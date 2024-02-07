package models

type (
	GetTags struct {
		TagSlug string `json:"tags_slug"`
		TagName string `json:"tags_name"`
	}
	PostTag struct {
		TagName string `json:"tags_name"`
	}
)
