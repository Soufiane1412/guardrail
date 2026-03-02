package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/soufiane1412/guardrail/internal/handlers"
	"github.com/soufiane1412/guardrail/internal/logger"
)

func main() {
	// 1. Initialise the Core Systems
	log := logger.New() // Booting up the JSON logger we built
	log.Info("initialising Guardrail Sentinel engine...")

	// 2. Initialise the Router (Multiplexer)
	mux := http.NewServeMux()

	// 3. Register the Endpoints
	mux.Handle("POST /api/v1/scans", handlers.HandleCreateScan(log))

	// 4. Configure the Server Port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// 5. Ignite the server
	addr := fmt.Sprintf(":%s", port)
	log.Info("server active and listening", "port", port)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Error("fatal server crash", "error", err.Error())
		os.Exit(1)
	}

}
