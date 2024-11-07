package repositories

import (
	"context"
	"errors"

	"github.com/Kei-K23/user-management-system-api/config"
	"github.com/Kei-K23/user-management-system-api/models"
	"github.com/jackc/pgx/v5"
)

var ErrUserNotFound = errors.New("user not found")

type UserRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	UpdateUser(id int, user *models.User) (*models.User, error)
	DeleteUser(id int) (int, error)
	GetUsers() ([]*models.User, error)
	GetUserById(id int) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
}

type userRepository struct {
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

// CreateUser implements UserRepository.
func (r *userRepository) CreateUser(user *models.User) (*models.User, error) {
	query := `INSERT INTO users (username, full_name, email, password_hash, role_id) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at, updated_at`

	err := config.DB.QueryRow(context.Background(), query, user.Username, user.FullName, user.Email, user.Password, user.RoleId).Scan(&user.Id, &user.CreatedAt, &user.UpdatedAt)

	return user, err
}

// GetUserById implements UserRepository.
func (r *userRepository) GetUserById(id int) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, username, full_name, email, role_id, created_at, updated_at FROM users WHERE id = $1`

	err := config.DB.QueryRow(context.Background(), query, id).Scan(&user.Id, &user.Username, &user.FullName, &user.Email, &user.RoleId, &user.CreatedAt, &user.UpdatedAt)

	if err == pgx.ErrNoRows {
		return nil, ErrUserNotFound
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByUsername implements UserRepository.
func (r *userRepository) GetUserByUsername(username string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, username, full_name, email, role_id, created_at, updated_at FROM users WHERE username = $1`

	err := config.DB.QueryRow(context.Background(), query, username).Scan(&user.Id, &user.Username, &user.FullName, &user.Email, &user.RoleId, &user.CreatedAt, &user.UpdatedAt)

	if err == pgx.ErrNoRows {
		return nil, ErrUserNotFound
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUsers implements UserRepository.
func (r *userRepository) GetUsers() ([]*models.User, error) {
	var users []*models.User
	query := `SELECT id, username, full_name, email, role_id, created_at, updated_at FROM users`

	rows, err := config.DB.Query(context.Background(), query)

	if err == pgx.ErrNoRows {
		return nil, ErrRoleNotFound
	}

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user := &models.User{}
		rows.Scan(&user.Id, &user.Username, &user.FullName, &user.Email, &user.RoleId, &user.CreatedAt, &user.UpdatedAt)

		users = append(users, user)
	}

	return users, nil
}

// DeleteUser implements UserRepository.
func (r *userRepository) DeleteUser(id int) (int, error) {
	query := `DELETE FROM users WHERE id = $1;`

	_, err := config.DB.Exec(context.Background(), query, id)
	return id, err
}

// UpdateUser implements UserRepository.
func (r *userRepository) UpdateUser(id int, user *models.User) (*models.User, error) {
	query := `UPDATE users 
	SET username = $1, full_name = $2, email = $3, password_hash = $4, role_id = $5 
	WHERE id = $6
	RETURNING id, username, full_name, email, role_id, created_at, updated_at`

	err := config.DB.QueryRow(context.Background(), query, user.Username, user.FullName, user.Email, user.Password, user.RoleId, id).Scan(&user.Id, &user.Username, &user.FullName, &user.Email, &user.RoleId, &user.CreatedAt, &user.UpdatedAt)

	return user, err
}
