package repositories

import "github.com/BrianToro/simple_api/domain/models"

type UsersRepository interface {
	Get(id string) (*models.Users, error)
	GetAll() ([]*models.Users, error)
	Create(user *models.Users) (*models.Users, error)
	Delete(id string) error
}
