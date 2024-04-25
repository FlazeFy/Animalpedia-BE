package repositories

import (
	"app/modules/animals/models"
	"app/packages/builders"
	"app/packages/database"
	"app/packages/helpers/auth"
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
	animalDesc := data.FormValue("animals_desc")

	// Command builder
	sqlStatement = "UPDATE " + baseTable + " SET animals_name= ?, animals_latin_name= ?, animals_region= ?, " +
		"animals_zone= ?, animals_status= ?, animals_category= ?, animals_desc= ?, updated_at= ?, updated_by= ? " +
		"WHERE animals_slug= ?"

	// Exec
	con := database.CreateCon()
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(animalName, animalLatinName, animalRegion, animalZone, animalStatus, animalCategory, animalDesc, dt, "1", slug)
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

func PostAnimal(d models.PostAnimal) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "animals"
	var sqlStatement string
	dt := time.Now().Format("2006-01-02 15:04:05")

	// Data
	id := uuid.Must(uuid.NewRandom())
	slug := generator.GetSlug(d.AnimalName)

	// Command builder
	sqlStatement = "INSERT INTO " + baseTable + " (id, animals_slug, animals_name, animals_desc, animals_latin_name, animals_img_url, animals_region, animals_zone, animals_status, animals_category, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by) " +
		"VALUES (?,?,?,?,?,?,?,?,?,?,?,?,null,null,null,null)"

	// Exec
	con := database.CreateCon()
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id, slug, d.AnimalName, d.AnimalDesc, d.AnimalLatinName, d.AnimalImgUrl, d.AnimalRegion, d.AnimalZone, d.AnimalStatus, d.AnimalCategory, dt, "1")
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
	res.Data = map[string]interface{}{
		"rows_affected": rowsAffected,
		"id":            id,
		"data":          d,
	}

	return res, nil
}

func PostNews(d models.PostNews) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "news"
	var sqlStatement string
	dt := time.Now().Format("2006-01-02 15:04:05")

	// Data
	id := uuid.Must(uuid.NewRandom())
	slug := generator.GetSlug(d.NewsName)

	// Command builder
	sqlStatement = "INSERT INTO " + baseTable + " (id, news_slug, news_name, news_tag, news_body, news_time_read, news_img_url, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by) " +
		"VALUES (?,?,?,?,?,?,?,?,?,null,null,null,null)"

	// Exec
	con := database.CreateCon()
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id, slug, d.NewsName, d.NewsTag, d.NewsBody, d.NewsTimeRead, d.NewsImgUrl, dt, "1")
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
	res.Data = map[string]interface{}{
		"rows_affected": rowsAffected,
		"id":            id,
		"data":          d,
	}

	return res, nil
}

func UpdateNewsBySlug(slug string, data echo.Context) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "news"
	var sqlStatement string
	dt := time.Now().Format("2006-01-02 15:04:05")
	_, userId := auth.GetTokenHeader(data)

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

	result, err := stmt.Exec(newsName, newsBody, newsTimeRead, dt, userId, slug)
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

func RecoverAnimalBySlug(slug, token string) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "animals"
	var sqlStatement string

	// Command builder
	sqlStatement = builders.GetTemplateCommand("recover", baseTable, baseTable+"_slug")

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
	res.Message = generator.GenerateCommandMsg(baseTable, "recover", int(rowsAffected))
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func RecoverNewsBySlug(slug, token string) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "news"
	var sqlStatement string

	// Command builder
	sqlStatement = builders.GetTemplateCommand("recover", baseTable, baseTable+"_slug")

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
	res.Message = generator.GenerateCommandMsg(baseTable, "recover", int(rowsAffected))
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}
