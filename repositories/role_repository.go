package repositories

import (
	"database/sql"

	"github.com/Kei-K23/user-management-system-api/models"
)

type RoleRepository interface {
	CreateRole(role models.Role) error
	GetRoleById(id int) (*models.Role, error)
	GetRoleByName(name string) (*models.Role, error)
}

type roleRepository struct {
	db *sql.DB
}

func NewRoleRepository(db *sql.DB) RoleRepository {
	return &roleRepository{db}
}

// CreateRole implements RoleRepository.
func (r *roleRepository) CreateRole(role models.Role) error {
	query := `INSERT INTO roles (name, description) VALUES ($1, $2) RETURNING id`
	return r.db.QueryRow(query, role.Name, role.Description).Scan(&role.Id)
}

// GetRoleById implements RoleRepository.
func (r *roleRepository) GetRoleById(id int) (*models.Role, error) {
	role := &models.Role{}
	query := `SELECT * from roles WHERE id = $1`

	err := r.db.QueryRow(query, id).Scan(&role.Id, &role.Name, &role.Description)
	return role, err
}

// GetRoleByName implements RoleRepository.
func (r *roleRepository) GetRoleByName(name string) (*models.Role, error) {
	role := &models.Role{}
	query := `SELECT * from roles WHERE name = $1`

	err := r.db.QueryRow(query, name).Scan(&role.Id, &role.Name, &role.Description)
	return role, err
}
