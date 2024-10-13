package customer

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/lcslucas/ipag-sdk-go/config"
	"github.com/lcslucas/ipag-sdk-go/pkg/model"
)

type deserializeMiddleware struct {
	next Service
}

func deserialize() ServiceMiddleware {
	return func(next Service) Service {
		return &deserializeMiddleware{
			next: next,
		}
	}
}

// func (m *deserializeMiddleware) GetCustomer(id string) (c Customer, err error) {
// 	c, err = m.Next.GetCustomer(id)
// 	if err != nil {
// 		return
// 	}

// 	//TODO: criar uma funçao genérica para deserializar no internal
// 	c.Deserialize()
// 	return
// }

func (mw deserializeMiddleware) Config() config.Config {
	return mw.next.Config()
}

func (mw deserializeMiddleware) Request() *http.Request {
	return mw.next.Request()
}

func (mw deserializeMiddleware) Response() *http.Response {
	return mw.next.Response()
}

func (mw deserializeMiddleware) Save(ctx context.Context, customer *model.Customer) (err error) {
	err = mw.next.Save(ctx, customer)

	if err != nil {
		return
	}

	response := mw.next.Response()

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return err
	}

	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
		var errorData any
		err = fmt.Errorf("request failed with status code %d", response.StatusCode)

		if err := json.Unmarshal(body, &errorData); err != nil {
			return fmt.Errorf("failed to parse error data: %w", err)
		}

		marshalErrorData, _ := json.MarshalIndent(errorData, "", "  ")

		err = fmt.Errorf("%w: %s", err, marshalErrorData)

		return err
	}

	err = json.Unmarshal(body, &customer)

	if err != nil {
		err = fmt.Errorf("failed to parse response data: %w", err)
	}

	return
}

func (mw deserializeMiddleware) Update(ctx context.Context, customer *model.Customer) (err error) {
	err = mw.next.Update(ctx, customer)
	return
}

func (mw deserializeMiddleware) Find(ctx context.Context, id uint32) (c model.Customer, err error) {
	c, err = mw.next.Find(ctx, id)
	return
}

func (mw deserializeMiddleware) FindAll(ctx context.Context, filters map[string]interface{}) (cs []model.Customer, err error) {
	cs, err = mw.next.FindAll(ctx, filters)
	return
}

func (mw deserializeMiddleware) Delete(ctx context.Context, id uint32) (err error) {
	err = mw.next.Delete(ctx, id)
	return
}
