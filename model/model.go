package model

// Building represents a building output model
type Building struct {
	Name     string    `json:"name"`
	Customer string    `json:"customer"`
	Meters   []Reading `json:"meters"`
}

// Reading represents a building power meter output model
type Reading struct {
	SerialID      string `json:"serialID"`
	TotalConsumed int    `json:"totalConsumed,omitempty"`
}
