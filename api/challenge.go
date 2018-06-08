package api

import (
	"net/http"
	"github.com/jumba-nl/hackathon-api"
	"github.com/jumba-nl/hackathon-api/db/mongodb"
	"github.com/gorilla/mux"
)

type ChallengeListHandler struct {
	Meta
	challengeService *mongodb.ChallengeService
}

func (h *ChallengeListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	challenges := make([]*hackathon_api.Challenge, 0)

	challenges = append(challenges, &hackathon_api.Challenge{Name: "test"})
	JsonResponse(w, r, http.StatusOK, challenges)
}

type ChallengeGetHandler struct {
	Meta
	challengeService *mongodb.ChallengeService
}

func (h *ChallengeGetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if _, ok  := vars["id"]; !ok {
		JsonResponse(w, r, http.StatusBadRequest, NewApiError("no id given"))
		return
	}

	challenge := &hackathon_api.Challenge{Name: "blabla"}

	JsonResponse(w, r, http.StatusOK, challenge)
}