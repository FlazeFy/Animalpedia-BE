package models

type (
	GetAnimalHeaders struct {
		AnimalSlug      string `json:"animals_slug"`
		AnimalName      string `json:"animals_name"`
		AnimalLatinName string `json:"animals_latin_name"`
		AnimalImgUrl    string `json:"animals_image_url"`
		AnimalRegion    string `json:"animals_region"`
		AnimalZone      string `json:"animals_zone"`
		AnimalStatus    string `json:"animals_status"`
		AnimalCategory  string `json:"animals_category"`
	}
	GetNewsHeaders struct {
		NewsSlug     string `json:"news_slug"`
		NewsName     string `json:"news_name"`
		NewsTag      string `json:"news_tag"`
		NewsBody     string `json:"news_body"`
		NewsTimeRead int    `json:"news_time_read"`
		NewsImgUrl   string `json:"news_image_url"`

		// Props
		CreatedAt string `json:"created_at"`
		CreatedBy string `json:"created_by"`
	}
)
