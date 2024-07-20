package model

import "encoding/json"

type Address struct {
	Street     string      `json:"street,omitempty"`
	Number     json.Number `json:"number,omitempty"`
	District   string      `json:"district,omitempty"`
	Complement string      `json:"complement,omitempty"`
	City       string      `json:"city,omitempty"`
	State      string      `json:"state,omitempty"`
	ZipCode    string      `json:"zipcode,omitempty"`
}
