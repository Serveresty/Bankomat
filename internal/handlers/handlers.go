package handlers

import (
	"Bankomat/internal/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func RegisterHandlers(r *mux.Router) {
	r.HandleFunc("/accounts", createAccountHandler).Methods("POST")
	r.HandleFunc("/accounts/{id}/deposit", depositHandler).Methods("POST")
	r.HandleFunc("/accounts/{id}/withdraw", withdrawHandler).Methods("POST")
	r.HandleFunc("/accounts/{id}/balance", balanceHandler).Methods("GET")
}

func createAccountHandler(w http.ResponseWriter, r *http.Request) {
	id := services.CreateAccount()
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(id))
}

func depositHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
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

func withdrawHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
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

func balanceHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	balance := services.GetBalance(id)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Balance: %f", balance)))
}
