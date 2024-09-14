package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/lcslucas/ipag-sdk-go/config"
	"github.com/lcslucas/ipag-sdk-go/credentials"
	"github.com/lcslucas/ipag-sdk-go/pkg/model"
	customerService "github.com/lcslucas/ipag-sdk-go/service/customer"
	"github.com/lcslucas/ipag-sdk-go/utils"
)

func main() {
	ctx := context.TODO()

	tBirth := time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC)

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
		BirthDate:       &tBirth,
		Address:         addr,
		BillingAddress:  addr,
		ShippingAddress: addr,
	}

	var service customerService.Service
	{
		service = customerService.NewService(config.Config{
			Credentials: credentials.Credentials{
				ApiID:       os.Getenv("API_ID"),
				ApiKey:      os.Getenv("API_KEY"),
				Environment: credentials.Environments.Sandbox,
				Version:     2,
			},
			Client: config.ClientConfig{
				Timeout: 30 * time.Second,
			},
		})
	}

	err := service.Save(ctx, &customer)

	if err != nil {
		panic(err)
	}

	// Do something with the customer
	utils.PrintPretty(customer)

}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
