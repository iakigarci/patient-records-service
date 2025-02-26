package ports

import "github.com/iakigarci/go-ddd-microservice-template/internal/domain/models/entities"

type AuthService interface {
	GenerateToken(user *entities.User) (string, error)
	ValidateToken(tokenString string) (string, error)
}
