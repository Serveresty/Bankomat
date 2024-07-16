package test

import (
	"Bankomat/internal/handlers"
	"Bankomat/internal/services"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	go services.SetupWorkers(1)
	req1, err := http.NewRequest("POST", "/accounts", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr1 := httptest.NewRecorder()
	handlers.CreateAccountHandler(rr1, req1)

	if rr1.Code != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr1.Code, http.StatusCreated)
	}

	if rr1.Body.String() == "" {
		t.Errorf("handler returned unexpected body: got %v want non-empty",
			rr1.Body.String())
	}

	acc, err := services.GetAccount(rr1.Body.String())
	t.Logf("Account : %v", acc)
	if err != nil {
		t.Errorf("ACC not found: %s", err.Error())
	}
}
