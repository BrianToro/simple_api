package migrations

import (
	"github.com/BrianToro/simple_api/config"
	"github.com/BrianToro/simple_api/domain/models"
)

func Migrate() error {
	conn, err := config.ConnectDB()
	if err != nil {
		return err
	}

	err = conn.Migrator().CreateTable(&models.Users{})
	if err != nil {
		return err
	}

	return nil
}
