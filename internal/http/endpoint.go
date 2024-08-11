package http

import (
	"errors"
	"net/http"

	"github.com/lcslucas/ipag-sdk-go/utils"
)

type Method string

const (
	GET    Method = http.MethodGet
	POST   Method = http.MethodPost
	PUT    Method = http.MethodPut
	DELETE Method = http.MethodDelete
)

type Endpoint struct {
	Method Method
	URI    string
}

type option func(*Endpoint)

func NewEndpoint(opts ...option) *Endpoint {
	e := &Endpoint{
		Method: GET,
		URI:    "",
	}

	for _, opt := range opts {
		opt(e)
	}

	return e
}

func WithMethod(method Method) option {
	methodsValid := []Method{
		GET,
		POST,
		PUT,
		DELETE,
	}

	if !utils.ExistsIn(methodsValid, method) {
		panic(errors.New("method not found in context"))
	}

	return func(e *Endpoint) {
		e.Method = method
	}
}

func WithURI(uri string) option {
	return func(e *Endpoint) {
		e.URI = uri
	}
}
