package main

import (
	"context"
	"time"

	"github.com/lcslucas/ipag-sdk-go/config"
	"github.com/lcslucas/ipag-sdk-go/pkg/model"
	customerService "github.com/lcslucas/ipag-sdk-go/service/customer"
	customerMiddleware "github.com/lcslucas/ipag-sdk-go/service/customer/middleware"
	"github.com/lcslucas/ipag-sdk-go/utils"
	"github.com/prometheus/client_golang/prometheus"
)

func prepareInstrumenting() (cMethods *prometheus.CounterVec, lMethods customerMiddleware.LatencyMethods) {
	cMethods = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "ipag_sdk",
			Subsystem: "customer_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		},
		[]string{"method"},
	)

	lMethods = customerMiddleware.LatencyMethods{
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

	countsService, latencyService := prepareInstrumenting()

	var service customerService.Service
	{
		service = customerService.NewService(config.Config{})
		service = customerMiddleware.InstrumentingMiddleware(countsService, latencyService)(service)
	}

	err := service.Save(ctx, &customer)

	if err != nil {
		panic(err)
	}

	// Do something with the customer
	utils.PrintPretty(customer)

}
