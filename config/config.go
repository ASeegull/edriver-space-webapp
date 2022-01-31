package config

import (
	"os"

	"github.com/spf13/viper"

	"github.com/ASeegull/edriver-space-webapp/logger"
)

//Config struct stores all configuration values for project using Viper
type Config struct {
	BaseUrl                  string `mapstructure:"BASE_URL"`
	UsersSignInUrl           string `mapstructure:"USERS_SIGN_IN_URL"`
	UsersSignUpUrl           string `mapstructure:"USERS_SIGN_UP_URL"`
	UsersSignOutUrl          string `mapstructure:"USERS_SIGN_OUT_URL"`
	UsersRefreshTokensUrl    string `mapstructure:"USERS_REFRESH_TOKENS_URL"`
	UsersAddDriverLicenceUrl string `mapstructure:"USERS_ADD_DRIVER_LICENCE_URL"`
	UsersGetFinesUrl         string `mapstructure:"USERS_GET_FINES_URL"`
	CookieName               string `mapstructure:"COOKIE_NAME"`
	UsrNotFoundMsg           string `mapstructure:"USER_DOESNT_EXCIST_MSG"`
	WrongPassMsg             string `mapstructure:"WRONG_PASS_MSG"`
	MainPageTitle            string `mapstructure:"MAIN_PAGE_TITLE"`
	PanelPageTitle           string `mapstructure:"PANEL_PAGE_TITLE"`
	MainAppAdr               string
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
	}

	//Parsing config vals from file (second step)
	err = viper.Unmarshal(config)

	return

}
