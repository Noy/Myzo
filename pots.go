package myzo

import "time"

type PotResponse struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Style string `json:"style"`
	Balance int64 `json:"balance"`
	Currency string `json:"currency"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
	Deleted bool `json:"deleted"`
}
