package customer

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/lcslucas/ipag-sdk-go/config"
	"github.com/lcslucas/ipag-sdk-go/internal/client"
	"github.com/lcslucas/ipag-sdk-go/pkg/model"
)

type contextKey string

const (
	ContextEndpointKey contextKey = "customer-endpoint"
)

type Service interface {
	Save(ctx context.Context, customer *model.Customer) error
	Update(ctx context.Context, customer *model.Customer) error
	Find(ctx context.Context, id uint32) (model.Customer, error)
	FindAll(ctx context.Context, filters map[string]interface{}) ([]model.Customer, error)
	Delete(ctx context.Context, id uint32) error
}

type customerService struct {
	client client.HTTPClient
	config config.Config
}

type ServiceMiddleware func(Service) Service

func NewService(config config.Config) Service {
	return EndpointMiddleware()(&customerService{
		client: client.NewHTTPClient(),
		config: config,
	})

}

func (c *customerService) Save(ctx context.Context, customer *model.Customer) error {
	endpoint, ok := ctx.Value(ContextEndpointKey).(client.Endpoint)

	if !ok {
		return errors.New("endpoint not found in context")
	}

	//TODO: mover tudo esse código para uma função que prepara Request
	endpointURL, err := url.Parse(fmt.Sprintf("%s/%s", c.config.Credentials.Environment, endpoint.URI))

	if err != nil {
		return errors.New("endpoint url unprocessable in context")
	}

	req, err := http.NewRequestWithContext(ctx, endpoint.Method, endpointURL.String(), nil)

	if err != nil {
		return err
	}

	userBasicAuth := fmt.Sprintf("%s:%s", c.config.Credentials.ApiID, c.config.Credentials.ApiKey)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "github.com/lcslucas/ipag-sdk-go")
	req.Header.Add("x-api-version", fmt.Sprintf("%d", c.config.Credentials.Version))
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(userBasicAuth))))

	fmt.Println(req.Header)

	// res, err := c.client.Do()

	// fmt.Println(res, err)

	return errors.New("not implemented")
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
