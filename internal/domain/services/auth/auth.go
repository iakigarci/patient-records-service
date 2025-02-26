package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/iakigarci/go-ddd-microservice-template/internal/domain/models/entities"
)

type authService struct {
	userRepository user.UserRepository
	jwtSecret      []byte
}

func New(userRepository user.UserRepository, jwtSecret []byte) *authService {
	return &authService{
		userRepository: userRepository,
		jwtSecret:      jwtSecret,
	}
}

func (svc *authService) GenerateToken(user *entities.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(svc.jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (svc *authService) ValidateToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.jwtSecret, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["user_id"].(string), nil
	}

	return "", fmt.Errorf("invalid token")
}
