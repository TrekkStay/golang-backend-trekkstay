package main

import (
	"flag"
	"os"
	"strconv"
	"trekkstay/api"
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

func main() {
	log.JsonLogger.Info("Starting server...")

	server.MustRun(api.NewServer())
}
