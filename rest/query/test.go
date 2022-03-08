package query

type Fleet struct {
	vehicles []Vehicle
}

type Vehicle struct {
	Class         string        `json:"__class__"`
	Available     bool          `json:"available"`
	Retired       bool          `json:"retired"`
	License       string        `json:"license"`
	ServRecs      []interface{} `json:"servRecs"`
	CurrentRental interface{}   `json:"currentRental"`
	PastRentals   []interface{} `json:"pastRentals"`
	Model         interface{}   `json:"model"`
	AtStall       interface{}   `json:"atStall"`
}
