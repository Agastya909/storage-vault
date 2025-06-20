package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	cfg, err := LoadConfig()
	if err != nil {
		fmt.Println("Error loading configuration:", err)
		return
	}

	Handlers, err := SetupHandlers(cfg)
	if err != nil {
		fmt.Println("Error setting up handlers:", err)
		return
	}

	r := chi.NewRouter()
	InitRoutes(r, Handlers)

	port := fmt.Sprintf(":%s", cfg.Port)
	if err := http.ListenAndServe(port, r); err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
