package customer

import (
	"context"
	"errors"

	"github.com/lcslucas/ipag-sdk-go/config"
	"github.com/lcslucas/ipag-sdk-go/internal/http"
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
}

type customerService struct {
	client http.HTTPClient
	config config.Config
}

type ServiceMiddleware func(Service) Service

func NewService(config config.Config) Service {
	return EndpointMiddleware()(
		SerializeMiddleware()(
			&customerService{
				client: http.NewHTTPClient(),
				config: config,
			}))
}

func (c *customerService) Save(ctx context.Context, customer *model.Customer) error {

	// request, err := c.client.BuilderRequest(ctx, endpoint, c.config, bytes.NewBuffer(requestData))

	// if err != nil {
	// 	return err
	// }

	// response, err := c.client.Do(request)

	// if err != nil {
	// 	return err
	// }

	// //TODO: refactor this
	// defer response.Body.Close()
	// body, err := io.ReadAll(response.Body)

	// if err != nil {
	// 	return err
	// }

	// if response.StatusCode < 200 || response.StatusCode >= 300 {
	// 	var errorData any
	// 	err = fmt.Errorf("request failed with status code %d", response.StatusCode)

	// 	if err := json.Unmarshal(body, &errorData); err != nil {
	// 		return fmt.Errorf("failed to parse error data: %w", err)
	// 	}

	// 	marshalErrorData, _ := json.MarshalIndent(errorData, "", "  ")

	// 	err = fmt.Errorf("%w: %s", err, marshalErrorData)

	// 	return err
	// }

	// err = json.Unmarshal(body, customer)

	// if err != nil {
	// 	err = fmt.Errorf("failed to parse response data: %w", err)
	// }

	// return err
	return nil
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

/**
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
*/
