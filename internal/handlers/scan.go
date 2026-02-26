package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

// ScanRequest defines the exact JSOM structure we expect from the React frontend
type ScanRequest struct {
	TargetURL string `json:"target_url"`
}

// HandleCreateScan processes incoming POST requests to register new URL
func HandleCreateScan(logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//1. Restrict to POST method only
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		//2. Decode the incoming JSON payload
		var req ScanRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			logger.Error("failed to decode request body", "error", err.Error())
			http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
			return
		}

		//3. Basic Validation
		if req.TargetURL == "" {
			logger.Warn("rerjected scan request: missing target_url")
			http.Error(w, "target_url is required", http.StatusBadRequest)
			return
		}

		// TODO: insert database logic here later

		// 4. Send Success Response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{
			"status":     "success",
			"message":    "Scan target registered",
			"target_url": req.TargetURL,
		})

		logger.Info("scan target registered successfully", "url", req.TargetURL)

	}

}
