package service

import (
	"github.com/Gergenus/internal/repository"
	hasherpkg "github.com/Gergenus/pkg/Hasher"
	"github.com/Gergenus/pkg/tokens"
)

type Auth interface {
	SignUp(userName string, password string) (int, error)
	SignIn(userName string, password string) (string, error)
}

type JWTauth struct {
	hasher       hasherpkg.HasherInterface
	userRepo     repository.UserRepository
	tokenManager tokens.TokenManager
}

func NewJWTauth(hasher hasherpkg.HasherInterface, userRepo repository.UserRepository, tokenManager tokens.TokenManager) *JWTauth {
	return &JWTauth{
		hasher:       hasher,
		userRepo:     userRepo,
		tokenManager: tokenManager,
	}
}

func (j *JWTauth) SignUp(userName string, password string) (int, error) {
	password = j.hasher.Hash(password)
	id, err := j.userRepo.CreateUser(userName, password)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (j *JWTauth) SignIn(userName string, password string) (string, error) {
	tkn, err := j.tokenManager.GenerateToken(userName, j.hasher.Hash(password))
	if err != nil {
		return "", err
	}
	return tkn, nil
}
