package repositories

import (
	"app/packages/builders"
	"app/packages/database"
	"app/packages/helpers/generator"
	"app/packages/helpers/response"
	"net/http"

	"github.com/labstack/echo"
)

func HardDelDictionaryById(id string) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "dictionaries"
	var sqlStatement string

	// Command builder
	sqlStatement = builders.GetTemplateCommand("hard_delete", baseTable, "id")

	// Exec
	con := database.CreateCon()
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateCommandMsg(baseTable, "permanently delete", int(rowsAffected))
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func PostDictionary(data echo.Context) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "dictionaries"
	var sqlStatement string

	// Data
	dctName := data.FormValue("dictionaries_name")
	dctSlug := generator.GetSlug(dctName)

	// Command builder
	sqlStatement = "INSERT INTO " + baseTable + " (id, dictionaries_type, dictionaries_name) " +
		"VALUES (null,?,?)"

	// Exec
	con := database.CreateCon()
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(dctName, dctSlug)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateCommandMsg(baseTable, "create", int(rowsAffected))
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}
