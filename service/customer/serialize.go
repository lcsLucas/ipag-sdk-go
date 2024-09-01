package customer

import (
	"context"

	"github.com/lcslucas/ipag-sdk-go/pkg/model"
)

type serializeMiddleware struct {
	next Service
}

func SerializeMiddleware() ServiceMiddleware {
	return func(next Service) Service {
		return &serializeMiddleware{
			next: next,
		}
	}
}

func (mw serializeMiddleware) Save(ctx context.Context, customer *model.Customer) (err error) {
	// endpoint, ok := ctx.Value(ContextEndpointKey).(*http.Endpoint)

	// if !ok {
	// 	return errors.New("endpoint not found in context")
	// }

	// endpointURL, err := url.Parse(fmt.Sprintf("%s/%s", config.Credentials.Environment, endpoint.URI))

	// requestData, err := json.Marshal(customer)

	// if err != nil {
	// 	return fmt.Errorf("failed to parse request data: %w", err)
	// }

	// request, err := c.client.BuilderRequest(ctx, endpoint, c.config, bytes.NewBuffer(requestData))

	err = mw.next.Save(ctx, customer)
	return
}
func (mw serializeMiddleware) Update(ctx context.Context, customer *model.Customer) (err error) {
	err = mw.next.Update(ctx, customer)
	return
}
func (mw serializeMiddleware) Find(ctx context.Context, id uint32) (c model.Customer, err error) {
	c, err = mw.next.Find(ctx, id)
	return
}
func (mw serializeMiddleware) FindAll(ctx context.Context, filters map[string]interface{}) (cs []model.Customer, err error) {
	cs, err = mw.next.FindAll(ctx, filters)
	return
}
func (mw serializeMiddleware) Delete(ctx context.Context, id uint32) (err error) {
	err = mw.next.Delete(ctx, id)
	return
}
