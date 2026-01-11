package repository

import (
	"github.com/r6rap/Quanta/pkg/database"
	"github.com/r6rap/Quanta/internal/domain"

	"errors"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *database.DB
}

func NewUserRepository(db *database.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *domain.User) error {
	query := `INSERT INTO users (email, password_hash, name) VALUES ($1, $2, $3) RETURNING id, created_at`

	row := r.db.Raw(query, user.Email, user.PasswordHash, user.Name).Row()

	return row.Scan(&user.ID, &user.CreatedAt)
}

func (r *UserRepository) GetByEmail(email string) (*domain.User, error) {
	user := &domain.User{}
	query := `SELECT id, email, password_hash, name, created_at FROM users WHERE email = $1`

	rows := r.db.Raw(query, email).Row()
	
	err := rows.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.Name, &user.CreatedAt)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return user, err
}

func (r *UserRepository) GetByID(id uint) (*domain.User, error) {
	user := &domain.User{}
	query := `SELECT id, email, password_hash, name, created_at FROM users WHERE id = $1`

	row := r.db.Raw(query, id).Row()
	
	err := row.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.Name, &user.CreatedAt)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return user, err
}