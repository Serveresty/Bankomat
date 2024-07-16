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

func TestDeposit(t *testing.T) {
	go services.SetupWorkers(1)
	accountID := services.CreateAccount()

	data := url.Values{}
	data.Set("amount", "123.45")

	url := "/accounts/" + accountID + "/deposit"

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

	acc, _ := services.GetAccount(accountID)
	t.Logf("ACCOUNT DETAILS: %v", acc)
}
