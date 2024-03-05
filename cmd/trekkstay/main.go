package main

import (
	"flag"
	"os"
	"strconv"
	"trekkstay/api"
	"trekkstay/config"
	"trekkstay/config/models"
	"trekkstay/docs"
	"trekkstay/pkgs/log"
	"trekkstay/pkgs/transport/http/server"
)

var (
	configPath = flag.String("conf", "./env/dev.env", "application config path")
	migration  = flag.Bool("migrate", false, "migrate database")
)

func init() {
	flag.Parse()

	if err := os.Setenv("CONFIG_PATH", *configPath); err != nil {
		panic(err)
	}

	if err := os.Setenv("CONFIG_ALLOW_MIGRATION",
		strconv.FormatBool(*migration)); err != nil {
		panic(err)
	}
}

// @title							Trekkstay - Hotel Booking System API
// @version        					v0.0.1
// @description     				API system for Trekkstay - Hotel Booking System
// @termsOfService  				https://swagger.io/

// @contact.name   					Trekkstay Team
// @contact.url    					https://www.trekkstay.com
// @contact.email  					support@trekkstay.com

// @license.name  					Apache 2.0
// @license.url   					https://www.apache.org/licenses/LICENSE-2.0.html

// @host      						52.163.61.213:8888
// @BasePath  						/api/v1
// @securitydefinitions.apikey  	JWT
// @in                          	header
// @name                        	Authorization
func main() {
	log.JsonLogger.Info("Starting server...")

	appConfig := config.LoadConfig(&models.AppConfig{}).(*models.AppConfig)

	docs.SwaggerInfo.Schemes = []string{"https"}
	if appConfig.BuildEnv != "prod" {
		docs.SwaggerInfo.Host = "localhost:8888"
		docs.SwaggerInfo.Schemes = []string{"http"}
	} else {
		docs.SwaggerInfo.Host = "52.163.61.213:8888"
		docs.SwaggerInfo.Schemes = []string{"http"}
	}

	server.MustRun(api.NewServer())
}
