package api

import (
	"net/http"
	"strings"
	"errors"
	"github.com/BeyondBankingDays/minions-api/db/mongodb"
	"github.com/BeyondBankingDays/minions-api"
	"github.com/BeyondBankingDays/minions-api/ext/gorctx"
)

func getContextUser(r *http.Request) *hackathon_api.User {
	if v := gorctx.Get(r, "user"); v != nil {
		if curUser, ok := v.(*hackathon_api.User); ok {
			return curUser
		}
	}
	return nil
}


func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := AuthHeader(r.Header.Get("Authorization"))
		if err != nil {
			JsonResponse(w, r, http.StatusUnauthorized, ApiError{err.Error()})
			return
		}

		gorctx.Set(r, "user", user)

		next.ServeHTTP(w, r)
	}
}

func AuthHeader(authHeader string) (*hackathon_api.User, error) {
	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 {
		return nil, errors.New("no token given")
	}

	return AuthToken(authHeaderParts[1])
}

func AuthToken(tokenVal string)  (*hackathon_api.User, error) {
	if tokenVal == "" {
		return nil, errors.New("no token given")
	}

	tokenService := mongodb.NewTokenService(mongodb.Conn)
	token, err := tokenService.GetTokenByToken(tokenVal)
	if err != nil {
		return nil, err
	}

	userService := mongodb.NewUserService(mongodb.Conn)
	user, err := userService.GetUser(token.UserId)
	if err != nil {
		return nil, err
	}

	return user, nil
}
