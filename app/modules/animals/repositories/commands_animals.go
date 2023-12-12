package repositories

import (
	"app/packages/builders"
	"app/packages/database"
	"app/packages/helpers/generator"
	"app/packages/helpers/response"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func SoftDelAnimalBySlug(slug string) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "animals"
	var sqlStatement string
	dt := time.Now().Format("2006-01-02 15:04:05")

	// Command builder
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

	// Command builder
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

func UpdateAnimalBySlug(slug string, data echo.Context) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "animals"
	var sqlStatement string
	dt := time.Now().Format("2006-01-02 15:04:05")

	// Data
	animalName := data.FormValue("animals_name")
	animalLatinName := data.FormValue("animals_latin_name")
	animalRegion := data.FormValue("animals_region")
	animalZone := data.FormValue("animals_zone")
	animalStatus := data.FormValue("animals_status")
	animalCategory := data.FormValue("animals_category")

	// Command builder
	sqlStatement = "UPDATE " + baseTable + " SET animals_name= ?, animals_latin_name= ?, animals_region= ?, " +
		"animals_zone= ?, animals_status= ?, animals_category= ?, updated_at= ?, updated_by= ? " +
		"WHERE animals_slug= ?"

	// Exec
	con := database.CreateCon()
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(animalName, animalLatinName, animalRegion, animalZone, animalStatus, animalCategory, dt, "1", slug)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateCommandMsg(baseTable, "update", int(rowsAffected))
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}
