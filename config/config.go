package config

import (
	"os"
	"trekkstay/pkgs/log"
	"trekkstay/utils"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

func LoadConfig(model interface{}) interface{} {
	configPath := os.Getenv("CONFIG_PATH")

	// Load the configuration from the file.
	config, err := loadConfigFromFile(configPath)
	if err != nil {
		log.JsonLogger.Error(err.Error())
		panic(err)
	}

	// Configure the decoder to allow weakly typed input and set the result to the AppConfig struct.
	customConfig := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &model,
	}

	// Create a new decoder with the custom configuration.
	decoder, err := mapstructure.NewDecoder(customConfig)
	if err != nil {
		log.JsonLogger.Error(err.Error())
		panic(err)
	}

	// Decode the appConfig into the result struct.
	err = decoder.Decode(config)
	if err != nil {
		log.JsonLogger.Error(err.Error())
		panic(err)
	}

	return model
}

// LoadConfigFromFile loads a configuration file from the given path and returns the loaded configuration.
// It expects the path to the configuration file and the name of the configuration file as parameters.
func loadConfigFromFile(configPath string) (config interface{}, err error) {
	var dirPath = utils.GetDirectoryPath(configPath)
	var fileName = utils.GetFileName(configPath)

	viper.AddConfigPath(dirPath)
	viper.SetConfigName(fileName)
	viper.SetConfigType("env")

	// Automatically read in environment variables that match the configuration keys.
	viper.AutomaticEnv()

	// Read in the configuration file.
	err = viper.ReadInConfig()
	if err != nil {
		log.JsonLogger.Error(err.Error())
		panic(err)
	}

	// Unmarshal the configuration file into the given config interface.
	err = viper.Unmarshal(&config)
	if err != nil {
		log.JsonLogger.Error(err.Error())
		panic(err)
	}

	return
}
