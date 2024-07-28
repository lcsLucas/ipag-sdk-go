package customer

import (
	"context"

	"github.com/lcslucas/ipag-sdk-go/internal/client"
	"github.com/lcslucas/ipag-sdk-go/pkg/model"
)

type endpoints struct {
	Save    client.Endpoint
	Update  client.Endpoint
	Find    client.Endpoint
	FindAll client.Endpoint
	Delete  client.Endpoint
}

var Endpoints = endpoints{
	Save:    client.NewEndpoint(client.POST, "/v1/customers"),
	Update:  client.NewEndpoint(client.PUT, "/v1/customers/{id}"),
	Find:    client.NewEndpoint(client.GET, "/v1/customers/{id}"),
	FindAll: client.NewEndpoint(client.GET, "/v1/customers"),
	Delete:  client.NewEndpoint(client.DELETE, "/v1/customers/{id}"),
}

type endpointMiddleware struct {
	next Service
}

func EndpointMiddleware() ServiceMiddleware {
	return func(next Service) Service {
		return endpointMiddleware{
			next: next,
		}
	}
}

func contextWithEndpoint(ctx context.Context, endpoint client.Endpoint) context.Context {
	return context.WithValue(ctx, contextEndpointKey, endpoint)
}

func (mw endpointMiddleware) Save(ctx context.Context, customer *model.Customer) (err error) {
	ctxValue := contextWithEndpoint(ctx, Endpoints.Save)

	err = mw.next.Save(ctxValue, customer)
	return
}
func (mw endpointMiddleware) Update(ctx context.Context, customer *model.Customer) (err error) {
	ctxValue := contextWithEndpoint(ctx, Endpoints.Update)

	err = mw.next.Update(ctxValue, customer)
	return
}
func (mw endpointMiddleware) Find(ctx context.Context, id uint32) (c model.Customer, err error) {
	ctxValue := contextWithEndpoint(ctx, Endpoints.Find)

	c, err = mw.next.Find(ctxValue, id)
	return
}
func (mw endpointMiddleware) FindAll(ctx context.Context, filters map[string]interface{}) (cs []model.Customer, err error) {
	ctxValue := contextWithEndpoint(ctx, Endpoints.FindAll)

	cs, err = mw.next.FindAll(ctxValue, filters)
	return
}
func (mw endpointMiddleware) Delete(ctx context.Context, id uint32) (err error) {
	ctxValue := contextWithEndpoint(ctx, Endpoints.Delete)

	err = mw.next.Delete(ctxValue, id)
	return
}
