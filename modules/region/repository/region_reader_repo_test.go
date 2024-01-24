package repository

import (
	"context"
	"os"
	"testing"
	"trekkstay/config"
	"trekkstay/config/models"
	database "trekkstay/pkgs/db"
)

func TestListData(t *testing.T) {
	err := os.Setenv("CONFIG_PATH", "../../../env/dev.env")
	if err != nil {
		return
	}

	dbConfig := config.LoadConfig(&models.DBConfig{}).(*models.DBConfig)

	connection := database.Connection{
		SSLMode:  database.Disable,
		Host:     dbConfig.DBHost,
		Port:     dbConfig.DBPort,
		Database: dbConfig.DBName,
		User:     dbConfig.DBUserName,
		Password: dbConfig.DBPassword,
	}

	db := database.InitDatabase(connection)

	repo := NewRegionReaderRepository(*db)

	t.Run("list province", func(t *testing.T) {
		provinces, _ := repo.ListProvinces(context.Background())

		if len(provinces) == 0 {
			t.Errorf("Province list is empty")
		}

		if len(provinces) != 63 {
			t.Errorf("Province list size is not 63")
		}

		district, _ := repo.ListDistricts(context.Background(), provinces[0].Code)

		if len(district) == 0 {
			t.Errorf("District list is empty")
		}

		ward, _ := repo.ListWards(context.Background(), district[0].Code)

		if len(ward) == 0 {
			t.Errorf("Ward list is empty")
		}
	})
}
