package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/lcslucas/ipag-sdk-go/pkg/model"
	customerService "github.com/lcslucas/ipag-sdk-go/service/customer"
	instrumenting "github.com/lcslucas/ipag-sdk-go/service/customer/instrumenting"
	"github.com/lcslucas/ipag-sdk-go/utils"
	"github.com/prometheus/client_golang/prometheus"
)

func prepareInstrumenting() (cMethods *prometheus.CounterVec, lMethods instrumenting.LatencyMethods) {
	cMethods = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "ipag_sdk",
			Subsystem: "customer_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		},
		[]string{"method"},
	)

	lMethods = instrumenting.LatencyMethods{
		Save: prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Namespace: "ipag_sdk",
				Subsystem: "customer_service",
				Name:      "save_method_seconds_latency",
				Help:      "Latency for save method requests.",
			},
		),
		Update: prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Namespace: "ipag_sdk",
				Subsystem: "customer_service",
				Name:      "update_method_seconds_latency",
				Help:      "Latency for update method requests.",
			},
		),
		Find: prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Namespace: "ipag_sdk",
				Subsystem: "customer_service",
				Name:      "find_method_seconds_latency",
				Help:      "Latency for find method requests.",
			},
		),
		FindAll: prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Namespace: "ipag_sdk",
				Subsystem: "customer_service",
				Name:      "find_all_method_seconds_latency",
				Help:      "Latency for find all method requests.",
			},
		),
		Delete: prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Namespace: "ipag_sdk",
				Subsystem: "customer_service",
				Name:      "delete_method_seconds_latency",
				Help:      "Latency for delete method requests.",
			},
		),
	}

	prometheus.MustRegister(cMethods)

	prometheus.MustRegister(lMethods.Save)
	prometheus.MustRegister(lMethods.Update)
	prometheus.MustRegister(lMethods.Find)
	prometheus.MustRegister(lMethods.FindAll)
	prometheus.MustRegister(lMethods.Delete)

	return
}

func main() {
	ctx := context.TODO()

	addr := &model.Address{
		Street:     "123 Main St",
		Number:     json.Number("100"),
		District:   "Downtown",
		Complement: "Apt 4B",
		City:       "Sample City",
		State:      "SC",
		ZipCode:    "12345-678",
	}

	customer := model.Customer{
		ID:              "1",
		UUID:            "123e4567-e89b-12d3-a456-426614174000",
		Name:            "John Doe",
		IsActive:        true,
		Email:           "johndoe@example.com",
		Phone:           "555-555-5555",
		CpfCnpj:         "123.456.789-00",
		TaxReceipt:      "123456789",
		BusinessName:    "John's Business",
		BirthDate:       time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
		Ip:              "192.168.1.1",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		Address:         addr,
		BillingAddress:  addr,
		ShippingAddress: addr,
	}

	countsService, latencyService := prepareInstrumenting()

	var service customerService.Service
	{
		service = customerService.NewService()
		service = instrumenting.InstrumentingMiddleware(countsService, latencyService)(service)
	}

	err := service.Save(ctx, &customer)

	if err != nil {
		panic(err)
	}

	// Do something with the customer
	utils.PrintPretty(customer)

}
