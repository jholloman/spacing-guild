package memproxysvc

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/bradfitz/gomemcache/memcache"
	"context"
)

type Endpoints struct {
	AddEndpoint            endpoint.Endpoint
	CompareAndSwapEndpoint endpoint.Endpoint
	DecrementEndpoint      endpoint.Endpoint
	DeleteEndpoint         endpoint.Endpoint
	GetEndpoint            endpoint.Endpoint
	GetMultiEndpoint       endpoint.Endpoint
	IncrementEndpoint      endpoint.Endpoint
	ReplaceEndpoint        endpoint.Endpoint
	SetEndpoint            endpoint.Endpoint
	TouchEndpoint          endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		AddEndpoint:            MakeAddEndpoint(s),
		CompareAndSwapEndpoint: MakeCompareAndSwapEndpoint(s),
		DecrementEndpoint:      MakeDecrementEndpoint(s),
		DeleteEndpoint:         MakeDeleteEndpoint(s),
		GetEndpoint:            MakeGetEndpoint(s),
		GetMultiEndpoint:       MakeGetMultiEndpoint(s),
		IncrementEndpoint:      MakeIncrementEndpoint(s),
		ReplaceEndpoint:        MakeReplaceEndpoint(s),
		SetEndpoint:            MakeSetEndpoint(s),
		TouchEndpoint:          MakeTouchEndpoint(s),
	}
}

func MakeAddEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(itemRequest)
		e := s.Add(ctx, req.Item)
		return baseResponse{Err: e}, nil
	}
}

func MakeCompareAndSwapEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(itemRequest)
		e := s.CompareAndSwap(ctx, req.Item)
		return baseResponse{Err: e}, nil
	}
}

func MakeDecrementEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(incrDecrRequest)
		n, e := s.Decrement(ctx, req.Key, req.Delta)
		return incrDecrResponse{NewValue: n, Err: e}, nil
	}
}

func MakeDeleteEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(keyRequest)
		e := s.Delete(ctx, req.Key)
		return baseResponse{Err: e}, nil
	}
}

func MakeGetEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(keyRequest)
		i, e := s.Get(ctx, req.Key)
		return getResponse{Item: i, Err: e}, nil
	}
}

func MakeGetMultiEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getMultiRequest)
		i, e := s.GetMulti(ctx, req.Keys)
		return getMultiResponse{Items: i, Err: e}, nil
	}
}

func MakeIncrementEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(incrDecrRequest)
		n, e := s.Increment(ctx, req.Key, req.Delta)
		return incrDecrResponse{NewValue: n, Err: e}, nil
	}
}

func MakeReplaceEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(itemRequest)
		e := s.Replace(ctx, req.Item)
		return baseResponse{Err: e}, nil
	}
}

func MakeSetEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(itemRequest)
		e := s.Set(ctx, req.Item)
		return baseResponse{Err: e}, nil
	}
}

func MakeTouchEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(touchRequest)
		e := s.Touch(ctx, req.Key, req.Seconds)
		return baseResponse{Err: e}, nil
	}
}

type itemRequest struct {
	Item *memcache.Item
}

type baseResponse struct {
	Err error `json:"err,omitempty"`
}
func (r baseResponse) error() error { return r.Err }

type incrDecrRequest struct {
	Key   string
	Delta uint64
}

type incrDecrResponse struct {
	NewValue uint64 `json:"newvalue,omitempty"`
	Err      error  `json:"err,omitempty"`
}

func (r incrDecrResponse) error() error { return r.Err }

type keyRequest struct {
	Key string
}

type getResponse struct {
	Item *memcache.Item `json:"item,omitempty"`
	Err  error          `json:"err,omitempty"`
}

func (r getResponse) error() error { return r.Err }

type getMultiRequest struct {
	Keys []string
}

type getMultiResponse struct {
	Items map[string]*memcache.Item `json:"items,omitempty"`
	Err   error            `json:"err,omitempty"`
}

func (r getMultiResponse) error() error { return r.Err }

type touchRequest struct {
	Key     string
	Seconds int32
}