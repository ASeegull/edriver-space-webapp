package model

// Session fields may be not final
type Session struct {
	UserIP       string `json:"userip"`
	UserLogin    string `json:"userlogin"`
	AccessToken  string `json:"accesstoken"`
	RefreshToken string `json:"refreshtoken"`
	Active       bool   `json:"sessionactive"`
}
