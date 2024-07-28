package client

import "net/http"

const (
	GET     = http.MethodGet
	POST    = http.MethodPost
	PUT     = http.MethodPut
	DELETE  = http.MethodDelete
	PATCH   = http.MethodPatch
	HEAD    = http.MethodHead
	OPTIONS = http.MethodOptions
)

type Endpoint struct {
	Method string
	URI    string
}

func NewEndpoint(method, uri string) Endpoint {
	return Endpoint{
		Method: method,
		URI:    uri,
	}
}
