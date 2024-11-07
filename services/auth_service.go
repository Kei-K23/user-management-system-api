package services

import (
	"log"

	"github.com/Kei-K23/user-management-system-api/models"
	"github.com/Kei-K23/user-management-system-api/repositories"
	"github.com/Kei-K23/user-management-system-api/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo repositories.UserRepository
}

func NewAuthService(userRepo repositories.UserRepository) *AuthService {
	return &AuthService{userRepo}
}

func (u *AuthService) Register(username, fullName, email, password string, roleId int) (*models.User, error) {

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

func (u *AuthService) Login(username, password string) (string, error) {

	user, err := u.userRepo.GetDetailUserByUsername(username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	token, err := utils.GenerateJWT(user.Id, user.RoleId)
	if err != nil {
		return "", err
	}

	return token, nil
}
