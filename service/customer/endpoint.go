package customer

import (
	"context"
	"net/http"

	"github.com/lcslucas/ipag-sdk-go/config"
	internalHttp "github.com/lcslucas/ipag-sdk-go/internal/http"
	"github.com/lcslucas/ipag-sdk-go/pkg/model"
)

type endpoints struct {
	Save    *internalHttp.Endpoint
	Update  *internalHttp.Endpoint
	Find    *internalHttp.Endpoint
	FindAll *internalHttp.Endpoint
	Delete  *internalHttp.Endpoint
}

var Endpoints = endpoints{
	Save: internalHttp.NewEndpoint(
		internalHttp.WithMethod(internalHttp.POST),
		internalHttp.WithURI("service/resources/customers")),
	// Update:  client.NewEndpoint(client.PUT, "v1/customers/{id}"),
	// Find:    client.NewEndpoint(client.GET, "v1/customers/{id}"),
	// FindAll: client.NewEndpoint(client.GET, "v1/customers"),
	// Delete:  client.NewEndpoint(client.DELETE, "v1/customers/{id}"),
}

type endpointMiddleware struct {
	next Service
}

func EndpointMiddleware() ServiceMiddleware {
	return func(next Service) Service {
		return &endpointMiddleware{
			next: next,
		}
	}
}

func contextWithEndpoint(ctx context.Context, endpoint *internalHttp.Endpoint) context.Context {
	return context.WithValue(ctx, ContextEndpointKey, endpoint)
}

func (mw endpointMiddleware) Config() config.Config {
	return mw.next.Config()
}

func (mw endpointMiddleware) Request() *http.Request {
	return mw.next.Request()
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
