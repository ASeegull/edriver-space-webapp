package auth

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/ASeegull/edriver-space-webapp/config"
	"github.com/ASeegull/edriver-space-webapp/logger"
	"github.com/ASeegull/edriver-space-webapp/model"
)

// LoginProcced() func delivers login request to main app and returns its response
func LoginProceed(data model.SingInData, config *config.Config) string {
	// Forming request body
	requestbody := strings.NewReader(`
	{
		"Email" : "` + data.Email + `"
		"Password" : "` + data.Password + `"
	}`)

	// Sending request to main app
	response, err := http.Post(config.MainAppAdr+config.SignInURL, "application/json", requestbody)
	if err != nil {
		logger.LogErr(err)
	}

	defer response.Body.Close()

	var content []byte

	content, err = io.ReadAll(response.Body)
	if err != nil {
		logger.LogErr(err)
	}

	return string(content)
}

// LogoutProceed() func delivers logout request to main app and returns an error if logout failed.
func LogoutProceed(accesstoken string, config *config.Config) error {
	// Forming request body
	requestbody := strings.NewReader(`
	{
		"AccessToken" : "` + accesstoken + `"
	}`)

	// Sending request to main app
	response, err := http.Post(config.MainAppAdr+config.SignOutURL, "application/json", requestbody)
	if err != nil {
		logger.LogErr(err)
	}

	defer response.Body.Close()

	var content []byte

	content, err = io.ReadAll(response.Body)
	if err != nil {
		logger.LogErr(err)
	}

	var authErr error
	if string(content) != "OK" {
		authErr = fmt.Errorf(string(content))
	}

	return authErr
}

// RefreshProceed() func delivers renew request (for refreshing tokens) to main app and returns an error if process failed.
func RefreshProceed(refreshtoken string, config *config.Config) error {
	// Sending request to main app
	response, err := http.Get(config.MainAppAdr + config.RefreshTokenURL)
	if err != nil {
		logger.LogErr(err)
	}

	defer response.Body.Close()

	var content []byte

	content, err = io.ReadAll(response.Body)
	if err != nil {
		logger.LogErr(err)
	}

	var authErr error
	if string(content) != "OK" {
		authErr = fmt.Errorf(string(content))
	}

	return authErr
}
