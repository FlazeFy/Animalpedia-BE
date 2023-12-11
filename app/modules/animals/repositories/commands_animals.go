package repositories

import (
	"app/packages/builders"
	"app/packages/database"
	"app/packages/helpers/generator"
	"app/packages/helpers/response"
	"net/http"
	"time"
)

func SoftDelAnimalBySlug(slug string) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "animals"
	var sqlStatement string
	dt := time.Now().Format("2006-01-02 15:04:05")

	// Command
	sqlStatement = builders.GetTemplateCommand("soft_delete", baseTable)

	// Exec
	con := database.CreateCon()
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(dt, "1", slug)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateCommandMsg(baseTable, "delete", int(rowsAffected))
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func SoftDelNewsBySlug(slug string) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "news"
	var sqlStatement string
	dt := time.Now().Format("2006-01-02 15:04:05")

	// Command
	sqlStatement = builders.GetTemplateCommand("soft_delete", baseTable)

	// Exec
	con := database.CreateCon()
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(dt, "1", slug)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateCommandMsg(baseTable, "delete", int(rowsAffected))
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}
