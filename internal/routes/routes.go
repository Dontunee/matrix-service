package routes

import (
	"net/http"

	"github.com/Dontunee/matrix-service/internal/handlers"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/echo", handlers.EchoHandler)
	mux.HandleFunc("/invert", handlers.InvertHandler)
	mux.HandleFunc("/flatten", handlers.FlattenHandler)
	mux.HandleFunc("/sum", handlers.SumHandler)
	mux.HandleFunc("/multiply", handlers.MultiplyHandler)
	return mux
}
