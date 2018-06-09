package api

import (
	"net/http"
	"strings"
	"errors"
	"github.com/BeyondBankingDays/minions-api/db/mongodb"
	"github.com/BeyondBankingDays/minions-api"
	"context"
)


func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := AuthHeader(r.Header.Get("Authorization"))
		if err != nil {
			JsonResponse(w, r, http.StatusUnauthorized, ApiError{err.Error()})
			return
		}

		c := context.Background()

		r.WithContext(context.WithValue(c, "user", user))

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

	tokenService := mongodb.Conn.NewTokenService()
	token, err := tokenService.GetTokenByToken(tokenVal)
	if err != nil {
		return nil, err
	}

	userService := mongodb.Conn.NewUserService()
	user, err := userService.GetUser(token.UserId)
	if err != nil {
		return nil, err
	}

	return user, nil
}
