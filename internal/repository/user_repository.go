package repository

import (
	"database/sql"
	"errors"

	"github.com/ahmadzakyarifin/gin-jwt-auth/internal/entity"
)

// ========================================
//go:generate mockery --name=AuthRepository --output=mocks --outpkg=mocks
// =======================================

type AuthRepository interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type authRepository struct {
	DB *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepository{DB: db}
}

func (r *authRepository) Create(user *entity.User) error {
	query := `
		INSERT INTO users (name, email, password, role, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err := r.DB.Exec(
		query,
		user.Name,
		user.Email,
		user.Password,
		user.Role,
		user.CreatedAt,
		user.UpdateAt,
	)
	return err
}

func (r *authRepository) FindByEmail(email string) (*entity.User, error) {
	query := `
		SELECT id, name, email, password, role
		FROM users
		WHERE email = ?
	`

	user := &entity.User{}
	err := r.DB.QueryRow(query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Role,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return user, nil
}
