package service

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"template/entities"
)

func (s *service) Login(userCred *entities.User) (token string, err error) {
	user, err := s.model.GetUser(userCred.Email)
	if err != nil {
		log.Default().Println(err.Error())
		return "", err
	}
	if user == nil {
		return "", errors.New("user not found")
	}

	if user.Password != userCred.Password {
		fmt.Println("Invalid password", user.Password, userCred.Password)
		return "", errors.New("invalid password")
	}

	return generateToken(), nil
}

// This function generates a random token. This function logic should be changed to create a jwt token later.
func generateToken() string {
	// Generate a random byte slice
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	// Encode the byte slice to a string using base64 encoding
	return base64.URLEncoding.EncodeToString(b)
}
