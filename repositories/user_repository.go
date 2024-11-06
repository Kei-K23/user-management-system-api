package repositories

import (
	"context"

	"github.com/Kei-K23/user-management-system-api/config"
	"github.com/Kei-K23/user-management-system-api/models"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserById(id int) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
}

type userRepository struct {
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

// CreateUser implements UserRepository.
func (r *userRepository) CreateUser(user *models.User) error {
	query := `INSERT INTO users (username, full_name, email, password_hashed, role_id) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	return config.DB.QueryRow(context.Background(), query, user.Username, user.FullName, user.Email, user.Password, user.RoleId).Scan(&user.Id)
}

// GetUserById implements UserRepository.
func (r *userRepository) GetUserById(id int) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, username, full_name, email, role_id, created_at, updated_at WHERE id = $1`

	err := config.DB.QueryRow(context.Background(), query, id).Scan(&user.Id, &user.Username, &user.FullName, &user.Email, &user.RoleId, &user.CreatedAt, &user.UpdatedAt)

	return user, err
}

// GetUserByUsername implements UserRepository.
func (r *userRepository) GetUserByUsername(username string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, username, full_name, email, role_id, created_at, updated_at WHERE username = $1`

	err := config.DB.QueryRow(context.Background(), query, username).Scan(&user.Id, &user.Username, &user.FullName, &user.Email, &user.RoleId, &user.CreatedAt, &user.UpdatedAt)
	return user, err
}
