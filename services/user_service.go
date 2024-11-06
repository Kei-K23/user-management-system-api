package services

import (
	"github.com/Kei-K23/user-management-system-api/models"
	"github.com/Kei-K23/user-management-system-api/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) *UserService {
	return &UserService{userRepo}
}

func (r *UserService) Create(username, fullName, email, password string, roleId int) error {

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &models.User{
		Username: username,
		FullName: fullName,
		Email:    email,
		Password: string(hashedPass),
		RoleId:   roleId,
	}

	return r.userRepo.CreateUser(user)
}
