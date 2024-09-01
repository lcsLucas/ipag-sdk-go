package customer

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	internalHttp "github.com/lcslucas/ipag-sdk-go/internal/http"
	"github.com/lcslucas/ipag-sdk-go/pkg/model"
)

type serializeMiddleware struct {
	next *customerService
}

func SerializeMiddleware() func(*customerService) Service {
	return func(next *customerService) Service {
		return &serializeMiddleware{
			next: next,
		}
	}
}

func contextWithRequest(ctx context.Context, r *http.Request) context.Context {
	return context.WithValue(ctx, ContextRequestKey, r)
}

func (sw serializeMiddleware) Save(ctx context.Context, customer *model.Customer) (err error) {
	endpoint, ok := ctx.Value(ContextEndpointKey).(*internalHttp.Endpoint)

	if !ok {
		return errors.New("endpoint not found in context")
	}

	requestData, err := json.Marshal(customer)

	if err != nil {
		return fmt.Errorf("failed to parse request data: %w", err)
	}

	endpointURL, err := url.Parse(fmt.Sprintf("%s/%s", sw.next.config.Credentials.Environment, endpoint.URI))

	if err != nil {
		return fmt.Errorf("failed to parse endpoint url: %w", err)
	}

	request, err := http.NewRequestWithContext(ctx, "POST", endpointURL.String(), bytes.NewBuffer(requestData))

	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	basicAuth := fmt.Sprintf("%s:%s", sw.next.config.Credentials.ApiID, sw.next.config.Credentials.ApiKey)

	request.Header.Add("Content-Type", internalHttp.ContentTypeDefault)
	request.Header.Add("User-Agent", internalHttp.UserAgentDefault)
	request.Header.Add("x-api-version", fmt.Sprintf("%d", sw.next.config.Credentials.Version))
	request.Header.Add("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(basicAuth))))

	ctxRequest := contextWithRequest(ctx, request)

	err = sw.next.Save(ctxRequest, customer)
	return
}
func (sw serializeMiddleware) Update(ctx context.Context, customer *model.Customer) (err error) {
	err = sw.next.Update(ctx, customer)
	return
}
func (sw serializeMiddleware) Find(ctx context.Context, id uint32) (c model.Customer, err error) {
	c, err = sw.next.Find(ctx, id)
	return
}
func (sw serializeMiddleware) FindAll(ctx context.Context, filters map[string]interface{}) (cs []model.Customer, err error) {
	cs, err = sw.next.FindAll(ctx, filters)
	return
}
func (sw serializeMiddleware) Delete(ctx context.Context, id uint32) (err error) {
	err = sw.next.Delete(ctx, id)
	return
}
