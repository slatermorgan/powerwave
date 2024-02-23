package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"powerwave/model"
	"time"

	"github.com/gorilla/mux"
)

// Service aims to decouple from database / handler logic
type Service interface {
	// DevicesByCustomerID retrieves a list of power meters installed at any of this customers buildings
	DevicesByCustomerID(id string) []model.Reading

	//DeviceReading at time T
	DeviceReading(id string, t time.Time) model.Reading
}

var service Service // Service implementation should be provided before running the server

func main() {
	http.HandleFunc("/customer/{id}/meters", makeHandler(customerDevices))
	http.HandleFunc("/device/{id}", makeHandler(device))

	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type handlerFunc func(http.ResponseWriter, *http.Request) error

func makeHandler(fn handlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// customerDevices returns all the devices for a given customer id
func customerDevices(w http.ResponseWriter, r *http.Request) error {
	customerID := mux.Vars(r)[IDPath]
	devices := service.DevicesByCustomerID(customerID)

	return json.NewEncoder(w).Encode(devices)
}

// REST API query parameter keys
const dateQuery = "date"
const IDPath = "id"

// device returns the device reading at time T
func device(w http.ResponseWriter, r *http.Request) error {
	serialID := r.URL.Query().Get(IDPath)
	dateStr := r.URL.Query().Get(dateQuery)

	date, err := time.Parse(time.Layout, dateStr)
	if err != nil {
		return fmt.Errorf("invalid date format")
	}

	deviceReading := service.DeviceReading(serialID, date)

	return json.NewEncoder(w).Encode(deviceReading)
}
