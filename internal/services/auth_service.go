package services

import (
	"os"
	"time"

	"github.com/r6rap/Quanta/internal/domain"
	"github.com/r6rap/Quanta/internal/repository"
	"github.com/r6rap/Quanta/pkg/jwt"

	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repository.UserRepository
	refreshRepo *repository.RefreshRepository
}

func NewAuthService(userRepo *repository.UserRepository, refreshRepo *repository.RefreshRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		refreshRepo: refreshRepo,
	}
}

func (s *AuthService) Register(req *domain.RegisterInput) (*domain.User, error) {
	existing, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if existing != nil {
		return nil, errors.New("Email already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Email:       req.Email,
		PasswordHash: string(hash),
		Name:        req.Name,
	}

	if err = s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) Login(req *domain.LoginInput) (*domain.LoginResponse, error) {
	user, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("Invalid crecedentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return nil, errors.New("Invalid crecedentials")
	}

	accessToken, _ := jwt.GenerateAccessToken(user.ID, user.Email)
	refreshToken, _ := jwt.GenerateRefreshToken(user.ID, user.Email)

	expiresStr := os.Getenv("JWT_REFRESH_EXPIRY")

	expiresAt, _ := time.ParseDuration(expiresStr)

	payload := &domain.RefreshToken {
		UserID: user.ID,
		Token: refreshToken,
		ExpiresAt: time.Now().Add(expiresAt),
	}

	s.refreshRepo.Create(payload)

	return &domain.LoginResponse{
		AccessToken: accessToken,
		User: *user,
	}, nil
}

func (s *AuthService) RefreshAccessToken(req *domain.RefreshInput) (string, error) {
	claims, err := jwt.ValidateToken(req.RefreshToken)
	if err != nil {
		return "", err
	}

	existingToken, err := s.refreshRepo.GetByToken(req.RefreshToken, claims.UserID)
	if err != nil {
		return "", err
	}
	
	if existingToken == nil {
		return "", errors.New("Token not found")
	}

	if time.Now().After(existingToken.ExpiresAt) {
		return "", s.refreshRepo.Delete(req.RefreshToken, claims.UserID)
	}

	accessToken, err := jwt.GenerateAccessToken(claims.UserID, claims.Email)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (s *AuthService) Logout(req *domain.LogoutInput) error {
	return s.refreshRepo.Delete(req.RefreshToken, req.UserID)
}