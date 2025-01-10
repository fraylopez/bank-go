package http_api

import (
	"encoding/json"
	"github.com/fraylopez/bank-go/internal"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Handler(b *internal.Bank) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/health", healthHandler).
		Methods(http.MethodGet)

	router.HandleFunc("/accounts", func(w http.ResponseWriter, r *http.Request) {
		openAccountHandler(w, r, b)
	}).
		Methods(http.MethodPost)

	return router
}

func Start(b *internal.Bank) {
	r := Handler(b)
	s := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	if err := http.ListenAndServe(s.Addr, s.Handler); err != nil {
		panic(err)
	}
	log.Printf("Starting server on %v", s.Addr)
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("OK"))
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

func openAccountHandler(w http.ResponseWriter, _ *http.Request, b *internal.Bank) {
	id, err := b.OpenAccount()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Printf("Error writing response: %v", err)
		}
		return
	}

	responseBody := struct {
		Id string `json:"account_id"`
	}{
		Id: id,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(responseBody); err != nil {
		log.Printf("Error writing response: %v", err)
	}

}
