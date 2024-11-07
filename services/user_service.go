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

func (u *UserService) Create(username, fullName, email, password string, roleId int) (*models.User, error) {

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: username,
		FullName: fullName,
		Email:    email,
		Password: string(hashedPass),
		RoleId:   roleId,
	}

	return u.userRepo.CreateUser(user)
}

func (u *UserService) Update(id int, username, fullName, email, password string, roleId int) (*models.User, error) {

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: username,
		FullName: fullName,
		Email:    email,
		Password: string(hashedPass),
		RoleId:   roleId,
	}

	return u.userRepo.UpdateUser(id, user)
}

func (u *UserService) Delete(id int) (int, error) {
	return u.userRepo.DeleteUser(id)
}

func (u *UserService) GetUsers() ([]*models.User, error) {
	return u.userRepo.GetUsers()
}

func (u *UserService) GetById(id int) (*models.User, error) {
	return u.userRepo.GetUserById(id)
}

func (u *UserService) GetByUsername(name string) (*models.User, error) {
	return u.userRepo.GetUserByUsername(name)
}
