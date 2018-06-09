package api

import (
	"net/http"
	"github.com/gorilla/mux"
	"io"
)

func PropertiesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if _, ok  := vars["id"]; !ok {
		JsonResponse(w, r, http.StatusBadRequest, NewApiError("no id given"))
		return
	}

	client := http.Client{}

	req, err := http.NewRequest("GET", "https://jumba.nl/v1/properties/"+vars["id"], nil)
	if err != nil {
		JsonResponse(w, r, http.StatusInternalServerError, NewApiError(err.Error()))
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		JsonResponse(w, r, http.StatusInternalServerError, NewApiError(err.Error()))
		return
	}
	defer resp.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

