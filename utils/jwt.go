package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtSecret = []byte("E822EBFC0D9E3DB3BFD2910242299A374F0BE293C7613639C6F86920AE6CFDC2")

func GenerateJWT(userId, roleId int) (string, error) {
	claims := jwt.MapClaims{
		"userId": userId,
		"roleId": roleId,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}
