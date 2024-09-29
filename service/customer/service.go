package customer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/lcslucas/ipag-sdk-go/config"
	internalHttp "github.com/lcslucas/ipag-sdk-go/internal/http"
	"github.com/lcslucas/ipag-sdk-go/pkg/model"
)

type contextKey string

const (
	ContextEndpointKey contextKey = "customer-endpoint"
	ContextRequestKey  contextKey = "customer-request"
)

type Service interface {
	Save(ctx context.Context, customer *model.Customer) error
	Update(ctx context.Context, customer *model.Customer) error
	Find(ctx context.Context, id uint32) (model.Customer, error)
	FindAll(ctx context.Context, filters map[string]interface{}) ([]model.Customer, error)
	Delete(ctx context.Context, id uint32) error
	Config() config.Config
	Request() *http.Request
}

type customerService struct {
	client  internalHttp.HTTPClient
	config  config.Config
	request *http.Request
}

type ServiceMiddleware func(Service) Service

func NewService(config config.Config) Service {
	baseService := &customerService{
		client: internalHttp.NewHTTPClient(),
		config: config,
	}

	return use(baseService, SerializeMiddleware(), EndpointMiddleware())

}

func use(service Service, middlewares ...ServiceMiddleware) Service {
	for _, mw := range middlewares {
		service = mw(service)
	}

	return service
}

func (c *customerService) Config() config.Config {
	return c.config
}

func (c *customerService) Request() *http.Request {
	return c.request
}

func (c *customerService) Save(ctx context.Context, customer *model.Customer) error {
	request, ok := ctx.Value(ContextRequestKey).(*http.Request)

	if !ok {
		return errors.New("request not found in context")
	}

	c.request = request

	response, err := c.client.Do(c.request)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return err
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		var errorData any
		err = fmt.Errorf("request failed with status code %d", response.StatusCode)

		if err := json.Unmarshal(body, &errorData); err != nil {
			return fmt.Errorf("failed to parse error data: %w", err)
		}

		marshalErrorData, _ := json.MarshalIndent(errorData, "", "  ")

		err = fmt.Errorf("%w: %s", err, marshalErrorData)

		return err
	}

	err = json.Unmarshal(body, customer)

	if err != nil {
		err = fmt.Errorf("failed to parse response data: %w", err)
	}

	return err
}

func (c *customerService) Update(ctx context.Context, customer *model.Customer) error {
	return errors.New("not implemented")
}

func (c *customerService) Find(ctx context.Context, id uint32) (model.Customer, error) {
	return model.Customer{}, errors.New("not implemented")
}

func (c *customerService) FindAll(ctx context.Context, filters map[string]interface{}) ([]model.Customer, error) {
	return []model.Customer{}, errors.New("not implemented")
}

func (c *customerService) Delete(ctx context.Context, id uint32) error {
	return errors.New("not implemented")
}
