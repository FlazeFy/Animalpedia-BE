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
	PostAnimalCountry struct {
		AnimalId    string `json:"animals_id"`
		CountryCode string `json:"countries_code"`
		Total       string `json:"total"`
	}
)
