package api

import (
	"github.com/BeyondBankingDays/minions-api/db/mongodb"
	"net/http"
	"encoding/json"
	"github.com/BeyondBankingDays/minions-api"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/validator.v2"
	"github.com/gorilla/mux"
)

func (h *Meta) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	user := getContextUser(r)
	user.Password = ""

	JsonResponse(w, r, http.StatusOK, user)
}

type UserPayBody struct {
	Points int `json:"points" validate:"nonzero"`
}


func (h *Meta) UserPayHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	userBody := UserPayBody{}
	if err := json.NewDecoder(r.Body).Decode(&userBody); err != nil {
		JsonResponse(w, r, http.StatusBadRequest, NewApiError(err.Error()))
		return
	}

	if errs := validator.Validate(userBody); errs != nil {
		JsonResponse(w, r, http.StatusBadRequest, NewApiError(errs.Error()))
		return
	}

	user := getContextUser(r)
	if userBody.Points > user.Score.Current {
		JsonResponse(w, r, http.StatusConflict, NewApiError("not enough points to spend"))
		return
	}

	user.Score.Current -= userBody.Points

	userService := mongodb.NewUserService(&h.DB)
	if _, err := userService.EditUser(user); err != nil {
		JsonResponse(w, r, http.StatusInternalServerError, NewApiError("something went wrong" + err.Error()))
		return
	}

	JsonResponse(w, r, http.StatusOK, Received{true, "points spent"})
}


func (h *Meta) UserDoChallengeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if _, ok  := vars["id"]; !ok {
		JsonResponse(w, r, http.StatusBadRequest, NewApiError("no id given"))
		return
	}

	challenge, err := hackathon_api.GetChallenge(vars["id"])
	if err != nil {
		JsonResponse(w, r, http.StatusBadRequest, NewApiError("challenge not found"))
		return
	}

	user := getContextUser(r)
	if inSlice(vars["id"], user.Challenges) {
		JsonResponse(w, r, http.StatusBadRequest, NewApiError("user already completed this challenge"))
		return
	}

	user.Challenges = append(user.Challenges, vars["id"])
	user.Score.Current += challenge.Points

	userService := mongodb.NewUserService(&h.DB)
	if _, err := userService.EditUser(user); err != nil {
		JsonResponse(w, r, http.StatusInternalServerError, NewApiError("something went wrong" + err.Error()))
		return
	}

	JsonResponse(w, r, http.StatusOK, Received{true, "challenge done"})
}

func inSlice(needle string, haystack []string) bool {
	for _, val := range haystack {
		if val == needle {
			return true
		}
	}

	return false
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

	h.userService = mongodb.NewUserService(&h.DB)
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

	h.userService = mongodb.NewUserService(&h.DB)
	user, err := h.userService.GetUserByUsername(userBody.Username)
	if err != nil {
		JsonResponse(w, r, http.StatusInternalServerError, NewApiError(err.Error()))
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userBody.Password)); err != nil {
		JsonResponse(w, r, http.StatusUnauthorized, NewApiError(err.Error()))
		return
	}

	h.tokenService = mongodb.NewTokenService(&h.DB)
	token := &hackathon_api.Token{UserId: user.Id.Hex()}
	if _, err := h.tokenService.CreateToken(token); err != nil {
		JsonResponse(w, r, http.StatusUnauthorized, NewApiError(err.Error()))
		return
	}

	JsonResponse(w, r, http.StatusOK, token)
}
