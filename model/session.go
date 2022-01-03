package model

// Session fields may be not final
type Session struct {
	ID        int    `json:"id"`
	UserIP    string `json:"userip"`
	UserLogin string `json:"userlogin"`
	Active    bool   `json:"sessionactive"`
}
