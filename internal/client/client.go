package client

import "net/http"

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type httpClient struct {
	client   *http.Client
	Endpoint Endpoint
}

func NewHTTPClient() *httpClient {
	return &httpClient{
		client: &http.Client{},
	}
}

func (httpClient *httpClient) Do(req *http.Request) (*http.Response, error) {
	return httpClient.client.Do(req)
}
