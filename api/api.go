package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand/v2"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr    string
	helloProducer HelloProducer
	orderProducer OrderProducer
}

func NewAPIServer(listenAddr string, hp HelloProducer, op OrderProducer) *APIServer {
	return &APIServer{
		listenAddr:    listenAddr,
		helloProducer: hp,
		orderProducer: op,
	}
}

func (s *APIServer) Run() {
	router := s.createRouter()
	log.Println("JSON API server running on port: ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

// Create a common router that will propogate requests to their relevant microservices
func (s *APIServer) createRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/hello", makeHTTPHandleFunc(s.handleHello))
	router.HandleFunc("/orders", makeHTTPHandleFunc(s.handleOrder))
	return router
}

func (s *APIServer) handleHello(_ http.ResponseWriter, _ *http.Request) error {
	return s.helloProducer.SayHello()
}

func (s *APIServer) handleOrder(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		return s.handleCreateOrder(r)
	} else {
		return fmt.Errorf("method not allowed %s", r.Method)
	}
}

func (s *APIServer) handleCreateOrder(r *http.Request) error {
	createOrderRequest := &OrderRequest{}
	if err := json.NewDecoder(r.Body).Decode(&createOrderRequest); err != nil {
		return err
	}
	// Not the best way to get an intent, but works for now
	createOrderRequest.Intent = rand.IntN(math.MaxInt)
	return s.orderProducer.PlaceOrder(*createOrderRequest)
}

// helper functions
func writeJSON(w http.ResponseWriter, status int, v any) error {
	// Note: header().Add calls need to happen before WriteHeader call
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			//handle the error
			writeJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}
