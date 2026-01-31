package handler

import (
	"encoding/json"
	"net/http"
	"time"

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

	return &AuthHandler{service: authService}
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

	http.SetCookie(w, &http.Cookie{
		Name: "refresh_token",
		Value: user.RefreshToken,
		Expires: time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure: true,
		SameSite: http.SameSiteStrictMode,
		Path: "/",
		Domain: "",
	})

	response.OK(w, http.StatusCreated, domain.RegisterResponse{
		AccessToken: user.AccessToken,
		User:        user.User,
	})
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req domain.LoginInput

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.Fail(w, http.StatusBadRequest, "INTERNAL_ERROR", "Invalid Request", err.Error())
		return
	}

	user, err := h.service.Login(&req)
	if err != nil {
		response.Fail(w, http.StatusUnauthorized, "UNAUTHORIZED", "Failed login", err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name: "refresh_token",
		Value: user.RefreshToken,
		Expires: time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure: true,
		SameSite: http.SameSiteStrictMode,
		Path: "/",
		Domain: "",
	})

	response.OK(w, http.StatusOK, domain.LoginResponse{
		AccessToken: user.AccessToken,
		User:        user.User,
	})
}

func (h *AuthHandler) RefreshAccessToken(w http.ResponseWriter, r *http.Request) {
	refreshToken, err := r.Cookie("refresh_token")
	if err != nil {
		response.Fail(w, http.StatusBadRequest, "INTERNAL_ERROR", "Invalid Request", err.Error())
		return
	}

	accessToken, err := h.service.RefreshAccessToken(&domain.RefreshInput{
		RefreshToken: refreshToken.Value,
	})
	if err != nil {
		response.Fail(w, http.StatusUnauthorized, "UNAUTHORIZED", "Invalid Refresh Token", err.Error())
		return
	}

	response.OK(w, http.StatusOK, accessToken)
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	refreshToken, err := r.Cookie("refresh_token")
	if err != nil {
		response.Fail(w, http.StatusBadRequest, "INTERNAL_ERROR", "Invalid Request", err.Error())
		return
	}

	err = h.service.Logout(&domain.LogoutInput{
		RefreshToken: refreshToken.Value,
	})
	if err != nil {
		response.Fail(w, http.StatusUnauthorized, "UNAUTHORIZED", "Invalid Refresh Token", err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name: "refresh_token",
		Value: "",
		Expires: time.Now().Add(-24 * time.Hour),
		HttpOnly: true,
		Secure: true,
		SameSite: http.SameSiteStrictMode,
		Path: "/",
		Domain: "",
	})

	response.OK(w, http.StatusOK, "Logout successfully")
}
