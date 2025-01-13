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

	router.HandleFunc("/accounts/{account_id}/deposit", func(w http.ResponseWriter, r *http.Request) {
		depositHandler(w, r, b)
	}).
		Methods(http.MethodPost)

	router.HandleFunc("/accounts/{account_id}/withdraw", func(w http.ResponseWriter, r *http.Request) {
		withdrawHandler(w, r, b)
	}).
		Methods(http.MethodPost)

	router.HandleFunc("/accounts/{account_id}/balance", func(w http.ResponseWriter, r *http.Request) {
		getBalanceHandler(w, r, b)
	}).Methods(http.MethodGet)

	return router
}

func Start(b *internal.Bank, address string) {
	r := Handler(b)
	s := &http.Server{
		Addr:    address,
		Handler: r,
	}
	if err := http.ListenAndServe(s.Addr, s.Handler); err != nil {
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

func openAccountHandler(w http.ResponseWriter, r *http.Request, b *internal.Bank) {
	var requestBody struct {
		Holder   string `json:"holder"`
		Currency string `json:"currency"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("Invalid request body"))
		if err != nil {
			log.Printf("Error writing response: %v", err)
		}
		return
	}

	id, err := b.OpenAccount(requestBody.Holder, requestBody.Currency)
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

func depositHandler(w http.ResponseWriter, r *http.Request, b *internal.Bank) {
	vars := mux.Vars(r)
	accountId := vars["account_id"]

	var requestBody struct {
		Amount   float64 `json:"amount"`
		Currency string  `json:"currency"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("Invalid request body"))
		if err != nil {
			log.Printf("Error writing response: %v", err)
		}
		return
	}

	if err := b.Deposit(accountId, requestBody.Amount, requestBody.Currency); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Printf("Error writing response: %v", err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err2 := w.Write([]byte("OK"))
	if err2 != nil {
		log.Printf("Error writing response: %v", err2)
	}
}

func withdrawHandler(w http.ResponseWriter, r *http.Request, b *internal.Bank) {
	vars := mux.Vars(r)
	accountId := vars["account_id"]

	var requestBody struct {
		Amount   float64 `json:"amount"`
		Currency string  `json:"currency"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("Invalid request body"))
		if err != nil {
			log.Printf("Error writing response: %v", err)
		}
		return
	}

	if err := b.Withdraw(accountId, requestBody.Amount, requestBody.Currency); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Printf("Error writing response: %v", err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err2 := w.Write([]byte("OK"))
	if err2 != nil {
		log.Printf("Error writing response: %v", err2)
	}
}

func getBalanceHandler(w http.ResponseWriter, r *http.Request, b *internal.Bank) {
	vars := mux.Vars(r)
	accountId := vars["account_id"]

	balance, err := b.GetBalance(accountId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Printf("Error writing response: %v", err)
		}
		return
	}

	responseBody := struct {
		Balance  float64 `json:"balance"`
		Currency string  `json:"currency"`
	}{
		Balance:  balance.Amount,
		Currency: balance.Currency,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(responseBody); err != nil {
		log.Printf("Error writing response: %v", err)
	}
}
