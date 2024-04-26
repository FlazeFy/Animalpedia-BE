package models

type (
	GetAnimalCountries struct {
		Id          string `json:"id"`
		CountryCode string `json:"countries_code"`
		CountryName string `json:"countries_name"`
		Total       int    `json:"total"`
	}
	GetCountries struct {
		CountryCode string `json:"countries_code"`
		CountryName string `json:"countries_name"`
	}
)
