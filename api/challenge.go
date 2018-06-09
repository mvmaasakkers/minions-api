package api

import (
	"net/http"
	"github.com/BeyondBankingDays/minions-api"
	"github.com/gorilla/mux"
)


func (h *Meta) ChallengeListHandler(w http.ResponseWriter, r *http.Request) {
	challenges := make([]*hackathon_api.Challenge, 0)

	challenges = append(challenges, &hackathon_api.Challenge{Name: "test"})
	JsonResponse(w, r, http.StatusOK, challenges)
}

func (h *Meta) ChallengeGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if _, ok  := vars["id"]; !ok {
		JsonResponse(w, r, http.StatusBadRequest, NewApiError("no id given"))
		return
	}

	challenge := &hackathon_api.Challenge{Name: "blabla"}

	JsonResponse(w, r, http.StatusOK, challenge)
}