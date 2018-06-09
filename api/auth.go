package api

import (
	"net/http"
	"strings"
	"errors"
	"github.com/BeyondBankingDays/minions-api/db/mongodb"
	"github.com/BeyondBankingDays/minions-api"
	"context"
)

// Set assigns value v under key k on given Request r's context.
func ContextSet(r *http.Request, k, v interface{}) {
	if v == nil {
		return
	}
	*r = *r.WithContext(context.WithValue(r.Context(), k, v))
}

// Get retrieves value registered under key k of given Request context.
func ContextGet(r *http.Request, k interface{}) interface{} {
	return r.Context().Value(k)
}

// GetOk retrieves value of key k from the given Request and indicates success or
// failure in 2nd return value.
func ContextGetOk(r *http.Request, k interface{}) (v interface{}, ok bool) {
	if v = ContextGet(r, k); v != nil {
		ok = true
	}
	return
}


func getContextUser(r *http.Request) *hackathon_api.User {
	if v := ContextGet(r, "user"); v != nil {
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

		ContextSet(r, "user", user)

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
		return nil, errors.New("token "+ err.Error())
	}

	userService := mongodb.NewUserService(mongodb.Conn)
	user, err := userService.GetUser(token.UserId)
	if err != nil {
		return nil, errors.New("user "+ err.Error())
	}

	return user, nil
}
