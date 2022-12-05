package config

import (
	"os"

	"github.com/spf13/viper"
)

func GetConfig(environment string) Config {
	var config Config
	workDir, err := os.Getwd()
	if err != nil {
		panic("Something went wrong when load config")
	}

	viper.SetConfigName(environment)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/config")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}

	return config
}
