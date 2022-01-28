package fines

import (
	"io"
	"net/http"
	"strings"

	"github.com/ASeegull/edriver-space-webapp/config"
	"github.com/ASeegull/edriver-space-webapp/logger"
)

// GetDriverFines() func delivers requests for getting fines by driver license to main app and returns its response
func GetDriverFines(license string, config *config.Config) string {
	// Forming request body
	requestbody := strings.NewReader(`
	{
		"LicenseNumber" : "` + license + `",	
	}`)

	// Sending request to main app
	response, err := http.Post(config.MainAppAdr+config.GetDriverFinesURL, "application/json", requestbody)
	if err != nil {
		logger.LogErr(err)
	}
	defer response.Body.Close()

	// Saving the response
	var content []byte
	content, err = io.ReadAll(response.Body)
	if err != nil {
		logger.LogErr(err)
	}

	return string(content)
}

// GetCarFines() func delivers requests for getting fines by car number plate to main app and returns its response
func GetCarFines(numberplate string, config *config.Config) string {
	// Forming request body
	requestbody := strings.NewReader(`
	{
		"CarNumber" : "` + numberplate + `",	
	}`)

	// Sending request to main app
	response, err := http.Post(config.MainAppAdr+config.GetCarFinesURL, "application/json", requestbody)
	if err != nil {
		logger.LogErr(err)
	}
	defer response.Body.Close()

	// Saving the response
	var content []byte
	content, err = io.ReadAll(response.Body)
	if err != nil {
		logger.LogErr(err)
	}

	return string(content)
}
