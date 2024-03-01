package seeders

import (
	"app/modules/systems/models"
	"app/modules/systems/repositories"
	"app/packages/helpers/response"
	"fmt"

	"github.com/bxcodec/faker/v3"
)

func SeedTags(total int, showRes bool) {
	var obj models.PostTag
	idx := 0
	var logs string

	for idx < total {
		// Data
		name := faker.Word()
		obj.TagName = name

		result, err := repositories.PostTag(obj)
		if err != nil {
			fmt.Println(err.Error())
		}

		if showRes {
			fmt.Println(result.Data)
			if strData, ok := result.Data.(string); ok {
				logs += strData + "\n"
			}
		}
		idx++
	}

	if showRes {
		response.ResponsePrinter("txt", "seeder_tags", logs)
	}
}
