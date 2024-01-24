package api

import (
	"time"
	"trekkstay/api/routes"
	"trekkstay/config"
	"trekkstay/config/models"
	database "trekkstay/pkgs/db"
	"trekkstay/pkgs/transport/http/server"
	"trekkstay/utils"
)

func NewServer() (*server.HTTPServer, error) {
	appConfig := config.LoadConfig(&models.AppConfig{}).(*models.AppConfig)
	dbConfig := config.LoadConfig(&models.DBConfig{}).(*models.DBConfig)

	connection := database.Connection{
		SSLMode:               database.Disable,
		Host:                  dbConfig.DBHost,
		Port:                  dbConfig.DBPort,
		Database:              dbConfig.DBName,
		User:                  dbConfig.DBUserName,
		Password:              dbConfig.DBPassword,
		MaxIdleConnections:    dbConfig.MaxIdleConnections,
		MaxOpenConnections:    dbConfig.MaxOpenConnections,
		ConnectionMaxIdleTime: time.Duration(dbConfig.ConnectionMaxIdleTime),
		ConnectionMaxLifeTime: time.Duration(dbConfig.ConnectionMaxLifeTime),
	}

	db := database.InitDatabase(connection)

	s := server.NewHTTPServer(
		server.AddName(appConfig.ServiceName),
		server.AddPort(appConfig.ServicePort),
		server.SetGracefulShutdownTimeout(time.Duration(appConfig.ServiceTimeout)),
	)

	requestValidator := utils.NewValidator()

	srv := &routes.RouteHandler{
		UserHandler:   routes.NewUserHandler(db, requestValidator),
		RegionHandler: routes.NewRegionHandler(db),
	}

	s.AddGroupRoutes(srv.InitGroupRoutes())

	return s, nil
}
