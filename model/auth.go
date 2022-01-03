package model

// SignInData fields may be not final
type SingInData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthData struct {
	AccessToken  string `json:"accesstoken"`
	RefreshToken string `json:"refreshtoken"`
}
