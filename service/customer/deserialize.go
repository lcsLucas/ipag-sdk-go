package customer

import (
	"context"
	"net/http"

	"github.com/lcslucas/ipag-sdk-go/config"
	"github.com/lcslucas/ipag-sdk-go/pkg/model"
)

type deserializeMiddleware struct {
	next Service
}

func DeserializeMiddleware() ServiceMiddleware {
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

func (mw deserializeMiddleware) Save(ctx context.Context, customer *model.Customer) (err error) {
	err = mw.next.Save(ctx, customer)
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
