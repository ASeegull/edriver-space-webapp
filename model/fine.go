package model

// Car fields may be not final
type Fine struct {
	VIN         string `json:"VIN"`
	NumberPlate string `json:"numberplate"`
	Date        string `json:"issue_date"`
	Place       string `json:"place"`
	Violation   string `json:"violation"`
	Ammount     string `json:"ammount"`
}
