package service

import (
	"encoding/json"
	"fmt"
	"os"
	"powerwave/model"
	"time"
)

// this uses the assumptions from part 1
type database struct {
	PowerMeters []struct {
		Building         string `json:"building"`
		Customer         string `json:"customer"`
		SerialID         string `json:"serialId"`
		DailyConsumption []struct {
			AmountConsumed int    `json:"amountConsumed"`
			Date           string `json:"date"`
		} `json:"dailyConsumption"`
	} `json:"powerMeters"`
}

type FakeService struct {
	database *database
}

// DevicesByCustomerID retrieves a list of power meters installed at any of this customers buildings
func (fs *FakeService) DevicesByCustomerID(id string) ([]model.Reading, error) {
	var result []model.Reading
	for _, pm := range fs.database.PowerMeters {
		if pm.Customer == id {
			result = append(result, model.Reading{
				SerialID: pm.SerialID,
			})
		}
	}
	return result, nil

}

// DeviceReading at time T
func (fs *FakeService) DeviceReading(id string, t time.Time) (*model.Reading, error) {
	for _, pm := range fs.database.PowerMeters {
		if pm.SerialID == id {
			for _, day := range pm.DailyConsumption {
				// get the date specified
				if t.Equal(stringToTime(day.Date)) {
					return &model.Reading{
						SerialID:      id,
						TotalConsumed: day.AmountConsumed,
					}, nil
				}
			}
		}
	}

	return nil, fmt.Errorf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())

}

func New(fileName string) *FakeService {
	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	result := &FakeService{
		// assign database to memory
		database: &database{},
	}

	err = json.Unmarshal(fileContent, result.database)
	if err != nil {
		panic(err)
	}

	return result
}

// stringToTime converts the time string in the json to base library time
func stringToTime(t string) time.Time {
	result, err := time.Parse("2006-01-02", t)
	if err != nil {
		panic(err)
	}

	return result
}
