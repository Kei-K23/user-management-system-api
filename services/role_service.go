package services

import (
	"github.com/Kei-K23/user-management-system-api/models"
	"github.com/Kei-K23/user-management-system-api/repositories"
)

type RoleService interface {
	CreateRole(role models.Role) error
	GetRoleById(id int) (*models.Role, error)
	GetRoleByName(name string) (*models.Role, error)
}

type roleService struct {
	roleRepo repositories.RoleRepository
}

func NewRoleService(roleRepo repositories.RoleRepository) RoleService {
	return &roleService{roleRepo}
}

// CreateRole implements RoleService.
func (r *roleService) CreateRole(role models.Role) error {
	return r.roleRepo.CreateRole(role)
}

// GetRoleById implements RoleService.
func (r *roleService) GetRoleById(id int) (*models.Role, error) {
	return r.roleRepo.GetRoleById(id)
}

// GetRoleByName implements RoleService.
func (r *roleService) GetRoleByName(name string) (*models.Role, error) {
	return r.roleRepo.GetRoleByName(name)
}
