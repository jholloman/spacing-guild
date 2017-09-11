package memproxysvc

import (
	"github.com/bradfitz/gomemcache/memcache"
	"context"
	"github.com/davecgh/go-spew/spew"
)

// Passthrough for gomemcache
type Service interface {
	Add(ctx context.Context, item *memcache.Item) error
	CompareAndSwap(ctx context.Context, item *memcache.Item) error
	Decrement(ctx context.Context, key string, delta uint64) (newValue uint64, err error)
	Delete(ctx context.Context, key string) error
	Get(ctx context.Context, key string) (item *memcache.Item, err error)
	GetMulti(ctx context.Context, keys []string) (map[string]*memcache.Item, error)
	Increment(ctx context.Context, key string, delta uint64) (newValue uint64, err error)
	Replace(ctx context.Context, item *memcache.Item) error
	Set(ctx context.Context, item *memcache.Item) error
	Touch(ctx context.Context, key string, seconds int32) (err error)
}

// memproxyService is the Service implementation and Service dependencies
type memproxyService struct {
	client*memcache.Client
}

// NewService returns a new memproxy Service
func NewService(client *memcache.Client) Service {
	return &memproxyService{
		client: client,
	}
}

func (s *memproxyService) Add(ctx context.Context, item *memcache.Item) error {
	return s.client.Add(item)
}

func (s *memproxyService) CompareAndSwap(ctx context.Context, item *memcache.Item) error {
	return s.client.CompareAndSwap(item)
}

func (s *memproxyService) Decrement(ctx context.Context, key string, delta uint64) (newValue uint64, err error) {
	return s.client.Decrement(key, delta)
}

func (s *memproxyService) Delete(ctx context.Context, key string) error {
	return s.client.Delete(key)
}

func (s *memproxyService) Get(ctx context.Context, key string) (item *memcache.Item, err error) {
	return s.client.Get(key)
}

func (s *memproxyService) GetMulti(ctx context.Context, keys []string) (map[string]*memcache.Item, error) {
	return s.client.GetMulti(keys)
}

func (s *memproxyService) Increment(ctx context.Context, key string, delta uint64) (newValue uint64, err error) {
	return s.client.Increment(key, delta)
}

func (s *memproxyService) Replace(ctx context.Context, item *memcache.Item) error {
	return s.client.Replace(item)
}

func (s *memproxyService) Set(ctx context.Context, item *memcache.Item) error {
	return s.client.Set(item)
}

func (s *memproxyService) Touch(ctx context.Context, key string, seconds int32) (err error) {
	return s.client.Touch(key, seconds)
}