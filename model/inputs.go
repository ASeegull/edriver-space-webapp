package model

type SignInInput struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type SignUpInput struct {
	Firstname string `json:"firstname" form:"firstname"`
	Lastname  string `json:"lastname" form:"lastname"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
}

type AddDriverLicenceInput struct {
	IndividualTaxNumber string `json:"individual_tax_number" form:"individual_tax_number"`
}
