package handlers

import (
	"Bankomat/internal/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func RegisterHandlers(r *mux.Router) {
	r.HandleFunc("/accounts", CreateAccountHandler).Methods("POST")
	r.HandleFunc("/accounts/{id}/deposit", DepositHandler).Methods("POST")
	r.HandleFunc("/accounts/{id}/withdraw", WithdrawHandler).Methods("POST")
	r.HandleFunc("/accounts/{id}/balance", BalanceHandler).Methods("GET")
}

func CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	id := services.CreateAccount()
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(id))
}

func DepositHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	amount, err := strconv.ParseFloat(r.FormValue("amount"), 64)
	if err != nil {
		http.Error(w, "Invalid amount", http.StatusBadRequest)
		return
	}
	err = services.Deposit(id, amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func WithdrawHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	amount, err := strconv.ParseFloat(r.FormValue("amount"), 64)
	if err != nil {
		http.Error(w, "Invalid amount", http.StatusBadRequest)
		return
	}

	err = services.Withdraw(id, amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func BalanceHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	balance := services.GetBalance(id)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Balance: %f", balance)))
}
