package api

import (
	"net/http"
	"github.com/BeyondBankingDays/minions-api"
)



func (h *Meta) ChallengeListHandler(w http.ResponseWriter, r *http.Request) {

	JsonResponse(w, r, http.StatusOK, hackathon_api.Challenges)
}
