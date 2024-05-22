package repositories

import (
	"context"
	"database/sql"
	"log"
	"math"
	"net/http"
	"tracker/modules/locations/models"
	"tracker/packages/builders"
	"tracker/packages/database"
	"tracker/packages/helpers/converter"
	"tracker/packages/helpers/generator"
	"tracker/packages/helpers/response"
	"tracker/packages/utils/pagination"
)

func GetAllLocations(page, pageSize int, path string) (response.Response, error) {
	// Declaration
	var obj models.Location
	var arrobj []models.Location
	var res response.Response
	var baseTable = "locations"
	var sqlStatement string

	// Converted Column
	var LocDesc sql.NullString
	var LocAddr sql.NullString
	var LocApps sql.NullString

	sqlStatement = "SELECT location_name, location_desc, location_lat, location_long, location_category, location_address, location_apps " +
		"FROM " + baseTable + " " +
		"ORDER BY location_name DESC " +
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
			&obj.LocName,
			&LocDesc,
			&obj.LocLat,
			&obj.LocLong,
			&obj.LocCat,
			&LocAddr,
			&LocApps,
		)

		if err != nil {
			return res, err
		}

		// Nullable check
		obj.LocDesc = converter.CheckNullString(LocDesc)
		obj.LocAddr = converter.CheckNullString(LocAddr)
		obj.LocApps = converter.CheckNullString(LocApps)

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

func GetAllLocationsFirebase() (response.Response, error) {
	// Declaration
	var obj map[string]models.Location
	var res response.Response
	var baseTable = "locations"

	// Exec
	ctx := context.Background()
	client, err := database.InitializeFirebaseDB(ctx)
	if err != nil {
		log.Fatalln("error in initializing firebase DB client: ", err)
		return res, err
	}

	ref := client.NewRef(baseTable)
	if err := ref.Get(ctx, &obj); err != nil {
		log.Fatalln("error in reading from firebase DB:", err)
		return res, err
	}

	var arrobj []models.Location
	for _, v := range obj {
		arrobj = append(arrobj, v)
	}

	// Build response
	total := len(arrobj)
	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg(baseTable, total)
	if total == 0 {
		res.Data = nil
	} else {
		res.Data = arrobj
	}

	return res, nil
}
