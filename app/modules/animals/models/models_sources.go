package models

type (
	GetSources struct {
		Id          string `json:"id"`
		SourceTitle string `json:"sources_title"`
		SourceUrl   string `json:"sources_url"`
	}
)
