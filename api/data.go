package api

import (
	"net/http"
	"github.com/BeyondBankingDays/minions-api"
	"encoding/json"
	"time"
	"gopkg.in/validator.v2"
)


type DataHandler struct {
	Meta
}

func (h *DataHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		JsonResponse(w, r, http.StatusOK, hackathon_api.Data{})
		return
	}
	defer r.Body.Close()
	data := &hackathon_api.Data{}

	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		JsonResponse(w, r, http.StatusInternalServerError, NewApiError(err.Error()))
		return
	}

	data.Time = time.Now()

	if errs := validator.Validate(data); errs != nil {
		JsonResponse(w, r, http.StatusBadRequest, NewApiError(errs.Error()))
		return
	}

	dataService := h.DB.NewDataService()
	if _, err := dataService.CreateData(data); err != nil {
		JsonResponse(w, r, http.StatusInternalServerError, NewApiError(err.Error()))
		return
	}

	JsonResponse(w, r, http.StatusCreated, Received{true, "data received"})
}

type Received struct {
	Success bool `json:"success"`
	Message string `json:"message"`
}