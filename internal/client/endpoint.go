package client

const (
	GET     = "GET"
	POST    = "POST"
	PUT     = "PUT"
	DELETE  = "DELETE"
	PATCH   = "PATCH"
	HEAD    = "HEAD"
	OPTIONS = "OPTIONS"
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
