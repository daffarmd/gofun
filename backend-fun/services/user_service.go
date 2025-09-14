package services

import (
	"time"

	"github.com/daffarmd/gofun/backend-fun/middleware"
	"github.com/daffarmd/gofun/backend-fun/repository"
	"github.com/golang-jwt/jwt/v5"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) Register(username, password string) {
	s.repo.Save(repository.User{
		Username: username,
		Password: password,
	})
}

func (s *UserService) Login(username, password string) bool {
	user, found := s.repo.FindByUsername(username)
	if !found {
		return false
	}

	return user.Password == password
}

func (s *UserService) generateJWT(username string) (string, error) {
	// Create claims
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // 24 hours expiration
		"iat":      time.Now().Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token
	tokenString, err := token.SignedString([]byte(middleware.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
