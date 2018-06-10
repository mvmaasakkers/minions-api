package api

import (
	"net/http"

	"github.com/BeyondBankingDays/minions-api"
)

func (h *Meta) ChallengeListHandler(w http.ResponseWriter, r *http.Request) {
	challenges := hackathon_api.Challenges
	user := getContextUser(r)
	for key, challenge := range challenges {
		if inSlice(challenge.Id, user.Challenges) {
			challenges[key].Done = true
		}
	}
	JsonResponse(w, r, http.StatusOK, challenges)
}
