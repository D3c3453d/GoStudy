package service

import (
	"GoStudy/internal/user/entity"
	"GoStudy/internal/user/repository"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	salt       = "asdasdq34234r32hurfh32vjh"
	tokenTime  = 5 * time.Minute
	signingKey = "afcsabdhbadf6y465weagdbanfsjhd4q3yehajsfm"
)

type AuthService struct {
	repo repository.AuthReposI
}

func NewAuthService(repo repository.AuthReposI) *AuthService {
	return &AuthService{repo: repo}
}
func (s *AuthService) Create(account entity.Account) (int, error) {
	account.Password = generatePasswordHash(account.Password)
	return s.repo.Create(account)
}

type tokenClaims struct {
	jwt.StandardClaims
	Id int `json:"id"`
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	account, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTime).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		account.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.Id, nil
}

func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
