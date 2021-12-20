package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

//Config struct stores all configuration values for project using Viper
type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	ServerPort    string `mapstructure:"SERVER_PORT"`
	PgUser        string
	PgDB          string
	PgHost        string
	PgPort        string
}

//LoadConfig reads configuration from .env  file
func LoadConfig(path string) (config Config, err error) {

	//Setting default path for config file
	if path == "" {
		path = "./config"
	}

	//Declareting path and type for config file
	viper.AddConfigPath(path)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	//Parsing config vals from file (first step)
	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	//Parsing config vals from file (second step)
	err = viper.Unmarshal(&config)

	//If it is nescessary to hardcore port - just add value for SERVER_PORT in .env file
	if config.ServerPort == "" {
		config.ServerPort = os.Getenv("SERVER_PORT")
	}

	viper.AutomaticEnv()
	viper.BindEnv("PgUser")
	os.Setenv("PGUSER", "postgress")
	config.PgUser = viper.GetString("PgUser")

	return

}

// GetConfigString returns specific value from config file
func GetConfigString(ValName string) (val string, err error) {

	//Loading config
	config, _ := LoadConfig("./config")

	switch ValName {
	case "DBDriver":
		val = config.DBDriver
	case "DBSource":
		val = config.DBSource
	case "ServerAddress":
		val = config.ServerAddress
	default:
		err = fmt.Errorf("Cannot find value " + ValName)
	}

	return

}