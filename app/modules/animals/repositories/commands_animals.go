package repositories

import (
	"app/packages/builders"
	"app/packages/database"
	"app/packages/helpers/generator"
	"app/packages/helpers/response"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

func SoftDelAnimalBySlug(slug string) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "animals"
	var sqlStatement string
	dt := time.Now().Format("2006-01-02 15:04:05")

	// Command builder
	sqlStatement = builders.GetTemplateCommand("soft_delete", baseTable, baseTable+"_slug")

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

func HardDelAnimalBySlug(slug string) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "animals"
	var sqlStatement string

	// Command builder
	sqlStatement = builders.GetTemplateCommand("hard_delete", baseTable, baseTable+"_slug")

	// Exec
	con := database.CreateCon()
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(slug)
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

func HardDelNewsBySlug(slug string) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "news"
	var sqlStatement string

	// Command builder
	sqlStatement = builders.GetTemplateCommand("hard_delete", baseTable, baseTable+"_slug")

	// Exec
	con := database.CreateCon()
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(slug)
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

func SoftDelNewsBySlug(slug string) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "news"
	var sqlStatement string
	dt := time.Now().Format("2006-01-02 15:04:05")

	// Command builder
	sqlStatement = builders.GetTemplateCommand("soft_delete", baseTable, baseTable+"_slug")

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

func PostAnimal(data echo.Context) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "animals"
	var sqlStatement string
	dt := time.Now().Format("2006-01-02 15:04:05")

	// Data
	id := uuid.Must(uuid.NewRandom())
	animalName := data.FormValue("animals_name")
	slug := generator.GetSlug(animalName)
	animalLatinName := data.FormValue("animals_latin_name")
	animalImgUrl := data.FormValue("animals_img_url")
	animalRegion := data.FormValue("animals_region")
	animalZone := data.FormValue("animals_zone")
	animalStatus := data.FormValue("animals_status")
	animalCategory := data.FormValue("animals_category")

	// Command builder
	sqlStatement = "INSERT INTO " + baseTable + " (id, animals_slug, animals_name, animals_latin_name, animals_img_url, animals_region, animals_zone, animals_status, animals_category, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by) " +
		"VALUES (?,?,?,?,?,?,?,?,?,?,?,null,null,null,null)"

	// Exec
	con := database.CreateCon()
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id, slug, animalName, animalLatinName, animalImgUrl, animalRegion, animalZone, animalStatus, animalCategory, dt, "1")
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

func PostNews(data echo.Context) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "news"
	var sqlStatement string
	dt := time.Now().Format("2006-01-02 15:04:05")

	// Data
	id := uuid.Must(uuid.NewRandom())
	newsName := data.FormValue("news_name")
	slug := generator.GetSlug(newsName)
	newsTag := data.FormValue("news_tag")
	newsBody := data.FormValue("news_body")
	newsTimeRead := data.FormValue("news_time_read")
	newsImgUrl := data.FormValue("news_image_url")

	// Command builder
	sqlStatement = "INSERT INTO " + baseTable + " (id, news_slug, news_name, news_tag, news_body, news_time_read, news_image_url, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by) " +
		"VALUES (?,?,?,?,?,?,?,?,?,null,null,null,null)"

	// Exec
	con := database.CreateCon()
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id, slug, newsName, newsTag, newsBody, newsTimeRead, newsImgUrl, dt, "1")
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

func UpdateNewsBySlug(slug string, data echo.Context) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "news"
	var sqlStatement string
	dt := time.Now().Format("2006-01-02 15:04:05")

	// Data
	newsName := data.FormValue("news_name")
	newsBody := data.FormValue("news_body")
	newsTimeRead := data.FormValue("news_time_read")

	// Command builder
	sqlStatement = "UPDATE " + baseTable + " SET news_name= ?, news_body= ?, news_time_read= ?, " +
		"updated_at= ?, updated_by= ? " +
		"WHERE news_slug= ?"

	// Exec
	con := database.CreateCon()
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(newsName, newsBody, newsTimeRead, dt, "1", slug)
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
