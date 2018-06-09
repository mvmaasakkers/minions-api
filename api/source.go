package api

import (
	"net/http"
	"github.com/BeyondBankingDays/minions-api"
	"encoding/json"
	"gopkg.in/validator.v2"
	"github.com/BeyondBankingDays/minions-api/db/mongodb"
	"github.com/gorilla/mux"
)

type SourceListHandler struct {
	Meta
	sourceService *mongodb.SourceService
}

func (h *SourceListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, err := AuthHeader(r.Header.Get("Authorization"))
	if user == nil {
		JsonResponse(w, r, http.StatusForbidden, NewApiError(err.Error()))
		return
	}
	h.sourceService = h.DB.NewSourceService()
	sources, err := h.sourceService.ListSources()
	if err != nil {
		JsonResponse(w, r, http.StatusInternalServerError, NewApiError(err.Error()))
		return
	}

	JsonResponse(w, r, http.StatusOK, sources)
}

type SourceGetHandler struct {
	Meta
	sourceService *mongodb.SourceService
}

func (h *SourceGetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, err := AuthHeader(r.Header.Get("Authorization"))
	if user == nil {
		JsonResponse(w, r, http.StatusForbidden, NewApiError(err.Error()))
		return
	}
	h.sourceService = h.DB.NewSourceService()
	vars := mux.Vars(r)
	if _, ok  := vars["id"]; !ok {
		JsonResponse(w, r, http.StatusBadRequest, NewApiError("no id given"))
		return
	}

	source, err := h.sourceService.GetSource(vars["id"])
	if err != nil {
		JsonResponse(w, r, http.StatusInternalServerError, NewApiError(err.Error()))
		return
	}

	JsonResponse(w, r, http.StatusOK, source)
}

type SourcePostHandler struct {
	Meta
	sourceService *mongodb.SourceService
}

func (h *SourcePostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, err := AuthHeader(r.Header.Get("Authorization"))
	if user == nil {
		JsonResponse(w, r, http.StatusForbidden, NewApiError(err.Error()))
		return
	}
	h.sourceService = h.DB.NewSourceService()
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

	sourceService := h.DB.NewSourceService()
	createdSource, err := sourceService.CreateSource(source)
	if err != nil {
		JsonResponse(w, r, http.StatusInternalServerError, NewApiError(err.Error()))
		return
	}

	JsonResponse(w, r, http.StatusCreated, createdSource)
}
