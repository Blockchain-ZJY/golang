package main

import (
	"context"
	"encoding/json"
	"golang/microservice/types"
	"math/rand"
	"net/http"
)

type APIFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

type JSONAPIServer struct {
	listenAddr string
	svc        PriceFetcher
}

func NewJSONAPIServer(listenAddr string, svc PriceFetcher) *JSONAPIServer {
	return &JSONAPIServer{
		listenAddr: listenAddr,
		svc:        svc,
	}
}

func (s *JSONAPIServer) Run() {
	http.Handle("/", makeHTTPAPIFunc(s.handleFetchPrice))
	http.ListenAndServe(s.listenAddr, nil)
}

func makeHTTPAPIFunc(f APIFunc) http.HandlerFunc {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "requestID", rand.Intn(10000))
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(ctx, w, r); err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		}
	}
}

func (s *JSONAPIServer) handleFetchPrice(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ticker := r.URL.Query().Get("ticker")

	price, err := s.svc.FetchPrice(ctx, ticker)

	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return err
	}

	rsp := types.PriceResponse{
		Ticker: ticker,
		Price:  price,
	}
	return writeJSON(w, http.StatusOK, rsp)
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
