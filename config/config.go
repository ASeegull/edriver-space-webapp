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
	UsrNotFoundMsg  string `mapstructure:"USER_DOESNT_EXCIST_MSG"`
	WrongPassMsg    string `mapstructure:"WRONG_PASS_MSG"`
	MainPageTitle   string `mapstructure:"MAIN_PAGE_TITLE"`
	PanelPageTitle  string `mapstructure:"PANEL_PAGE_TITLE"`
	MainAppAdr      string
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

	viper.AutomaticEnv()
	viper.BindEnv("MainAppAdr")

	// Temporary setting MAINAPPADR manually for correct work with Login plug
	os.Setenv("MAINAPPADR", "http://localhost:5050")

	config.MainAppAdr = viper.GetString("MainAppAdr")
	return

}
