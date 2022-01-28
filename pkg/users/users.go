package users

import (
	"io"
	"net/http"
	"strings"

	"github.com/ASeegull/edriver-space-webapp/config"
	"github.com/ASeegull/edriver-space-webapp/logger"
	"github.com/ASeegull/edriver-space-webapp/model"
)

// UpdateUserInfo() func delivers requests for updating user data to main app and returns its response
func UpdateUserInfo(newinfo model.User, config *config.Config) string {
	// Forming request body
	requestbody := strings.NewReader(`
	{
		"Login" : "` + newinfo.Login + `",	
		"Name" : "` + newinfo.FullName + `",	
		"DateOfBirth" : "` + newinfo.DateOfBirth + `",	
		"PlaceOfBirth" : "` + newinfo.PlaceOfBirth + `",
		"LicenseNumber" : "` + newinfo.LicenseNumber + `",
		"LicenseIssueDate" : "` + newinfo.DateOfIssue + `",
		"LicenseExpireDate" : "` + newinfo.ExpireDate + `",
		"IndividualTaxNumber" : "` + newinfo.IndividualTaxNumber + `",
		"DriverCategory" : "` + newinfo.Category + `",
		"CategoryIssueDate" : "` + newinfo.CategoryIssuingDate + `",
		"CategoryExpireDate" : "` + newinfo.CategoryExpire + `",
	}`)

	// Sending request to main app
	response, err := http.Post(config.MainAppAdr+config.UpdateUserURL, "application/json", requestbody)
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
