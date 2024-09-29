package customer

import (
	"context"
	"errors"
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
	Response() *http.Response
}

type customerService struct {
	client   internalHttp.HTTPClient
	config   config.Config
	request  *http.Request
	response *http.Response
}

type ServiceMiddleware func(Service) Service

func NewService(config config.Config) Service {
	baseService := &customerService{
		client: internalHttp.NewHTTPClient(),
		config: config,
	}

	return use(baseService, deserialize(), serialize(), endpoint())

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

func (c *customerService) Response() *http.Response {
	return c.response
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

	c.response = response

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
