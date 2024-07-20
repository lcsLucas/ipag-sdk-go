package model

import (
	"time"
)

type Customer struct {
	ID           string    `json:"id,omitempty"`
	UUID         string    `json:"uuid,omitempty"`
	Name         string    `json:"name,omitempty"`
	IsActive     bool      `json:"is_active,omitempty"`
	Email        string    `json:"email,omitempty"`
	Phone        string    `json:"phone,omitempty"`
	CpfCnpj      string    `json:"cpf_cnpj,omitempty"`
	TaxReceipt   string    `json:"tax_receipt,omitempty"`
	BusinessName string    `json:"business_name,omitempty"`
	BirthDate    time.Time `json:"birthdate,omitempty"`
	Ip           string    `json:"ip,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`

	Address         *Address `json:"address,omitempty"`
	BillingAddress  *Address `json:"billing_address,omitempty"`
	ShippingAddress *Address `json:"shipping_address,omitempty"`
}
