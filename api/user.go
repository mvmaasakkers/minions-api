package api

import (
	"github.com/BeyondBankingDays/minions-api/db/mongodb"
	"net/http"
	"encoding/json"
	"github.com/BeyondBankingDays/minions-api"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/validator.v2"
)

type GetUserHandler struct {
	Meta
	userService *mongodb.UserService
}

func (h *GetUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, err := AuthHeader(r.Header.Get("Authorization"))
	if user == nil {
		JsonResponse(w, r, http.StatusForbidden, NewApiError(err.Error()))
		return
	}

	user.Password = ""
	JsonResponse(w, r, http.StatusOK, user)
}

type CreateUserHandler struct {
	Meta
	userService *mongodb.UserService
}

type UserRequestBody struct {
	Username string `json:"username" validate:"nonzero"`
	Password string `json:"password" validate:"nonzero"`
}

func (h *CreateUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	userBody := UserRequestBody{}
	if err := json.NewDecoder(r.Body).Decode(&userBody); err != nil {
		JsonResponse(w, r, http.StatusBadRequest, NewApiError(err.Error()))
		return
	}

	if errs := validator.Validate(userBody); errs != nil {
		JsonResponse(w, r, http.StatusBadRequest, NewApiError(errs.Error()))
		return
	}

	h.userService = h.DB.NewUserService()
	_, err := h.userService.GetUserByUsername(userBody.Username)
	if err == nil {
		JsonResponse(w, r, http.StatusInternalServerError, NewApiError("username already in use"))
		return
	}

	pw, err := bcrypt.GenerateFromPassword([]byte(userBody.Password), 10)
	if err != nil {
		JsonResponse(w, r, http.StatusInternalServerError, NewApiError(err.Error()))
		return
	}

	user := &hackathon_api.User{
		Username: userBody.Username,
		Password: string(pw),
	}

	if _, err := h.userService.CreateUser(user); err != nil {
		JsonResponse(w, r, http.StatusInternalServerError, NewApiError(err.Error()))
		return
	}

	JsonResponse(w, r, http.StatusOK, Received{true, "user created"})
}

type LoginHandler struct {
	Meta
	userService  *mongodb.UserService
	tokenService *mongodb.TokenService
}

func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	userBody := UserRequestBody{}
	if err := json.NewDecoder(r.Body).Decode(&userBody); err != nil {
		JsonResponse(w, r, http.StatusBadRequest, NewApiError(err.Error()))
		return
	}

	if errs := validator.Validate(userBody); errs != nil {
		JsonResponse(w, r, http.StatusBadRequest, NewApiError(errs.Error()))
		return
	}

	h.userService = h.DB.NewUserService()
	user, err := h.userService.GetUserByUsername(userBody.Username)
	if err != nil {
		JsonResponse(w, r, http.StatusInternalServerError, NewApiError(err.Error()))
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userBody.Password)); err != nil {
		JsonResponse(w, r, http.StatusUnauthorized, NewApiError(err.Error()))
		return
	}

	h.tokenService = h.DB.NewTokenService()
	token := &hackathon_api.Token{UserId: user.Id.Hex()}
	if _, err := h.tokenService.CreateToken(token); err != nil {
		JsonResponse(w, r, http.StatusUnauthorized, NewApiError(err.Error()))
		return
	}

	JsonResponse(w, r, http.StatusOK, token)
}
