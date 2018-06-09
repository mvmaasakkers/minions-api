package api

import (
	"net/http"
	"github.com/BeyondBankingDays/minions-api/ext/bb"
)

func (h *Meta) BankGetData(w http.ResponseWriter, r *http.Request) {
	user := getContextUser(r)
	// Mock first bank account user?
	if len(user.BankUsers) == 0 {
		JsonResponse(w, r, http.StatusBadRequest, NewApiError("no bank users connected"))
		return
	}

	bankUser := user.BankUsers[0]
	tokenResponse, err := bb.Login(bankUser)
	if err != nil {
		JsonResponse(w, r, http.StatusBadRequest, NewApiError(err.Error()))
		return
	}

	conn := bb.Conn{Token: tokenResponse.Token}
	accounts, err := conn.GetAccounts()
	if err != nil {
		JsonResponse(w, r, http.StatusBadRequest, NewApiError(err.Error()))
		return
	}

	JsonResponse(w, r, http.StatusOK, accounts)
}
