package models

type (
	GetDictionaryByType struct {
		ID      int    `json:"id"`
		DctName string `json:"dictionary_name"`
		DctType string `json:"dictionary_type"`
	}
	PostDictionaryByType struct {
		DctType string `json:"dictionaries_type"`
		DctName string `json:"dictionaries_name"`
	}
)
