package model_test

import (
	"encoding/json"
	"testing"

	"github.com/lcslucas/ipag-sdk-go/pkg/model"
)

func TestAddressSerialization(t *testing.T) {
	addr := model.Address{
		Street:     "123 Main St",
		Number:     "100",
		District:   "Downtown",
		Complement: "Apt 4B",
		City:       "Sample City",
		State:      "SC",
		ZipCode:    "12345-678",
	}

	jsonData, err := json.Marshal(addr)

	if err != nil {
		t.Fatalf("Failed to serialize Address: %v", err)
	}

	var deserializedAddr model.Address

	err = json.Unmarshal(jsonData, &deserializedAddr)

	if err != nil {
		t.Fatalf("Failed to deserialize JSON to Address: %v", err)
	}

	if addr != deserializedAddr {
		t.Errorf("Deserialized Address does not match the original. Expected %v, got %v", addr, deserializedAddr)
	}
}

func TestAddressSerializationWithOmittedFields(t *testing.T) {
	addr := model.Address{
		Street: "123 Main St",
		City:   "Sample City",
		State:  "SC",
	}

	jsonData, err := json.Marshal(addr)

	if err != nil {
		t.Fatalf("Failed to serialize Address: %v", err)
	}

	var deserializedAddr model.Address

	err = json.Unmarshal(jsonData, &deserializedAddr)

	if err != nil {
		t.Fatalf("Failed to deserialize JSON to Address: %v", err)
	}

	if addr != deserializedAddr {
		t.Errorf("Deserialized Address does not match the original. Expected %v, got %v", addr, deserializedAddr)
	}
}
