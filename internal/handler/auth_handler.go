package handler

import (
	"encoding/json"
	"net/http"

	"github.com/r6rap/Quanta/internal/domain"
	"github.com/r6rap/Quanta/internal/repository"
	"github.com/r6rap/Quanta/internal/services"
	"github.com/r6rap/Quanta/pkg/database"
	"github.com/r6rap/Quanta/pkg/response"
)

type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHandler(db *database.DB) *AuthHandler {
	userRepo := repository.NewUserRepository(db)
	refreshRepo := repository.NewRefreshRepository(db)

	authService := services.NewAuthService(userRepo, refreshRepo)

	return &AuthHandler {service: authService}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req domain.RegisterInput

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.Fail(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Invalid Request", err.Error())
		return
	}

	user, err := h.service.Register(&req)
	if err != nil {
		response.Fail(w, http.StatusBadRequest, "FAILED_REGIST", "Bad Request", err.Error())
		return
	}

	response.OK(w, http.StatusCreated, user)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req domain.LoginInput

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.Fail(w, http.StatusBadRequest, "INTERNAL_ERROR", "Invalid Request", err.Error())
		return
	}

	user, err := h.service.Login(&req)
	if  err != nil {
		response.Fail(w, http.StatusUnauthorized, "UNAUTHORIZED", "Failed login", err.Error())
		return
	}

	response.OK(w, http.StatusOK, user)
}

func (h *AuthHandler) RefreshAccessToken(w http.ResponseWriter, r *http.Request) {
	var req domain.RefreshInput

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.Fail(w, http.StatusBadRequest, "INTERNAL_ERROR", "Invalid Request", err.Error())
		return
	}

	accessToken, err := h.service.RefreshAccessToken(&req)
	if err != nil {
		response.Fail(w, http.StatusUnauthorized, "UNAUTHORIZED", "Invalid Refresh Token", err.Error())
		return
	}

	response.OK(w, http.StatusOK, accessToken)
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	var req domain.LogoutInput

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.Fail(w, http.StatusBadRequest, "INTERNAL_ERROR", "Invalid Request", err.Error())
		return
	}

	h.service.Logout(&req)
	
	response.OK(w, http.StatusOK, "Logout successfully")
}