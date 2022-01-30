package api_client

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/ASeegull/edriver-space-webapp/config"
	"github.com/ASeegull/edriver-space-webapp/model"
	log "github.com/sirupsen/logrus"
)

type UsersRequests struct {
	HttpClient *http.Client
	cfg        *config.Config
	logger     *log.Logger
}

func NewUsersRequests(cfg *config.Config) *UsersRequests {
	return &UsersRequests{
		HttpClient: &http.Client{
			Timeout: time.Second * 10,
		},
		cfg:    cfg,
		logger: log.New(),
	}
}

func (u *UsersRequests) SignIn(input model.SignInInput) (model.ApiResponseWithCookies, error) {

	requestBody, err := u.createRequestBody(input)
	if err != nil {
		return model.ApiResponseWithCookies{}, err
	}

	req, err := http.NewRequest("POST", u.cfg.BaseUrl+u.cfg.UsersSignInUrl, requestBody)
	if err != nil {
		return model.ApiResponseWithCookies{}, err
	}

	return u.getTokensAndCookiesFromResponse(req)
}

func (u *UsersRequests) SignUp(input model.SignUpInput) (model.ApiResponseWithCookies, error) {

	requestBody, err := u.createRequestBody(input)
	if err != nil {
		return model.ApiResponseWithCookies{}, err
	}

	req, err := http.NewRequest("POST", u.cfg.BaseUrl+u.cfg.UsersSignUpUrl, requestBody)
	if err != nil {
		return model.ApiResponseWithCookies{}, err
	}

	return u.getTokensAndCookiesFromResponse(req)
}

func (u *UsersRequests) RefreshTokens(cookie *http.Cookie) (model.ApiResponseWithCookies, error) {

	req, err := http.NewRequest("GET", u.cfg.BaseUrl+u.cfg.UsersRefreshTokensUrl, nil)
	if err != nil {
		return model.ApiResponseWithCookies{}, err
	}

	req.AddCookie(cookie)

	return u.getTokensAndCookiesFromResponse(req)
}

func (u *UsersRequests) getTokensAndCookiesFromResponse(req *http.Request) (model.ApiResponseWithCookies, error) {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err := u.HttpClient.Do(req)
	if err != nil {
		return model.ApiResponseWithCookies{}, err
	}

	defer closeWithLogOnErr(u.logger, resp.Body)

	if resp.StatusCode != http.StatusOK {
		byteResponseBody, _ := io.ReadAll(resp.Body)

		return model.ApiResponseWithCookies{
			ApiResponse: model.ApiResponse{
				StatusCode: resp.StatusCode,
				Body:       string(byteResponseBody),
			},
			Cookies: nil,
		}, nil
	}

	var tokens model.Tokens

	if err := u.unmarshalJson(resp.Body, &tokens); err != nil {
		return model.ApiResponseWithCookies{}, err
	}

	return model.ApiResponseWithCookies{
		ApiResponse: model.ApiResponse{
			StatusCode: resp.StatusCode,
			Body:       tokens,
		},
		Cookies: resp.Cookies(),
	}, nil
}

func (u *UsersRequests) SignOut(cookie *http.Cookie) (model.ApiResponse, error) {

	req, err := http.NewRequest("POST", u.cfg.BaseUrl+u.cfg.UsersSignOutUrl, nil)
	if err != nil {
		return model.ApiResponse{}, err
	}

	req.AddCookie(cookie)

	resp, err := u.HttpClient.Do(req)
	if err != nil {
		return model.ApiResponse{}, err
	}

	defer closeWithLogOnErr(u.logger, resp.Body)

	byteRespBody, _ := io.ReadAll(resp.Body)

	return model.ApiResponse{StatusCode: resp.StatusCode, Body: string(byteRespBody)}, nil
}

func (u *UsersRequests) GetFines(jwtHeader string) (model.ApiResponse, error) {

	req, err := http.NewRequest("GET", u.cfg.BaseUrl+u.cfg.UsersGetFinesUrl, nil)
	if err != nil {
		return model.ApiResponse{}, err
	}

	req.Header.Set("Authorization", jwtHeader)

	resp, err := u.HttpClient.Do(req)
	if err != nil {
		return model.ApiResponse{}, err
	}

	if resp.StatusCode != http.StatusOK {
		byteRespBody, _ := io.ReadAll(resp.Body)

		return model.ApiResponse{StatusCode: resp.StatusCode, Body: string(byteRespBody)}, nil
	}

	defer closeWithLogOnErr(u.logger, resp.Body)

	var fines model.Fines

	if err := u.unmarshalJson(resp.Body, &fines); err != nil {
		return model.ApiResponse{}, err
	}

	return model.ApiResponse{
		StatusCode: resp.StatusCode,
		Body:       fines,
	}, nil
}

func (u *UsersRequests) AddDriverLicense(input model.AddDriverLicenceInput, jwtHeader string) (model.ApiResponse, error) {

	reqBody, err := u.createRequestBody(input)
	if err != nil {
		return model.ApiResponse{}, err
	}

	req, err := http.NewRequest("POST", u.cfg.BaseUrl+u.cfg.UsersAddDriverLicenceUrl, reqBody)
	if err != nil {
		return model.ApiResponse{}, err
	}
	req.Header.Set("Authorization", jwtHeader)

	resp, err := u.HttpClient.Do(req)
	if err != nil {
		return model.ApiResponse{}, err
	}

	byteRespBody, _ := io.ReadAll(resp.Body)

	return model.ApiResponse{
		StatusCode: resp.StatusCode,
		Body:       string(byteRespBody),
	}, nil
}

func (u *UsersRequests) createRequestBody(i interface{}) (*strings.Reader, error) {
	byteBody, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}

	requestBody := strings.NewReader(string(byteBody))

	return requestBody, nil
}

func (u *UsersRequests) unmarshalJson(r io.Reader, i interface{}) error {
	byteRespBody, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(byteRespBody, i); err != nil {
		return err
	}
	return nil
}

// closeWithLogOnErr is making sure we log every error, even those from the best effort tiny closers.
func closeWithLogOnErr(log *log.Logger, closer io.Closer) {
	err := closer.Close()
	if err == nil {
		return
	}

	log.Warn(err)
}
