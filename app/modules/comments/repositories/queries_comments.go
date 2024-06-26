package repositories

import (
	"app/modules/comments/models"
	"app/packages/builders"
	"app/packages/database"
	"app/packages/helpers/generator"
	"app/packages/helpers/response"
	"app/packages/utils/pagination"
	"math"
	"net/http"
)

func GetAllCommentBySlug(page, pageSize int, path string, types string, slug string) (response.Response, error) {
	// Declaration
	var obj models.GetAllComment
	var arrobj []models.GetAllComment
	var res response.Response
	var baseTable = "comments"
	var sqlStatement string

	// Query builder
	propsTemplate := builders.GetTemplateSelect("properties_time", &baseTable, nil)
	joinTemplate := builders.GetTemplateJoin("total", baseTable, "context_id", types, "id", false)
	activeTemplate := builders.GetTemplateLogic("active")
	order := builders.GetTemplateOrder("permanent_data", baseTable, "updated_at DESC")

	sqlStatement = "SELECT comments_body, " + propsTemplate + " " +
		"FROM " + baseTable + " " +
		joinTemplate + " " +
		"WHERE " + baseTable + "." + activeTemplate + " " +
		"AND context_type='" + types + "' AND " + types + "_slug='" + slug + "' " +
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
			&obj.CommentBody,
			&obj.CreatedAt,
			&obj.CreatedBy,
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
