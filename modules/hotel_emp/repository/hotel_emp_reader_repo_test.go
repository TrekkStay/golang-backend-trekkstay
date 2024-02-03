package repository

import (
	"context"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"trekkstay/config"
	"trekkstay/config/models"
	"trekkstay/pkgs/dbs/postgres"
)

func TestFindEmpByCondition(t *testing.T) {
	err := os.Setenv("CONFIG_PATH", "../../../env/dev.env")
	if err != nil {
		return
	}

	dbConfig := config.LoadConfig(&models.DBConfig{}).(*models.DBConfig)

	connection := postgres.Connection{
		SSLMode:  postgres.Disable,
		Host:     dbConfig.DBHost,
		Port:     dbConfig.DBPort,
		Database: dbConfig.DBName,
		User:     dbConfig.DBUserName,
		Password: dbConfig.DBPassword,
	}

	db := postgres.InitDatabase(connection)

	repo := New(*db)

	t.Run("should find hotel employee by condition", func(t *testing.T) {
		_, err := repo.FindHotelEmpByCondition(context.Background(), map[string]interface{}{
			"email": "WQhjz@example.com",
		})

		assert.Nil(t, err)
	})
}
