package repositories

import (
	"app/modules/animals/models"
	"app/packages/builders"
	"app/packages/database"
	"app/packages/helpers/converter"
	"app/packages/helpers/generator"
	"app/packages/helpers/response"
	"app/packages/utils/pagination"
	"database/sql"
	"math"
	"net/http"
	"strconv"
)

func GetAllAnimalHeaders(page, pageSize int, path string, ord string) (response.Response, error) {
	// Declaration
	var obj models.GetAnimalHeaders
	var arrobj []models.GetAnimalHeaders
	var res response.Response
	var baseTable = "animals"
	var sqlStatement string

	// Nullable column
	var AnimalImgUrl sql.NullString

	// Query builder
	selectTemplate := builders.GetTemplateSelect("content_info", &baseTable, nil)
	activeTemplate := builders.GetTemplateLogic("active")
	order := builders.GetTemplateOrder("dynamic_data", baseTable, "animals_name")

	sqlStatement = "SELECT " + selectTemplate + ", animals_latin_name, animals_img_url, animals_region, animals_zone, animals_status, animals_category " +
		"FROM " + baseTable + " " +
		"WHERE " + activeTemplate + " " +
		"ORDER BY " + order + " " + ord + " " +
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
			&AnimalImgUrl,
			&obj.AnimalRegion,
			&obj.AnimalZone,
			&obj.AnimalStatus,
			&obj.AnimalCategory,
		)

		if err != nil {
			return res, err
		}

		// Nullable
		obj.AnimalImgUrl = converter.CheckNullString(AnimalImgUrl)

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

func GetAnimalDetail(slug string) (response.Response, error) {
	// Declaration
	var obj models.GetAnimalDetail
	var res response.Response
	var baseTable = "animals"
	var sqlStatement string

	// Nullable column
	var UpdatedAt sql.NullString
	var DeletedAt sql.NullString
	var AnimalImgUrl sql.NullString

	// Query builder
	selectTemplate := builders.GetTemplateSelect("content_info", &baseTable, nil)

	sqlStatement = "SELECT " + selectTemplate + ", animals_desc, animals_latin_name, animals_img_url, animals_region, animals_zone, animals_status, animals_category, created_at, updated_at, deleted_at " +
		"FROM " + baseTable + " " +
		"WHERE animals_slug = '" + slug + "'"

	// Exec
	con := database.CreateCon()
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	// Map
	for rows.Next() {
		err = rows.Scan(
			&obj.AnimalSlug,
			&obj.AnimalName,
			&obj.AnimalDesc,
			&obj.AnimalLatinName,
			&AnimalImgUrl,
			&obj.AnimalRegion,
			&obj.AnimalZone,
			&obj.AnimalStatus,
			&obj.AnimalCategory,
			&obj.CreatedAt,
			&UpdatedAt,
			&DeletedAt,
		)

		// Nullable
		obj.UpdatedAt = converter.CheckNullString(UpdatedAt)
		obj.DeletedAt = converter.CheckNullString(DeletedAt)
		obj.AnimalImgUrl = converter.CheckNullString(AnimalImgUrl)

		if err != nil {
			return res, err
		}
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg(baseTable, 1)
	res.Data = obj

	return res, nil
}

func GetAllNewsHeaders(page, pageSize int, path string, ord string) (response.Response, error) {
	// Declaration
	var obj models.GetNewsHeaders
	var arrobj []models.GetNewsHeaders
	var res response.Response
	var baseTable = "news"
	var sqlStatement string

	// Converted Column
	var NewsTimeRead string

	// Query builder
	selectTemplate := builders.GetTemplateSelect("content_info", &baseTable, nil)
	propsTemplate := builders.GetTemplateSelect("properties_time", &baseTable, nil)
	activeTemplate := builders.GetTemplateLogic("active")
	order := builders.GetTemplateOrder("dynamic_data", baseTable, "news_name")

	sqlStatement = "SELECT " + selectTemplate + ", news_tag, news_body, news_time_read, news_img_url, " + propsTemplate + " " +
		"FROM " + baseTable + " " +
		"WHERE " + activeTemplate + " " +
		"ORDER BY " + order + " " + ord + " " +
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
			&obj.NewsSlug,
			&obj.NewsName,
			&obj.NewsTag,
			&obj.NewsBody,
			&NewsTimeRead,
			&obj.NewsImgUrl,
			&obj.CreatedAt,
			&obj.CreatedBy,
		)

		if err != nil {
			return res, err
		}

		// Converted
		intNewsTimeRead, err := strconv.Atoi(NewsTimeRead)
		if err != nil {
			return res, err
		}

		obj.NewsTimeRead = intNewsTimeRead

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

func GetNewsDetail(slug string) (response.Response, error) {
	// Declaration
	var obj models.GetNewsHeaders
	var res response.Response
	var baseTable = "news"
	var sqlStatement string

	// Converted Column
	var NewsTimeRead string

	// Nullable column
	var UpdatedAt sql.NullString
	var DeletedAt sql.NullString

	// Query builder
	selectTemplate := builders.GetTemplateSelect("content_info", &baseTable, nil)
	propsTemplate := builders.GetTemplateSelect("properties_time", &baseTable, nil)

	sqlStatement = "SELECT " + selectTemplate + ", news_tag, news_body, news_time_read, news_img_url, " + propsTemplate + ", updated_at, deleted_at " +
		"FROM " + baseTable + " " +
		"WHERE news_slug = '" + slug + "'"

	// Exec
	con := database.CreateCon()
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	// Map
	for rows.Next() {
		err = rows.Scan(
			&obj.NewsSlug,
			&obj.NewsName,
			&obj.NewsTag,
			&obj.NewsBody,
			&NewsTimeRead,
			&obj.NewsImgUrl,
			&obj.CreatedAt,
			&obj.CreatedBy,
			&UpdatedAt,
			&DeletedAt,
		)

		if err != nil {
			return res, err
		}

		// Converted
		intNewsTimeRead, err := strconv.Atoi(NewsTimeRead)
		if err != nil {
			return res, err
		}

		obj.NewsTimeRead = intNewsTimeRead

		// Nullable
		obj.UpdatedAt = converter.CheckNullString(UpdatedAt)
		obj.DeletedAt = converter.CheckNullString(DeletedAt)
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg(baseTable, 1)
	res.Data = obj

	return res, nil
}

func GetNewsByTags(page, pageSize int, path string, slug string) (response.Response, error) {
	// Declaration
	var obj models.GetNewsSearch
	var arrobj []models.GetNewsSearch
	var res response.Response
	var baseTable = "news"
	var sqlStatement string

	// Query builder
	selectTemplate := builders.GetTemplateSelect("content_info", &baseTable, nil)
	activeTemplate := builders.GetTemplateLogic("active")
	order := builders.GetTemplateOrder("dynamic_data", baseTable, "news_name")
	filterTemplate := builders.GetTemplateCommand("filter_tag", baseTable, slug)

	sqlStatement = "SELECT " + selectTemplate + ", news_body " +
		"FROM " + baseTable + " " +
		"WHERE " + activeTemplate + " AND " + filterTemplate + " " +
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
			&obj.NewsSlug,
			&obj.NewsName,
			&obj.NewsBody,
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

func GetSourcesByContextSlug(types string, slug string) (response.Response, error) {
	// Declaration
	var obj models.GetSources
	var arrobj []models.GetSources
	var res response.Response
	var baseTable = "sources"
	var sqlStatement string

	// Query builder
	joinTemplate := builders.GetTemplateJoin("total", baseTable, "context_id", types, "id", false)
	order := builders.GetTemplateOrder("permanent_data", baseTable, "sources_title DESC")

	sqlStatement = "SELECT " + baseTable + ".id, sources_title, sources_url " +
		"FROM " + baseTable + " " +
		joinTemplate + " " +
		"WHERE context_type='" + types + "' AND " + types + "_slug='" + slug + "' " +
		"ORDER BY " + order

	// Exec
	con := database.CreateCon()
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	// Map
	for rows.Next() {
		err = rows.Scan(
			&obj.Id,
			&obj.SourceTitle,
			&obj.SourceUrl,
		)

		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg(baseTable, len(arrobj))
	if len(arrobj) == 0 {
		res.Data = nil
	} else {
		res.Data = arrobj
	}

	return res, nil
}

func GetAnimalCountryBySlug(slug string) (response.Response, error) {
	// Declaration
	var obj models.GetAnimalCountries
	var arrobj []models.GetAnimalCountries
	var res response.Response
	var baseTable = "animals_location"
	var sqlStatement string

	// Converted column
	var Total sql.NullInt32

	// Query builder
	joinTemplate1 := builders.GetTemplateJoin("total", baseTable, "animals_id", "animals", "id", false)
	joinTemplate2 := builders.GetTemplateJoin("total", baseTable, "countries_code", "countries", "countries_code", false)

	sqlStatement = "SELECT " + baseTable + ".id, LOWER(countries.countries_code), countries_name, total " +
		"FROM " + baseTable + " " +
		joinTemplate1 + " " + joinTemplate2 + " " +
		"WHERE animals_slug='" + slug + "' " +
		"ORDER BY total DESC"

	// Exec
	con := database.CreateCon()
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	// Map
	for rows.Next() {
		err = rows.Scan(
			&obj.Id,
			&obj.CountryCode,
			&obj.CountryName,
			&Total,
		)

		if err != nil {
			return res, err
		}

		// Converted
		var intTotal int
		if Total.Valid {
			intTotal = int(Total.Int32)
		}
		obj.Total = intTotal

		arrobj = append(arrobj, obj)
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg(baseTable, len(arrobj))
	if len(arrobj) == 0 {
		res.Data = nil
	} else {
		res.Data = arrobj
	}

	return res, nil
}

func GetAllCountries() (response.Response, error) {
	// Declaration
	var obj models.GetCountries
	var arrobj []models.GetCountries
	var res response.Response
	var baseTable = "countries"
	var sqlStatement string

	sqlStatement = "SELECT countries_code, countries_name " +
		"FROM " + baseTable + " " +
		"ORDER BY countries_name ASC"

	// Exec
	con := database.CreateCon()
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	// Map
	for rows.Next() {
		err = rows.Scan(
			&obj.CountryCode,
			&obj.CountryName,
		)

		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg(baseTable, len(arrobj))
	if len(arrobj) == 0 {
		res.Data = nil
	} else {
		res.Data = arrobj
	}

	return res, nil
}
