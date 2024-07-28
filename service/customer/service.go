package customer

import (
	"context"
	"errors"

	"github.com/lcslucas/ipag-sdk-go/config"
	"github.com/lcslucas/ipag-sdk-go/internal/client"
	"github.com/lcslucas/ipag-sdk-go/pkg/model"
)

type contextKey string

const contextEndpointKey contextKey = "customer_endpoint"

type Service interface {
	Save(ctx context.Context, customer *model.Customer) error
	Update(ctx context.Context, customer *model.Customer) error
	Find(ctx context.Context, id uint32) (model.Customer, error)
	FindAll(ctx context.Context, filters map[string]interface{}) ([]model.Customer, error)
	Delete(ctx context.Context, id uint32) error
}

type customerService struct {
	client client.HTTPClient
}

type ServiceMiddleware func(Service) Service

func NewService(config config.Config) Service {
	c := client.NewHTTPClient()

	return EndpointMiddleware()(&customerService{
		client: c,
	})

}

func (c *customerService) Save(ctx context.Context, customer *model.Customer) error {
	// endpoint, ok := ctx.Value(contextEndpointKey).(client.Endpoint)

	// if !ok {
	// 	return errors.New("endpoint not found in context")
	// }

	// requestURL := fmt.Sprintf("%s/%s", endpoint., endpoint.Path)

	// req, err := http.NewRequestWithContext(ctx, endpoint.Method, endpoint.URI, nil)

	// if err != nil {
	// 	return err
	// }

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
