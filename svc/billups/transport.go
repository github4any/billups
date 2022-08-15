package billups

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/billups/api/internal/httpencoder"
	"github.com/go-chi/chi"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
)

type (
	logger interface {
		Log(keyvals ...interface{}) error
	}
)

// MakeHTTPHandler ...
func MakeHTTPHandler(e Endpoints, log logger) http.Handler {
	r := chi.NewRouter()

	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(log)),
		httptransport.ServerErrorEncoder(httpencoder.EncodeError(log, codeAndMessageFrom)),
	}

	r.Get("/choices", httptransport.NewServer(
		e.Choices,
		decodeDefaultRequest,
		httpencoder.EncodeResponse,
		options...,
	).ServeHTTP)

	r.Get("/choice", httptransport.NewServer(
		e.Choice,
		decodeDefaultRequest,
		httpencoder.EncodeResponse,
		options...,
	).ServeHTTP)

	r.Post("/play", httptransport.NewServer(
		e.Play,
		decodePlayRequest,
		httpencoder.EncodeResponse,
		options...,
	).ServeHTTP)

	r.Get("/scoreboard", httptransport.NewServer(
		e.Scoreboard,
		decodeDefaultRequest,
		httpencoder.EncodeResponse,
		options...,
	).ServeHTTP)

	r.Delete("/reset", httptransport.NewServer(
		e.Reset,
		decodeDefaultRequest,
		httpencoder.EncodeResponse,
		options...,
	).ServeHTTP)

	return r
}

func decodeDefaultRequest(ctx context.Context, _ *http.Request) (interface{}, error) {
	return nil, nil
}

func decodePlayRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req PlayerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, fmt.Errorf("could not decode request body: %w", err)
	}

	return req, nil
}

// returns http error code by error type
func codeAndMessageFrom(err error) (int, interface{}) {
	return httpencoder.CodeAndMessageFrom(err)
}
