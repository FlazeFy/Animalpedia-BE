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
)
