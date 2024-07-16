package test

import (
	"Bankomat/internal/handlers"
	"Bankomat/internal/services"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gorilla/mux"
)

func TestWithdraw(t *testing.T) {
	go services.SetupWorkers(1)
	accountID := services.CreateAccount()
	err := services.Deposit(accountID, 100)
	if err != nil {
		t.Errorf("Deposit error: %v", err)
	}

	acc, _ := services.GetAccount(accountID)
	t.Logf("ACCOUNT DETAILS AFTER DEPOSIT: %v", acc)

	data := url.Values{}
	data.Set("amount", "50")

	url := "/accounts/" + accountID + "/withdraw"

	req2, err := http.NewRequest("POST", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	req2.URL.RawQuery = data.Encode()

	rr2 := httptest.NewRecorder()

	r := mux.NewRouter()
	handlers.RegisterHandlers(r)
	r.ServeHTTP(rr2, req2)

	if rr2.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr2.Code, http.StatusOK)
	}

	acc, _ = services.GetAccount(accountID)
	t.Logf("ACCOUNT DETAILS AFTER WITHDRAW: %v", acc)
}
