package main

import (
	"flag"
	"os"
	"strconv"
	"trekkstay/pkgs/log"
)

var (
	configPath = flag.String("conf", "./env/dev.prod", "application config path")
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
}
