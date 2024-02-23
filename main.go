package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"powerwave/model"
	fakeService "powerwave/service"
	"time"

	"github.com/gorilla/mux"
)

// Service aims to decouple from database / handler logic
type Service interface {
	// DevicesByCustomerID retrieves a list of power meters installed at any of this customers buildings
	DevicesByCustomerID(id string) ([]model.Reading, error)

	//DeviceReading at time T
	DeviceReading(id string, t time.Time) (*model.Reading, error)
}

var service Service // Service implementation should be provided before running the server

func main() {
	r := mux.NewRouter()
	serv := fakeService.New("data.json")

	r.HandleFunc("/customer/{id}/meters", makeHandler(customerDevices, serv))
	r.HandleFunc("/device/{id}", makeHandler(device, serv))

	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// handler stuff would go into its own package if I had more time
type handlerFunc func(http.ResponseWriter, *http.Request, Service) error

func makeHandler(fn handlerFunc, service Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r, service); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// customerDevices returns all the devices for a given customer id
func customerDevices(w http.ResponseWriter, r *http.Request, service Service) error {
	customerID := mux.Vars(r)[IDPath]
	devices, err := service.DevicesByCustomerID(customerID)
	if err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(devices)
}

// REST API query parameter keys (single source of truth)
const dateQuery = "date"
const IDPath = "id"

// device returns the device reading at time T
func device(w http.ResponseWriter, r *http.Request, service Service) error {
	serialID := mux.Vars(r)[IDPath]
	dateStr := r.URL.Query().Get(dateQuery)
	fmt.Println(serialID)

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return fmt.Errorf("invalid date format")
	}

	deviceReading, err := service.DeviceReading(serialID, date)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("device not found"))
		return nil
	}

	return json.NewEncoder(w).Encode(deviceReading)
}
