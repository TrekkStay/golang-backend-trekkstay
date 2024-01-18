package api

import (
	"time"
	"trekkstay/api/routes"
	"trekkstay/config"
	"trekkstay/config/models"
	"trekkstay/pkgs/transport/http/server"
)

func NewServer() (*server.HTTPServer, error) {
	appConfig := config.LoadConfig(&models.AppConfig{}).(*models.AppConfig)
	//dbConfig := config.LoadConfig(&models.DBConfig{}).(*models.DBConfig)

	s := server.NewHTTPServer(
		server.AddName(appConfig.ServiceName),
		server.AddPort(appConfig.ServicePort),
		server.SetGracefulShutdownTimeout(3*time.Second),
	)

	s.AddRoutes(routes.InitRoutes())
	s.AddGroupRoutes(routes.InitGroupRoutes())

	return s, nil
}
