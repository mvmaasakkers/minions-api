package api

import (
	"net/http"
	"github.com/BeyondBankingDays/minions-api"
	"encoding/json"
	"time"
	"gopkg.in/validator.v2"
	"github.com/BeyondBankingDays/minions-api/db/mongodb"
)


type DataHandler struct {
	Meta
}

func (h *Meta) DataHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		JsonResponse(w, r, http.StatusOK, hackathon_api.Data{})
		return
	}

	if r.Body == nil {
		JsonResponse(w, r, http.StatusBadRequest, NewApiError("no body given"))
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

	dataService := mongodb.NewDataService(&h.DB)
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