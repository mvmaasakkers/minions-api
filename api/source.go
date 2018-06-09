package api

import (
	"net/http"
	"github.com/BeyondBankingDays/minions-api"
	"encoding/json"
	"gopkg.in/validator.v2"
	"github.com/gorilla/mux"
)


func (h *Meta) SourceListHandler(w http.ResponseWriter, r *http.Request) {
	sourceService := h.DB.NewSourceService()
	sources, err := sourceService.ListSources()
	if err != nil {
		JsonResponse(w, r, http.StatusInternalServerError, NewApiError(err.Error()))
		return
	}

	JsonResponse(w, r, http.StatusOK, sources)
}


func (h *Meta) SourceGetHandler(w http.ResponseWriter, r *http.Request) {
	sourceService := h.DB.NewSourceService()
	vars := mux.Vars(r)
	if _, ok  := vars["id"]; !ok {
		JsonResponse(w, r, http.StatusBadRequest, NewApiError("no id given"))
		return
	}

	source, err := sourceService.GetSource(vars["id"])
	if err != nil {
		JsonResponse(w, r, http.StatusInternalServerError, NewApiError(err.Error()))
		return
	}

	JsonResponse(w, r, http.StatusOK, source)
}


func (h *Meta) SourcePostHandler(w http.ResponseWriter, r *http.Request) {
	sourceService := h.DB.NewSourceService()
	defer r.Body.Close()
	source := &hackathon_api.Source{}

	if err := json.NewDecoder(r.Body).Decode(source); err != nil {
		JsonResponse(w, r, http.StatusInternalServerError, NewApiError(err.Error()))
		return
	}

	if errs := validator.Validate(source); errs != nil {
		JsonResponse(w, r, http.StatusBadRequest, NewApiError(errs.Error()))
		return
	}

	createdSource, err := sourceService.CreateSource(source)
	if err != nil {
		JsonResponse(w, r, http.StatusInternalServerError, NewApiError(err.Error()))
		return
	}

	JsonResponse(w, r, http.StatusCreated, createdSource)
}
