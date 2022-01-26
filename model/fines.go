package model

type Fines struct {
	DriversFines []DriversFine `json:"drivers_fines"`
	CarsFines    []CarsFine    `json:"cars_fines"`
}

type DriversFine struct {
	Id                        string `json:"id"`
	LicenceNumber             string `json:"licence_number"`
	FineNum                   string `json:"fine_num"`
	DataAndTime               string `json:"data_and_time"`
	Place                     string `json:"place"`
	FileLawArticle            string `json:"file_law_article"`
	Price                     int    `json:"price"`
	VehicleRegistrationNumber string `json:"vehicle_registration_number"`
}

type CarsFine struct {
	Id                        string `json:"id"`
	VehicleRegistrationNumber string `json:"vehicle_registration_number"`
	FineNum                   string `json:"fine_num"`
	DataAndTime               string `json:"data_and_time"`
	Place                     string `json:"place"`
	FileLawArticle            string `json:"file_law_article"`
	Price                     int    `json:"price"`
	Info                      string `json:"info"`
	ImdUrl                    string `json:"imd_url"`
}
