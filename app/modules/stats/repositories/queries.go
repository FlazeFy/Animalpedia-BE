package repositories

import (
	"app/modules/stats/models"
	"app/packages/builders"
	"app/packages/database"
	"app/packages/helpers/generator"
	"app/packages/helpers/response"
	"fmt"
	"net/http"
	"strconv"
)

func GetTotalStats(ord string, mainCol string, baseTable string, joinArgs *string) (response.Response, error) {
	// Declaration
	var obj models.GetMostAppear
	var arrobj []models.GetMostAppear
	var res response.Response
	var sqlStatement string

	// Converted column
	var totalStr string

	// Query builder
	sqlStatement = builders.GetTemplateStats(mainCol, baseTable, "most_appear", ord, joinArgs)

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
			&obj.Context,
			&totalStr)

		if err != nil {
			return res, err
		}

		// Converted
		totalInt, err := strconv.Atoi(totalStr)
		if err != nil {
			return res, err
		}

		obj.Total = totalInt
		arrobj = append(arrobj, obj)
	}

	// Total
	total, err := builders.GetTotalCount(con, baseTable, nil)
	if err != nil {
		return res, err
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg("Stats", total)
	if total == 0 {
		res.Data = nil
	} else {
		res.Data = arrobj
	}

	return res, nil
}

func GetTotalAnimalPopulationByCountry(ord string) (response.Response, error) {
	// Declaration
	var obj models.GetMostAppear
	var arrobj []models.GetMostAppear
	var res response.Response
	var sqlStatement string

	// Converted column
	var totalStr string

	// Query builder
	sqlStatement = "SELECT countries_name as context, SUM(total) AS total " +
		"FROM animals_location al " +
		"JOIN countries c ON c.countries_code = al.countries_code " +
		"JOIN animals a ON a.id = al.animals_id " +
		"GROUP BY context ORDER BY total " + ord + " LIMIT 8"

	fmt.Println(sqlStatement)

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
			&obj.Context,
			&totalStr)

		if err != nil {
			return res, err
		}

		// Converted
		totalInt, err := strconv.Atoi(totalStr)
		if err != nil {
			return res, err
		}

		obj.Total = totalInt
		arrobj = append(arrobj, obj)
	}

	// Total
	total, err := builders.GetTotalCount(con, "animals_location", nil)
	if err != nil {
		return res, err
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg("Stats", total)
	if total == 0 {
		res.Data = nil
	} else {
		res.Data = arrobj
	}

	return res, nil
}
