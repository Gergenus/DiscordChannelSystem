package tokens

import (
	"errors"
	"time"

	"github.com/Gergenus/internal/repository"
	"github.com/golang-jwt/jwt/v5"
)

const (
	Code = "Steve"
)

type TokenManager interface {
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type JWTTokenManager struct {
	userRepo repository.UserRepository
}

func NewJWTTokenManager(userRepo repository.UserRepository) *JWTTokenManager {
	return &JWTTokenManager{userRepo: userRepo}
}

type TokenClaims struct {
	jwt.RegisteredClaims
	UserId int
}

func (j *JWTTokenManager) GenerateToken(username, password string) (string, error) {
	user, err := j.userRepo.GetUser(username)
	if err != nil {
		return "", err
	}
	if username != user.UserName || password != user.HashPassword {
		return "", errors.New("password is incorrect")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		user.Uid,
	})
	generatedToken, err := token.SignedString([]byte(Code))
	if err != nil {
		return "", err
	}
	return generatedToken, nil
}

func (j *JWTTokenManager) ParseToken(token string) (int, error) {
	tkn, err := jwt.ParseWithClaims(token, &TokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid signing method")
		}
		return []byte(Code), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := tkn.Claims.(*TokenClaims)
	if !ok {
		return 0, errors.New("token has invalid type")
	}
	return claims.UserId, nil
}
