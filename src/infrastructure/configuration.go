package infrastructure

import (
	"log"

	"github.com/spf13/viper"
)

var (
	configuration *viper.Viper
	isLoaded      bool = false
)

func loadEnvironmentConfigurationFile() {
	configuration = viper.New()

	configuration.AddConfigPath("./environment")
	configuration.SetConfigName("app")
	configuration.SetConfigType("json")
}

func GetConfig(key string) string {
	if isLoaded {
		return configuration.Get(key).(string)
	}

	loadEnvironmentConfigurationFile()

	err := configuration.ReadInConfig()
	if err != nil {
		log.Printf("[ERROR] Fail to read configuration file: %s \n", err.Error())
		return ""
	}

	isLoaded = true
	return configuration.Get(key).(string)
}
