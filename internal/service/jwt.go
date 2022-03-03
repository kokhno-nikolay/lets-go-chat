package service

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/kokhno-nikolay/lets-go-chat/internal/repository"
)

type JWTService struct {
	repo      repository.JWT
	secretKey string
}

func NewJWTService(repo repository.JWT, secretKey string) *JWTService {
	return &JWTService{
		repo:      repo,
		secretKey: secretKey,
	}
}

type authCustomClaims struct {
	UUID     string `json:"uuid"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func (s *JWTService) GenerateToken(uuid, username string) string {
	claims := &authCustomClaims{
		uuid,
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		panic(err)
	}

	return t
}

func (s *JWTService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		return []byte(s.secretKey), nil
	})
}

func (s *JWTService) ExtractClaims(tokenStr string) (jwt.MapClaims, bool) {
	hmacSecret := []byte(s.secretKey)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}
