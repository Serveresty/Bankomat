package test

import (
	"Bankomat/internal/handlers"
	"Bankomat/internal/services"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetBalance(t *testing.T) {
	go services.SetupWorkers(1)
	accountID := services.CreateAccount()

	req2, err := http.NewRequest("GET", "/accounts/"+accountID+"/balance", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr2 := httptest.NewRecorder()
	r := mux.NewRouter()
	handlers.RegisterHandlers(r)
	r.ServeHTTP(rr2, req2)

	if rr2.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr2.Code, http.StatusOK)
	}

	t.Log(rr2.Body.String())
}
