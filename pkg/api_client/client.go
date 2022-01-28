package api_client

import (
	"github.com/ASeegull/edriver-space-webapp/config"
	"github.com/ASeegull/edriver-space-webapp/model"
	"net/http"
)

type Users interface {
	SignIn(input model.SignInInput) (model.ApiResponseWithCookies, error)
	SignUp(input model.SignUpInput) (model.ApiResponseWithCookies, error)
	RefreshTokens(cookie *http.Cookie) (model.ApiResponseWithCookies, error)
	SignOut(cookie *http.Cookie) (model.ApiResponse, error)
	GetFines(jwtHeader string) (model.ApiResponse, error)
	AddDriverLicense(input model.AddDriverLicenceInput, jwtHeader string) (model.ApiResponse, error)
}

type ApiClient struct {
	Users Users
}

func NewApiClient(cfg *config.Config) *ApiClient {
	return &ApiClient{
		Users: NewUsersRequests(cfg),
	}
}
