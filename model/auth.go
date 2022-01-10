package model

// SignInData fields may be not final
type SingInData struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

// AuthData fields may be not final
type AuthData struct {
	AccessToken  string `json:"accesstoken"`
	RefreshToken string `json:"refreshtoken"`
}
