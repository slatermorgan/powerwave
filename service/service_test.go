package service

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// I would write more test cases here if I had time
func TestDevicesByCustomerID(t *testing.T) {
	service := New("../data.json")

	devices, err := service.DevicesByCustomerID("Aquaflow")

	assert.NoError(t, err)
	assert.Len(t, devices, 2)
	assert.Equal(t, "1111-1111-1111", devices[0].SerialID)
	assert.Equal(t, "1111-1111-2222", devices[1].SerialID)
}

// I would write more test cases here if I had time
func TestDeviceReading(t *testing.T) {
	service := New("../data.json")

	time := time.Date(2024, 02, 29, 0, 0, 0, 0, time.UTC)

	device, err := service.DeviceReading("1111-1111-1111", time)

	assert.NoError(t, err)
	assert.Equal(t, 140, device.TotalConsumed)
}
