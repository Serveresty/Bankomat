package main

import (
	"Bankomat/internal/handlers"
	"Bankomat/internal/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	countWorkers int = 1
)

func main() {
	go services.SetupWorkers(countWorkers)

	r := mux.NewRouter()
	handlers.RegisterHandlers(r)
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Error starting server: %s", err.Error())
	}
}
