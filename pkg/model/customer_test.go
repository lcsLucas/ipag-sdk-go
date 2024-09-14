package model_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/lcslucas/ipag-sdk-go/pkg/model"
)

func TestCustomerSerialization(t *testing.T) {
	tBirthDate := time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC)

	addr := &model.Address{
		Street:     "123 Main St",
		Number:     "100",
		District:   "Downtown",
		Complement: "Apt 4B",
		City:       "Sample City",
		State:      "SC",
		ZipCode:    "12345-678",
	}

	customer := model.Customer{
		Name:            "John Doe",
		IsActive:        true,
		Email:           "johndoe@example.com",
		Phone:           "555-555-5555",
		CpfCnpj:         "748.980.410-86",
		TaxReceipt:      "123456789",
		BusinessName:    "John's Business",
		BirthDate:       &tBirthDate,
		Address:         addr,
		BillingAddress:  addr,
		ShippingAddress: addr,
	}

	jsonData, err := json.Marshal(customer)

	if err != nil {
		t.Fatalf("Failed to serialize Customer: %v", err)
	}

	var deserializedCustomer model.Customer

	err = json.Unmarshal(jsonData, &deserializedCustomer)

	if err != nil {
		t.Fatalf("Failed to deserialize JSON to Customer: %v", err)
	}

	if customer.ID != deserializedCustomer.ID ||
		customer.UUID != deserializedCustomer.UUID ||
		customer.Name != deserializedCustomer.Name ||
		customer.IsActive != deserializedCustomer.IsActive ||
		customer.Email != deserializedCustomer.Email ||
		customer.Phone != deserializedCustomer.Phone ||
		customer.CpfCnpj != deserializedCustomer.CpfCnpj ||
		customer.TaxReceipt != deserializedCustomer.TaxReceipt ||
		customer.BusinessName != deserializedCustomer.BusinessName ||
		!customer.BirthDate.Equal(*deserializedCustomer.BirthDate) ||
		customer.Ip != deserializedCustomer.Ip ||
		*customer.Address != *deserializedCustomer.Address ||
		*customer.BillingAddress != *deserializedCustomer.BillingAddress ||
		*customer.ShippingAddress != *deserializedCustomer.ShippingAddress {
		t.Errorf("Deserialized Customer does not match the original. Expected %v, got %v", customer, deserializedCustomer)
	}
}

func TestCustomerSerializationWithOmittedFields(t *testing.T) {
	customer := model.Customer{
		Name:  "John Doe",
		Email: "johndoe@example.com",
	}

	jsonData, err := json.Marshal(customer)

	if err != nil {
		t.Fatalf("Failed to serialize Customer: %v", err)
	}

	var deserializedCustomer model.Customer

	err = json.Unmarshal(jsonData, &deserializedCustomer)

	if err != nil {
		t.Fatalf("Failed to deserialize JSON to Customer: %v", err)
	}

	if customer.Name != deserializedCustomer.Name ||
		customer.Email != deserializedCustomer.Email {
		t.Errorf("Deserialized Customer does not match the original. Expected %+v, got %+v", customer, deserializedCustomer)
	}
}
