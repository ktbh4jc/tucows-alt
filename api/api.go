package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr    string
	helloProducer HelloProducer
}

func NewAPIServer(listenAddr string, helloProducer HelloProducer) *APIServer {
	return &APIServer{
		listenAddr:    listenAddr,
		helloProducer: helloProducer,
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
	return router
}

func (s *APIServer) handleHello(_ http.ResponseWriter, _ *http.Request) error {
	return s.helloProducer.SayHello()
}

// helper functions
func WriteJSON(w http.ResponseWriter, status int, v any) error {
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
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}
