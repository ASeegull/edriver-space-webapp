package config

import (
	"os"

	"github.com/spf13/viper"

	"github.com/ASeegull/edriver-space-webapp/logger"
)

//Config struct stores all configuration values for project using Viper
type Config struct {
	SignInURL       string `mapstructure:"SIGN_IN_URL"`
	SignUpURL       string `mapstructure:"SIGN_UP_URL"`
	SignOutURL      string `mapstructure:"SIGN_OUT_URL"`
	RefreshTokenURL string `mapstructure:"REFRESH_TOKEN_URL"`
	MainAppAdr      string
	PgUser          string
	PgDB            string
	PgHost          string
	PgPort          string
}

//LoadConfig reads configuration from .env  file
func LoadConfig(path string) (config *Config, err error) {

	config = new(Config)
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
		logger.LogErr(err)
		// is there a sense to fail if you don't use anything config?
	}

	//Parsing config vals from file (second step)
	err = viper.Unmarshal(config)

	viper.AutomaticEnv()
	viper.BindEnv("MainAppAdr")
	os.Setenv("MAINAPPADR", "http://localhost:5050")
	config.MainAppAdr = viper.GetString("MainAppAdr")
	return

}
