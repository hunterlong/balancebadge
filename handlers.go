package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.Handle("/token/{token}/{address}.svg", http.HandlerFunc(TokenBadgeHandler))
	r.Handle("/token/{token}/{address}/usd.svg", http.HandlerFunc(TokenBadgeUSDHandler))
	r.Handle("/{coin}/{address}.svg", http.HandlerFunc(NormalBadgeHandler))
	r.Handle("/{coin}/{address}/usd.svg", http.HandlerFunc(USDBadgeHandler))
	r.Handle("/", http.HandlerFunc(IndexHandler))
	return r
}

func StartHTTPServer() {
	fmt.Println("BALANCE BADGE running on http://localhost:9090")
	r := Router()
	srv := &http.Server{
		Addr:         "0.0.0.0:9090",
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
		IdleTimeout:  time.Second * 30,
		Handler:      r,
	}
	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	stats := &Server{
		Online:   true,
		HourHits: statusHourHits,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stats)
}

func USDBadgeHandler(w http.ResponseWriter, r *http.Request) {
	badge := NewBadge(r).SetBalance().SetUSD().Clean()
	badge.Serve(w, r)
}

func TokenBadgeUSDHandler(w http.ResponseWriter, r *http.Request) {
	badge := NewBadge(r).TokenBadge().SetUSD().Clean()
	badge.Serve(w, r)
}

func TokenBadgeHandler(w http.ResponseWriter, r *http.Request) {
	badge := NewBadge(r).TokenBadge().Clean()
	badge.Serve(w, r)
}

func NormalBadgeHandler(w http.ResponseWriter, r *http.Request) {
	badge := NewBadge(r).SetBalance().Clean()
	badge.Serve(w, r)
}

func (b *Badge) Serve(w http.ResponseWriter, r *http.Request) {
	load := time.Now().Sub(b.start)
	fmt.Printf("%v | %v | %v | %v | %v | %v | %v\n", time.Now().Format(time.RFC3339), load, r.URL.Path, b.Coin, b.FullAddress, r.RemoteAddr, r.Header.Get("X-Forwarded-For"))
	w.Header().Set("Content-Type", "image/svg+xml")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	if b.error == nil {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
	svgTemplate.Execute(w, b)
	statusHourHits++
}

func httpGet(url string, method string, data []byte) (*http.Response, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	httpClient := &http.Client{Timeout: 30 * time.Second}
	resp, err := httpClient.Do(req)
	return resp, err
}
