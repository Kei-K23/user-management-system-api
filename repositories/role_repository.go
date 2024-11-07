package repositories

import (
	"context"
	"errors"

	"github.com/Kei-K23/user-management-system-api/config"
	"github.com/Kei-K23/user-management-system-api/models"
	"github.com/jackc/pgx/v5"
)

var ErrRoleNotFound = errors.New("role not found")

type RoleRepository interface {
	CreateRole(role *models.Role) (*models.Role, error)
	GetRoles() ([]*models.Role, error)
	GetRoleById(id int) (*models.Role, error)
	GetRoleByName(name string) (*models.Role, error)
	DeleteRole(id int) (int, error)
	UpdateRole(id int, role *models.Role) (*models.Role, error)
}

type roleRepository struct {
}

func NewRoleRepository() RoleRepository {
	return &roleRepository{}
}

// CreateRole implements RoleRepository.
func (r *roleRepository) CreateRole(role *models.Role) (*models.Role, error) {
	query := `INSERT INTO roles (name, description) VALUES ($1, $2) RETURNING id;`
	err := config.DB.QueryRow(context.Background(), query, role.Name, role.Description).Scan(&role.Id)
	return role, err
}

// GetRoleById implements RoleRepository.
func (r *roleRepository) GetRoleById(id int) (*models.Role, error) {
	role := &models.Role{}
	query := `SELECT * from roles WHERE id = $1;`

	err := config.DB.QueryRow(context.Background(), query, id).Scan(&role.Id, &role.Name, &role.Description)

	if err == pgx.ErrNoRows {
		return nil, ErrRoleNotFound
	}

	if err != nil {
		return nil, err
	}

	return role, err
}

// GetRoleByName implements RoleRepository.
func (r *roleRepository) GetRoleByName(name string) (*models.Role, error) {
	role := &models.Role{}
	query := `SELECT * from roles WHERE name LIKE $1`

	err := config.DB.QueryRow(context.Background(), query, name).Scan(&role.Id, &role.Name, &role.Description)

	if err == pgx.ErrNoRows {
		return nil, ErrRoleNotFound
	}

	if err != nil {
		return nil, err
	}

	return role, err
}

// GetRoles implements RoleRepository.
func (r *roleRepository) GetRoles() ([]*models.Role, error) {
	var roles []*models.Role

	query := `SELECT * from roles;`

	rows, err := config.DB.Query(context.Background(), query)

	if err == pgx.ErrNoRows {
		return nil, ErrRoleNotFound
	}

	if err != nil {
		return nil, err
	}

	// Read the rows
	for rows.Next() {
		role := &models.Role{}
		rows.Scan(&role.Id, &role.Name, &role.Description)
		roles = append(roles, role)
	}

	return roles, nil
}

// DeleteRole implements RoleRepository.
func (r *roleRepository) DeleteRole(id int) (int, error) {
	query := `DELETE FROM roles WHERE id = $1;`

	_, err := config.DB.Exec(context.Background(), query, id)
	return id, err
}

// UpdateRole implements RoleRepository.
func (r *roleRepository) UpdateRole(id int, role *models.Role) (*models.Role, error) {
	query := `UPDATE roles 
	SET name = $1, description = $2
	WHERE id = $3`

	err := config.DB.QueryRow(context.Background(), query, role.Name, role.Description, id).Scan(&role.Id, &role.Name, &role.Description)

	return role, err
}
