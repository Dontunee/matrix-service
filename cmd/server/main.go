package main

import (
	"log"
	"net/http"

	"github.com/Dontunee/matrix-service/internal/routes"
)

func main() {
	r := routes.SetupRoutes()
	log.Println("Starting Matrix service on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Failed to start service: %v", err)
	}
}
