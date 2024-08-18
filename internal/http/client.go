package http

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/lcslucas/ipag-sdk-go/config"
)

const (
	UserAgentDefault   = "github.com/lcslucas/ipag-sdk-go"
	ContentTypeDefault = "application/json"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
	BuilderRequest(ctx context.Context, endpoint *Endpoint, config config.Config, body *bytes.Buffer) (*http.Request, error)
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

func (httpClient *httpClient) BuilderRequest(ctx context.Context, endpoint *Endpoint, config config.Config, body *bytes.Buffer) (*http.Request, error) {
	endpointURL, err := url.Parse(fmt.Sprintf("%s/%s", config.Credentials.Environment, endpoint.URI))

	if err != nil {
		return nil, errors.New("endpoint url unprocessable in context")
	}

	request, err := http.NewRequestWithContext(ctx, string(endpoint.Method), endpointURL.String(), body)

	if err != nil {
		return nil, err
	}

	userBasicAuth := fmt.Sprintf("%s:%s", config.Credentials.ApiID, config.Credentials.ApiKey)

	request.Header.Add("Content-Type", ContentTypeDefault)
	request.Header.Add("User-Agent", UserAgentDefault)
	request.Header.Add("x-api-version", fmt.Sprintf("%d", config.Credentials.Version))
	request.Header.Add("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(userBasicAuth))))

	return request, nil
}
