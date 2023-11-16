package repositories

import (
	"app/modules/systems/models"
	"app/packages/builders"
	"app/packages/database"
	"app/packages/helpers/generator"
	"app/packages/helpers/response"
	"app/packages/utils/pagination"
	"math"
	"net/http"
	"strconv"
)

func GetDictionaryByType(page, pageSize int, path string, dctType string) (response.Response, error) {
	// Declaration
	var obj models.GetDictionaryByType
	var arrobj []models.GetDictionaryByType
	var res response.Response
	var baseTable = "dictionaries"
	var sqlStatement string

	// Converted Column
	var Id string

	// Query builder
	where := "dictionaries_type = '" + dctType + "' "
	order := "dictionaries_name DESC "

	sqlStatement = "SELECT id, dictionaries_type, dictionaries_name " +
		"FROM " + baseTable + " " +
		"WHERE " + where +
		"ORDER BY " + order +
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
			&Id,
			&obj.DctName,
			&obj.DctType,
		)

		if err != nil {
			return res, err
		}

		// Converted
		intID, err := strconv.Atoi(Id)
		if err != nil {
			return res, err
		}

		obj.ID = intID

		arrobj = append(arrobj, obj)
	}

	// Page
	total, err := builders.GetTotalCount(con, baseTable, &where)
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
