package models

type (
	GetAnimalHeaders struct {
		AnimalSlug      string `json:"animals_slug"`
		AnimalName      string `json:"animals_name"`
		AnimalLatinName string `json:"animals_latin_name"`
		AnimalImgUrl    string `json:"animals_img_url"`
		AnimalRegion    string `json:"animals_region"`
		AnimalZone      string `json:"animals_zone"`
		AnimalStatus    string `json:"animals_status"`
		AnimalCategory  string `json:"animals_category"`
	}
	GetAnimalDetail struct {
		AnimalSlug      string `json:"animals_slug"`
		AnimalName      string `json:"animals_name"`
		AnimalDesc      string `json:"animals_desc"`
		AnimalLatinName string `json:"animals_latin_name"`
		AnimalImgUrl    string `json:"animals_img_url"`
		AnimalRegion    string `json:"animals_region"`
		AnimalZone      string `json:"animals_zone"`
		AnimalStatus    string `json:"animals_status"`
		AnimalCategory  string `json:"animals_category"`

		// Props
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
		DeletedAt string `json:"deleted_at"`
	}
	PostAnimal struct {
		AnimalName      string `json:"animals_name"`
		AnimalDesc      string `json:"animals_desc"`
		AnimalLatinName string `json:"animals_latin_name"`
		AnimalImgUrl    string `json:"animals_img_url"`
		AnimalRegion    string `json:"animals_region"`
		AnimalZone      string `json:"animals_zone"`
		AnimalStatus    string `json:"animals_status"`
		AnimalCategory  string `json:"animals_category"`
	}
	UpdateAnimalHeaders struct {
		AnimalName      string `json:"animals_name"`
		AnimalLatinName string `json:"animals_latin_name"`
		AnimalRegion    string `json:"animals_region"`
		AnimalZone      string `json:"animals_zone"`
		AnimalStatus    string `json:"animals_status"`
		AnimalCategory  string `json:"animals_category"`
	}
	PostNews struct {
		NewsName     string `json:"news_name"`
		NewsTag      string `json:"news_tag"`
		NewsBody     string `json:"news_body"`
		NewsTimeRead int    `json:"news_time_read"`
		NewsImgUrl   string `json:"news_img_url"`
	}
	GetNewsHeaders struct {
		NewsSlug     string `json:"news_slug"`
		NewsName     string `json:"news_name"`
		NewsTag      string `json:"news_tag"`
		NewsBody     string `json:"news_body"`
		NewsTimeRead int    `json:"news_time_read"`
		NewsImgUrl   string `json:"news_img_url"`

		// Props
		CreatedAt string `json:"created_at"`
		CreatedBy string `json:"created_by"`
		UpdatedAt string `json:"updated_at"`
		DeletedAt string `json:"deleted_at"`
	}
	GetNewsSearch struct {
		NewsSlug string `json:"news_slug"`
		NewsName string `json:"news_name"`
		NewsBody string `json:"news_body"`
	}
)
