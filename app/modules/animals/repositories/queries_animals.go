package repositories

import (
	"app/modules/animals/models"
	"app/packages/builders"
	"app/packages/database"
	"app/packages/helpers/generator"
	"app/packages/helpers/response"
	"app/packages/utils/pagination"
	"math"
	"net/http"
)

func GetAllAnimalHeaders(page, pageSize int, path string, ord string) (response.Response, error) {
	// Declaration
	var obj models.GetAnimalHeaders
	var arrobj []models.GetAnimalHeaders
	var res response.Response
	var baseTable = "animals"
	var sqlStatement string

	// Query builder
	selectTemplate := builders.GetTemplateSelect("content_info", &baseTable, nil)
	order := builders.GetTemplateOrder("dynamic_data", baseTable, "animals_name")

	sqlStatement = "SELECT " + selectTemplate + ", animals_latin_name, animals_img_url, animals_region, animals_zone, animals_status, animals_category " +
		"FROM " + baseTable + " " +
		"ORDER BY " + order + " " +
		"LIMIT ? OFFSET ?"

	// Exec
	con := database.CreateCon()
	offset := (page - 1) * pageSize
	rows, err := con.Query(sqlStatement, pageSize, offset)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	// Map
	for rows.Next() {
		err = rows.Scan(
			&obj.AnimalSlug,
			&obj.AnimalName,
			&obj.AnimalLatinName,
			&obj.AnimalImgUrl,
			&obj.AnimalRegion,
			&obj.AnimalZone,
			&obj.AnimalStatus,
			&obj.AnimalCategory,
		)

		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	// Page
	total, err := builders.GetTotalCount(con, baseTable, nil)
	if err != nil {
		return res, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	pagination := pagination.BuildPaginationResponse(page, pageSize, total, totalPages, path)

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg(baseTable, total)
	if total == 0 {
		res.Data = nil
	} else {
		res.Data = map[string]interface{}{
			"current_page":   page,
			"data":           arrobj,
			"first_page_url": pagination.FirstPageURL,
			"from":           pagination.From,
			"last_page":      pagination.LastPage,
			"last_page_url":  pagination.LastPageURL,
			"links":          pagination.Links,
			"next_page_url":  pagination.NextPageURL,
			"path":           pagination.Path,
			"per_page":       pageSize,
			"prev_page_url":  pagination.PrevPageURL,
			"to":             pagination.To,
			"total":          total,
		}
	}

	return res, nil
}
