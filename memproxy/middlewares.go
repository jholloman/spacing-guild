package memproxysvc

import (
	"github.com/go-kit/kit/log"
	"context"
	"github.com/bradfitz/gomemcache/memcache"
	"time"
)

type Middleware func(Service) Service

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next Service) Service {
		return &loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

type loggingMiddleware struct {
	next   Service
	logger log.Logger
}

func (mw loggingMiddleware) Add(ctx context.Context, item *memcache.Item) (err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "Add", "key", item.Key, "elapsed", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.Add(ctx, item)
}

func (mw loggingMiddleware) CompareAndSwap(ctx context.Context, item *memcache.Item) (err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "CompareAndSwap", "key", item.Key, "elapsed", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.CompareAndSwap(ctx, item)
}

func (mw loggingMiddleware) Decrement(ctx context.Context, key string, delta uint64) (newValue uint64, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "Decrement", "key", key, "elapsed", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.Decrement(ctx, key, delta)
}

func (mw loggingMiddleware) Delete(ctx context.Context, key string) (err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "Decrement", "key", key, "elapsed", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.Delete(ctx, key)
}

func (mw loggingMiddleware) Get(ctx context.Context, key string) (item *memcache.Item, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "Get", "key", key, "elapsed", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.Get(ctx, key)
}

func (mw loggingMiddleware) GetMulti(ctx context.Context, keys []string) (items map[string]*memcache.Item, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "GetMulti", "elapsed", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.GetMulti(ctx, keys)
}

func (mw loggingMiddleware) Increment(ctx context.Context, key string, delta uint64) (newValue uint64, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "increment", "key", key, "elapsed", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.Increment(ctx, key, delta)
}

func (mw loggingMiddleware) Replace(ctx context.Context, item *memcache.Item) (err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "Replace", "key", item.Key, "elapsed", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.Replace(ctx, item)
}

func (mw loggingMiddleware) Set(ctx context.Context, item *memcache.Item) (err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "Set", "key", item.Key, "elapsed", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.Set(ctx, item)
}

func (mw loggingMiddleware) Touch(ctx context.Context, key string, seconds int32) (err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "Touch", "key", key, "elapsed", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.Touch(ctx, key, seconds)
}