package services

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/r6rap/Quanta/internal/domain"
	"github.com/r6rap/Quanta/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

type fakeUserRepo struct {
	users     map[string]*domain.User
	getErr    error
	createErr error
}

func (f *fakeUserRepo) GetByEmail(email string) (*domain.User, error) {
	if f.getErr != nil {
		return nil, f.getErr
	}
	return f.users[email], nil
}

func (f *fakeUserRepo) Create(user *domain.User) error {
	if f.createErr != nil {
		return f.createErr
	}
	user.ID = uint(len(f.users) + 1)
	f.users[user.Email] = user
	return nil
}

type fakeRefreshRepo struct {
	tokens    map[string]*domain.RefreshToken
	createErr error
	getErr    error
	deleteErr error
}

func (f *fakeRefreshRepo) Create(req *domain.RefreshToken) error {
	if f.createErr != nil {
		return f.createErr
	}
	f.tokens[req.Token] = req
	return nil
}

func (f *fakeRefreshRepo) GetByToken(token string, userId uint) (*domain.RefreshToken, error) {
	if f.getErr != nil {
		return nil, f.getErr
	}
	stored := f.tokens[token]
	if stored == nil || stored.UserID != userId {
		return nil, nil
	}
	return stored, nil
}

func (f *fakeRefreshRepo) Delete(token string, userId uint) error {
	if f.deleteErr != nil {
		return f.deleteErr
	}
	delete(f.tokens, token)
	return nil
}

func withEnvFile(t *testing.T) {
	t.Helper()

	dir := t.TempDir()
	envPath := filepath.Join(dir, ".env")
	envContents := []byte("JWT_SECRET_KEY=secret\nJWT_ACCESS_EXPIRY=1h\nJWT_REFRESH_EXPIRY=24h\n")
	if err := os.WriteFile(envPath, envContents, 0o600); err != nil {
		t.Fatalf("failed to write .env file: %v", err)
	}

	origDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get working directory: %v", err)
	}

	if err := os.Chdir(dir); err != nil {
		t.Fatalf("failed to change working directory: %v", err)
	}

	t.Cleanup(func() {
		_ = os.Chdir(origDir)
	})
}

func TestAuthServiceRegisterSuccess(t *testing.T) {
	userRepo := &fakeUserRepo{users: map[string]*domain.User{}}
	refreshRepo := &fakeRefreshRepo{tokens: map[string]*domain.RefreshToken{}}

	service := NewAuthService(userRepo, refreshRepo)
	user, err := service.Register(&domain.RegisterInput{
		Email:    "new@example.com",
		Password: "secure-password",
		Name:     "New User",
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if user.ID == 0 {
		t.Fatal("expected user ID to be set")
	}

	if user.Email != "new@example.com" {
		t.Fatalf("expected email new@example.com, got %s", user.Email)
	}

	if userRepo.users["new@example.com"] == nil {
		t.Fatal("expected user to be stored in repository")
	}
}

func TestAuthServiceRegisterDuplicateEmail(t *testing.T) {
	userRepo := &fakeUserRepo{
		users: map[string]*domain.User{
			"existing@example.com": {ID: 10, Email: "existing@example.com"},
		},
	}
	refreshRepo := &fakeRefreshRepo{tokens: map[string]*domain.RefreshToken{}}

	service := NewAuthService(userRepo, refreshRepo)
	_, err := service.Register(&domain.RegisterInput{
		Email:    "existing@example.com",
		Password: "password",
		Name:     "Existing User",
	})
	if err == nil {
		t.Fatal("expected error for duplicate email, got nil")
	}
}

func TestAuthServiceLoginInvalidCredentials(t *testing.T) {
	userRepo := &fakeUserRepo{users: map[string]*domain.User{}}
	refreshRepo := &fakeRefreshRepo{tokens: map[string]*domain.RefreshToken{}}

	service := NewAuthService(userRepo, refreshRepo)
	_, err := service.Login(&domain.LoginInput{
		Email:    "missing@example.com",
		Password: "password",
	})
	if err == nil {
		t.Fatal("expected error for invalid credentials, got nil")
	}
}

func TestAuthServiceLoginSuccess(t *testing.T) {
	withEnvFile(t)

	hashed, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("failed to hash password: %v", err)
	}

	userRepo := &fakeUserRepo{
		users: map[string]*domain.User{
			"user@example.com": {
				ID:           1,
				Email:        "user@example.com",
				PasswordHash: string(hashed),
				Name:         "Test User",
			},
		},
	}
	refreshRepo := &fakeRefreshRepo{tokens: map[string]*domain.RefreshToken{}}

	service := NewAuthService(userRepo, refreshRepo)
	resp, err := service.Login(&domain.LoginInput{
		Email:    "user@example.com",
		Password: "password",
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if resp.AccessToken == "" {
		t.Fatal("expected access token to be set")
	}

	if resp.User.Email != "user@example.com" {
		t.Fatalf("expected email user@example.com, got %s", resp.User.Email)
	}

	if len(refreshRepo.tokens) != 1 {
		t.Fatal("expected refresh token to be stored")
	}
}

func TestAuthServiceRefreshAccessTokenSuccess(t *testing.T) {
	withEnvFile(t)

	refreshToken, err := jwt.GenerateRefreshToken(4, "user@example.com")
	if err != nil {
		t.Fatalf("failed to generate refresh token: %v", err)
	}

	userRepo := &fakeUserRepo{users: map[string]*domain.User{}}
	refreshRepo := &fakeRefreshRepo{
		tokens: map[string]*domain.RefreshToken{
			refreshToken: {
				UserID:    4,
				Token:     refreshToken,
				ExpiresAt: time.Now().Add(2 * time.Hour),
			},
		},
	}

	service := NewAuthService(userRepo, refreshRepo)
	accessToken, err := service.RefreshAccessToken(&domain.RefreshInput{
		RefreshToken: refreshToken,
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if accessToken == "" {
		t.Fatal("expected access token to be set")
	}
}
