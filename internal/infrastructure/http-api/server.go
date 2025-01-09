package http_api

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func NewServer() *http.Server {
	router := mux.NewRouter()

	router.HandleFunc("/health", healthHandler).
		Methods(http.MethodGet)

	return &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
}

func Start(s *http.Server) {
	log.Printf("Starting server on %v", s.Addr)
	if err := http.ListenAndServe(s.Addr, s.Handler); err != nil {
		panic(err)
	}
}

func Shutdown(s *http.Server) {
	log.Printf("Shutting down server")
	if err := s.Shutdown(nil); err != nil {
		panic(err)
	}
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("OK"))
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
}
