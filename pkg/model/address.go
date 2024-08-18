package model

type Address struct {
	Street     string `json:"street,omitempty"`
	Number     string `json:"number,omitempty"`
	District   string `json:"district,omitempty"`
	Complement string `json:"complement,omitempty"`
	City       string `json:"city,omitempty"`
	State      string `json:"state,omitempty"`
	ZipCode    string `json:"zipcode,omitempty"`
}
