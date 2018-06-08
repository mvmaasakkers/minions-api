package api

import (
	"net/http"
	"github.com/jumba-nl/hackathon-api"
	"encoding/json"
	"gopkg.in/validator.v2"
	"github.com/jumba-nl/hackathon-api/db/mongodb"
)

type SourceHandler struct {
	Meta
	sourceService *mongodb.SourceService
}

func (h *SourceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.sourceService = h.DB.NewSourceService()
	switch r.Method {
	case "GET":
		if "" != r.URL.Query().Get("id") {
			h.get(w, r)
			return
		}

		h.list(w, r)
	case "POST":
		h.post(w, r)
	case "OPTIONS":
		JsonResponse(w, r, http.StatusOK, hackathon_api.Source{})
		return
	}
}

func (h *SourceHandler) list(w http.ResponseWriter, r *http.Request) {
	sources, err := h.sourceService.ListSources()
	if err != nil {
		JsonResponse(w, r, http.StatusInternalServerError, NewApiError(err.Error()))
		return
	}

	JsonResponse(w, r, http.StatusOK, sources)
}

func (h *SourceHandler) get(w http.ResponseWriter, r *http.Request) {
	source, err := h.sourceService.GetSource(r.URL.Query().Get("id"))
	if err != nil {
		JsonResponse(w, r, http.StatusInternalServerError, NewApiError(err.Error()))
		return
	}

	JsonResponse(w, r, http.StatusOK, source)
}

func (h *SourceHandler) post(w http.ResponseWriter, r *http.Request) {
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
