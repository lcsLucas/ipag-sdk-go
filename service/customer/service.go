package customer

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/lcslucas/ipag-sdk-go/config"
	"github.com/lcslucas/ipag-sdk-go/internal/http"
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
	client http.HTTPClient
	config config.Config
}

type ServiceMiddleware func(Service) Service

func NewService(config config.Config) Service {
	return EndpointMiddleware()(&customerService{
		client: http.NewHTTPClient(),
		config: config,
	})

}

func (c *customerService) Save(ctx context.Context, customer *model.Customer) error {
	endpoint, ok := ctx.Value(ContextEndpointKey).(*http.Endpoint)

	if !ok {
		return errors.New("endpoint not found in context")
	}

	requestData, err := json.Marshal(customer)

	if err != nil {
		return fmt.Errorf("failed to parse request data: %w", err)
	}

	request, err := c.client.BuilderRequest(ctx, endpoint, c.config, bytes.NewBuffer(requestData))

	if err != nil {
		return err
	}

	response, err := c.client.Do(request)

	if err != nil {
		return err
	}

	//TODO: refactor this
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
