package api

import (
	"net/http"
	"encoding/json"
	"log"
)

func JsonResponse(w http.ResponseWriter, r *http.Request, statusCode int, body interface{}) {
	if statusCode == 0 {
		statusCode = http.StatusOK
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(body); err != nil  {
		log.Println("Could not decode json body", body, err)
	}
}
