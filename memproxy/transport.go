package memproxysvc

import (
	"github.com/go-kit/kit/log"
	"net/http"
	"github.com/gorilla/mux"
	httptransport "github.com/go-kit/kit/transport/http"
	"context"
	"encoding/json"
)

func MakeHTTPHandler(s Service, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	e := MakeServerEndpoints(s)
	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeError),
	}

	r.Methods("POST").Path("/item/add/").Handler(httptransport.NewServer(
		e.AddEndpoint,
		decodeItemRequest,
		encodeResponse,
		options...,
	))

	r.Methods("PUT").Path("/item/compareandswap/").Handler(httptransport.NewServer(
		e.CompareAndSwapEndpoint,
		decodeItemRequest,
		encodeResponse,
		options...,
	))

	r.Methods("PUT").Path("/item/decrement/").Handler(httptransport.NewServer(
		e.DecrementEndpoint,
		decodeIncrDecrRequest,
		encodeResponse,
		options...,
	))

	r.Methods("DELETE").Path("/item/").Handler(httptransport.NewServer(
		e.DeleteEndpoint,
		decodeKeyRequest,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path("/item/").Handler(httptransport.NewServer(
		e.GetEndpoint,
		decodeKeyRequest,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path("/items/").Handler(httptransport.NewServer(
		e.GetMultiEndpoint,
		decodeMultiKeyRequest,
		encodeResponse,
		options...,
	))

	r.Methods("PUT").Path("/item/increment/").Handler(httptransport.NewServer(
		e.IncrementEndpoint,
		decodeIncrDecrRequest,
		encodeResponse,
		options...,
	))

	r.Methods("PUT").Path("/item/replace/").Handler(httptransport.NewServer(
		e.ReplaceEndpoint,
		decodeItemRequest,
		encodeResponse,
		options...,
	))

	r.Methods("POST").Path("/item/").Handler(httptransport.NewServer(
		e.SetEndpoint,
		decodeItemRequest,
		encodeResponse,
		options...,
	))

	r.Methods("PUT").Path("/item/touch/").Handler(httptransport.NewServer(
		e.TouchEndpoint,
		decodeTouchRequest,
		encodeResponse,
		options...,
	))

	return r
}

func decodeItemRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req itemRequest
	if e := json.NewDecoder(r.Body).Decode(&req.Item); e != nil {
		return nil, e
	}
	return req, nil
}

func decodeIncrDecrRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req incrDecrRequest
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

func decodeKeyRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req keyRequest
	e := r.ParseForm()
	if e != nil {
		return nil, e
	}
	req.Key = r.Form.Get("key")
	return req, nil
}

func decodeMultiKeyRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req getMultiRequest
	e := r.ParseForm()
	if e != nil {
		return nil, e
	}
	req.Keys = r.Form["keys"]
	return req, nil
}

func decodeTouchRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req touchRequest
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

type errorer interface {
	error() error
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}