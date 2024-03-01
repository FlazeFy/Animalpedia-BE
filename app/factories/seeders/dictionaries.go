package seeders

import (
	"app/factories/dummies"
	"app/modules/systems/models"
	"app/modules/systems/repositories"
	"app/packages/helpers/generator"
	"app/packages/helpers/response"
	"fmt"

	"github.com/bxcodec/faker/v3"
)

func SeedDictionaries(total int, showRes bool) {
	var obj models.PostDictionaryByType
	idx := 0
	var logs string

	for idx < total {
		// Data
		obj.DctType = generator.GetSlug(dummies.DummyDctType())
		obj.DctName = faker.Word()

		result, err := repositories.PostDictionary(obj)
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
		response.ResponsePrinter("txt", "seeder_dictionaries", logs)
	}
}
