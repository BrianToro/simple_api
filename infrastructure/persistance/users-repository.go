package persistance

import (
	"github.com/BrianToro/simple_api/domain/models"
	"github.com/BrianToro/simple_api/domain/repositories"
	"gorm.io/gorm"
)

type UsersRepositoryImpl struct {
	Conn *gorm.DB
}

func NewUsersRepositoryImpl(conn *gorm.DB) repositories.UsersRepository {
	return &UsersRepositoryImpl{Conn: conn}
}

func (repo *UsersRepositoryImpl) Get(id string) (*models.Users, error) {
	userFind := &models.Users{Id: id}
	user := &models.Users{}
	if err := repo.Conn.First(&user, userFind).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *UsersRepositoryImpl) GetAll() ([]*models.Users, error) {
	var users []*models.Users
	if err := repo.Conn.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (repo *UsersRepositoryImpl) Create(user *models.Users) (*models.Users, error) {
	if err := repo.Conn.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UsersRepositoryImpl) Delete(id string) error {
	userFind := &models.Users{Id: id}
	if err := repo.Conn.Delete(&userFind).Error; err != nil {
		return err
	}
	return nil
}
