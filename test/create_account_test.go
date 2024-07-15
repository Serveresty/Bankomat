package test

import (
	"Bankomat/internal/handlers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestCreateAccount(t *testing.T) {
	r := mux.NewRouter()
	handlers.RegisterHandlers(r)

	req, err := http.NewRequest("POST", "/accounts", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	if rr.Body.String() == "" {
		t.Errorf("handler returned unexpected body: got %v want non-empty",
			rr.Body.String())
	}
}
