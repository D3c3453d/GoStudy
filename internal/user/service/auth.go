package service

import (
	"GoStudy/internal/user/entity"
	"GoStudy/internal/user/repository"
	"crypto/sha256"
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
	id int `json:"id"`
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

func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
