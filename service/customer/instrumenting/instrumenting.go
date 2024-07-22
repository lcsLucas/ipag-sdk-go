package instrumenting

import (
	"context"
	"time"

	"github.com/lcslucas/ipag-sdk-go/pkg/model"
	"github.com/lcslucas/ipag-sdk-go/service/customer"
	"github.com/prometheus/client_golang/prometheus"
)

type LatencyMethods struct {
	Save    prometheus.Histogram
	Update  prometheus.Histogram
	Find    prometheus.Histogram
	FindAll prometheus.Histogram
	Delete  prometheus.Histogram
}

type instrumentingMiddleware struct {
	countMethods   *prometheus.CounterVec
	latencyMethods LatencyMethods
	next           customer.Service
}

func InstrumentingMiddleware(cMethods *prometheus.CounterVec, lMethods LatencyMethods) customer.ServiceMiddleware {
	return func(next customer.Service) customer.Service {
		return instrumentingMiddleware{
			countMethods:   cMethods,
			latencyMethods: lMethods,
			next:           next,
		}
	}
}

func (mw instrumentingMiddleware) registerGeralMetrics(next func(time.Time)) func(time.Time) {
	return func(begin time.Time) {
		mw.countMethods.WithLabelValues("total").Inc()
		next(begin)
	}
}

func (mw instrumentingMiddleware) Save(ctx context.Context, customer *model.Customer) (err error) {
	defer mw.registerGeralMetrics(func(begin time.Time) {
		mw.latencyMethods.Save.Observe(time.Since(begin).Seconds())
		mw.countMethods.WithLabelValues("save").Inc()
	})(time.Now())

	mw.next.Save(ctx, customer)
	return
}

func (mw instrumentingMiddleware) Update(ctx context.Context, customer *model.Customer) (err error) {
	defer mw.registerGeralMetrics(func(begin time.Time) {
		mw.latencyMethods.Update.Observe(time.Since(begin).Seconds())
		mw.countMethods.WithLabelValues("update").Inc()
	})(time.Now())

	mw.next.Update(ctx, customer)
	return
}

func (mw instrumentingMiddleware) Find(ctx context.Context, id uint32) (c model.Customer, err error) {
	defer mw.registerGeralMetrics(func(begin time.Time) {
		mw.latencyMethods.Find.Observe(time.Since(begin).Seconds())
		mw.countMethods.WithLabelValues("find").Inc()
	})(time.Now())

	c, err = mw.next.Find(ctx, id)
	return
}

func (mw instrumentingMiddleware) FindAll(ctx context.Context, filters map[string]interface{}) (cs []model.Customer, err error) {
	defer mw.registerGeralMetrics(func(begin time.Time) {
		mw.latencyMethods.FindAll.Observe(time.Since(begin).Seconds())
		mw.countMethods.WithLabelValues("findAll").Inc()
	})(time.Now())

	cs, err = mw.next.FindAll(ctx, filters)
	return
}

func (mw instrumentingMiddleware) Delete(ctx context.Context, id uint32) (err error) {
	defer mw.registerGeralMetrics(func(begin time.Time) {
		mw.latencyMethods.Delete.Observe(time.Since(begin).Seconds())
		mw.countMethods.WithLabelValues("delete").Inc()
	})(time.Now())

	err = mw.next.Delete(ctx, id)
	return
}
