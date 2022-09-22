package application

import (
	"github.com/BrianToro/simple_api/config"
	"github.com/BrianToro/simple_api/domain/models"
	"github.com/BrianToro/simple_api/infrastructure/persistance"
)

func GetUser(id string) (*models.Users, error) {
	conn, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}

	repo := persistance.NewUsersRepositoryImpl(conn)
	return repo.Get(id)
}

func GetAll() ([]*models.Users, error) {
	conn, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}

	repo := persistance.NewUsersRepositoryImpl(conn)
	return repo.GetAll()
}

func Create(user *models.Users) (*models.Users, error) {
	conn, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}

	repo := persistance.NewUsersRepositoryImpl(conn)
	return repo.Create(user)
}

func Delete(id string) error {
	conn, err := config.ConnectDB()
	if err != nil {
		return err
	}
	repo := persistance.NewUsersRepositoryImpl(conn)

	// Validate if the user exist
	user, err := repo.Get(id)
	if err != nil {
		return err
	}

	return repo.Delete(user.Id)
}
