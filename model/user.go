package model

// User fields may be not final
type User struct {
	ID                  int    `json:"id"`
	Login               string `json:"login"`
	LicenseNumber       string `json:"license_number"`
	DateOfIssue         string `json:"date_of_issue"`
	ExpireDate          string `json:"expire_date"`
	IndividualTaxNumber string `json:"individual_tax_number"`

	Category            string `json:"category"`
	CategoryIssuingDate string `json:"category_issuing_date"`
	CategoryExpire      string `json:"category_expire"`

	FullName     string `json:"full_name"`
	DateOfBirth  string `json:"date_of_birth"`
	PlaceOfBirth string `json:"place_of_birth"`

	Restrictions string `json:"restrictions"`
}
