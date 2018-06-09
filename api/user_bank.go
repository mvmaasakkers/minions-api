package api

import (
	"net/http"
	"encoding/json"
	"github.com/BeyondBankingDays/minions-api"
	"gopkg.in/validator.v2"
	"github.com/BeyondBankingDays/minions-api/db/mongodb"
)


func (m *Meta) AddBankUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data := hackathon_api.BankUser{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {

	}

	if errs := validator.Validate(data); errs != nil {
		JsonResponse(w, r, http.StatusBadRequest, NewApiError(errs.Error()))
		return
	}

	user := getContextUser(r)

	for _, bu := range user.BankUsers {
		if bu.Username == data.Username {
			JsonResponse(w, r, http.StatusBadRequest, NewApiError("username already connected"))
			return
		}
	}

	user.BankUsers = append(user.BankUsers, data)

	userService := mongodb.NewUserService(&m.DB)
	if _, err := userService.EditUser(user); err != nil {
		JsonResponse(w, r, http.StatusInternalServerError, NewApiError(err.Error()))
		return
	}

	JsonResponse(w, r, http.StatusOK, Received{true, "bank user created"})

}
