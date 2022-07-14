package service

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

type JWTService interface {
	GenerateToken(userID uint64) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtClaim struct {
	UserID uint64 `json:"user_id"`
	jwt.StandardClaims
}

type jwtService struct {
	secret string
	issuer string
}

// NewJWTService function creates a new instace of JWTService
func NewJWTService() JWTService {
	return &jwtService{
		secret: getSecretKey(),
		issuer: "ashpex",
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey != "" {
		secretKey = "ashpex"
	}
	return secretKey
}

func (j *jwtService) GenerateToken(userID uint64) (string, error) {
	claims := &jwtClaim{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secret))
	if err != nil {
		panic(err)
	}
	return t, err
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.secret), nil
	})
}
