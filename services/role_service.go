package services

import (
	"github.com/Kei-K23/user-management-system-api/models"
	"github.com/Kei-K23/user-management-system-api/repositories"
)

type RoleService struct {
	roleRepo repositories.RoleRepository
}

func NewRoleService(roleRepo repositories.RoleRepository) *RoleService {
	return &RoleService{roleRepo}
}

// CreateRole implements RoleService.
func (r *RoleService) Create(name, description string) (*models.Role, error) {
	role := &models.Role{
		Name:        name,
		Description: description,
	}

	return r.roleRepo.CreateRole(role)
}

// GetRoleById implements RoleService.
func (r *RoleService) GetRoles() ([]*models.Role, error) {
	return r.roleRepo.GetRoles()
}

// GetRoleById implements RoleService.
func (r *RoleService) GetById(id int) (*models.Role, error) {
	return r.roleRepo.GetRoleById(id)
}

// GetRoleByName implements RoleService.
func (r *RoleService) GetByName(name string) (*models.Role, error) {
	return r.roleRepo.GetRoleByName(name)
}

func (r *RoleService) Update(id int, name, description string) (*models.Role, error) {
	role := &models.Role{
		Name:        name,
		Description: description,
	}

	return r.roleRepo.UpdateRole(id, role)
}

func (r *RoleService) Delete(id int) (int, error) {
	return r.roleRepo.DeleteRole(id)
}
