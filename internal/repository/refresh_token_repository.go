package repository

import (
	"database/sql"
	"errors"

	"github.com/r6rap/Quanta/internal/domain"
	"github.com/r6rap/Quanta/pkg/database"
	"gorm.io/gorm"
)

type RefreshRepository struct {
	db *database.DB
}

func NewRefreshRepository(db *database.DB) *RefreshRepository {
	return &RefreshRepository{db: db}
}

func (r *RefreshRepository) Create(req *domain.RefreshToken) error {
	query := `INSERT INTO refresh_tokens (user_id, token, expires_at) VALUES ($1, $2, $3) RETURNING id, created_at`

	row := r.db.Raw(query, req.UserID, req.Token, req.ExpiresAt).Row()

	return row.Scan(&req.ID, &req.CreatedAt)
}

func (r *RefreshRepository) GetByToken(token string, userId uint) (*domain.RefreshToken, error) {
	storedToken := domain.RefreshToken{}

	query := `SELECT id, user_id, token, expires_at FROM refresh_tokens WHERE token = $1 AND user_id = $2`

	row := r.db.Raw(query, token, userId).Row()

	err := row.Scan(&storedToken.ID, &storedToken.UserID, &storedToken.Token, &storedToken.ExpiresAt)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return &storedToken, err
}

func (r *RefreshRepository) Delete(token string, userId uint) error {
	query := `DELETE FROM refresh_tokens WHERE token = $1 AND user_id = $2`

	row := r.db.Exec(query, token, userId).RowsAffected

	if row == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
