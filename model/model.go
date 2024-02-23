package model

// Building represents a building where power meters are installed.
type Building struct {
	Name     string
	Customer string
	Meters   []Reading
}

// Reading represents a building power meter.
type Reading struct {
	SerialID      string
	TotalConsumed int //since device began
}
