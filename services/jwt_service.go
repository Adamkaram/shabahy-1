package services

import (
	"fmt"
	"github.com/ElegantSoft/shabahy/common"
	"github.com/dgrijalva/jwt-go"
	"log"
	"os"
	"time"
)

type JWTService interface {
	GenerateToken(id uint) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type jwtCustomClaims struct {
	ID uint `json:"id"`
	jwt.StandardClaims
}

type JwtService struct {
	issuer    string
}

func (j *JwtService) GenerateToken(id uint) (string, error) {
	claims := &jwtCustomClaims{
		ID:             id,
		StandardClaims: jwt.StandardClaims{
			Issuer: j.issuer,
			IssuedAt: time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv(common.KJwtSecret)))

}

func (j *JwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv(common.KJwtSecret)), nil
	})
}

func getSecretKey() string {
	secret := os.Getenv(common.KJwtSecret)
	log.Println("secret from func", secret)

	if secret == "" {
		secret = "secretfskfjsdkfjslfsdjl"
	}
	log.Println("secret from func", secret)
	return secret
}

func NewJWTService() JWTService {
	return &JwtService{
		issuer:    "shabahy.com",
	}
}

