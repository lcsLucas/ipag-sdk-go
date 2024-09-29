package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/lcslucas/ipag-sdk-go/config"
	"github.com/lcslucas/ipag-sdk-go/pkg/model"
	"github.com/lcslucas/ipag-sdk-go/service/customer"
)

type loggingMiddleware struct {
	next   customer.Service
	logger log.Logger
}

func LoggingMiddleware(logger log.Logger) customer.ServiceMiddleware {
	return func(next customer.Service) customer.Service {
		return loggingMiddleware{next, logger}
	}
}

func (mw loggingMiddleware) Config() config.Config {
	return mw.next.Config()
}

func (mw loggingMiddleware) Request() *http.Request {
	return mw.next.Request()
}

func (mw loggingMiddleware) Save(ctx context.Context, customer *model.Customer) (err error) {
	defer func(begin time.Time) {
		logger := log.With(mw.logger, "service", "customer", "method", "save")
		level.Info(logger).Log(
			"parameter", fmt.Sprintf("%v", customer),
			"error", err,
			"ended", time.Now(),
			"duration", fmt.Sprintf("%vms", time.Since(begin).Milliseconds()),
		)
	}(time.Now())

	err = mw.next.Save(ctx, customer)
	return
}

func (mw loggingMiddleware) Update(ctx context.Context, customer *model.Customer) (err error) {
	defer func(begin time.Time) {
		logger := log.With(mw.logger, "service", "customer", "method", "update")
		level.Info(logger).Log(
			"parameter", fmt.Sprintf("%v", customer),
			"error", err,
			"ended", time.Now(),
			"duration", fmt.Sprintf("%vms", time.Since(begin).Milliseconds()),
		)
	}(time.Now())

	err = mw.next.Update(ctx, customer)
	return
}

func (mw loggingMiddleware) Find(ctx context.Context, id uint32) (c model.Customer, err error) {
	defer func(begin time.Time) {
		logger := log.With(mw.logger, "service", "customer", "method", "find")
		level.Info(logger).Log(
			"parameter", id,
			"error", err,
			"ended", time.Now(),
			"duration", fmt.Sprintf("%vms", time.Since(begin).Milliseconds()),
		)
	}(time.Now())

	c, err = mw.next.Find(ctx, id)
	return
}

func (mw loggingMiddleware) FindAll(ctx context.Context, filters map[string]interface{}) (cs []model.Customer, err error) {
	defer func(begin time.Time) {
		logger := log.With(mw.logger, "service", "customer", "method", "findAll")
		level.Info(logger).Log(
			"parameter", filters,
			"error", err,
			"ended", time.Now(),
			"duration", fmt.Sprintf("%vms", time.Since(begin).Milliseconds()),
		)
	}(time.Now())

	cs, err = mw.next.FindAll(ctx, filters)
	return
}

func (mw loggingMiddleware) Delete(ctx context.Context, id uint32) (err error) {
	defer func(begin time.Time) {
		logger := log.With(mw.logger, "service", "customer", "method", "delete")
		level.Info(logger).Log(
			"parameter", id,
			"error", err,
			"ended", time.Now(),
			"duration", fmt.Sprintf("%vms", time.Since(begin).Milliseconds()),
		)
	}(time.Now())

	err = mw.next.Delete(ctx, id)
	return
}
