package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/iakigarci/go-ddd-microservice-template/internal/domain/models/entities"
	"github.com/iakigarci/go-ddd-microservice-template/internal/domain/ports"
	"go.uber.org/zap"
)

type authService struct {
	userService ports.UserService
	jwtSecret   []byte
	logger      *zap.Logger
}

func New(userService ports.UserService, jwtSecret []byte, logger *zap.Logger) *authService {
	return &authService{
		userService: userService,
		jwtSecret:   jwtSecret,
		logger:      logger,
	}
}

func (svc *authService) GenerateToken(user *entities.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.Email,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(svc.jwtSecret)
	if err != nil {
		svc.logger.Error("failed to generate token", zap.Error(err))
		return "", err
	}

	return tokenString, nil
}

func (svc *authService) ValidateToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return svc.jwtSecret, nil
	})

	if err != nil {
		svc.logger.Error("failed to validate token", zap.Error(err))
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		err = fmt.Errorf("invalid token")
		svc.logger.Error(err.Error())
		return "", err
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		err = fmt.Errorf("invalid user id")
		svc.logger.Error(err.Error())
		return "", err
	}

	return userID, nil
}
