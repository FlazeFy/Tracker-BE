package repositories

import (
	"net/http"
	"time"
	"tracker/modules/locations/models"
	"tracker/packages/database"
	"tracker/packages/helpers/converter"
	"tracker/packages/helpers/generator"
	"tracker/packages/helpers/response"
)

func PostLocation(d models.Location) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "locations"
	var sqlStatement string
	id, _ := generator.GenerateUUID(32)
	dt := time.Now().Format("2006-01-02 15:04:05")

	// Command builder
	sqlStatement = "INSERT INTO " + baseTable + " (id, location_name, location_desc, location_lat, location_long, location_category, location_address, location_apps, created_at,created_by) " +
		"VALUES (?,?,?,?,?,?,?,?,?,?)"

	// Exec - MySQL
	con := database.CreateCon()
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	result, err := stmt.Exec(id, d.LocName, d.LocDesc, d.LocLat, d.LocLong, d.LocCat, d.LocAddr, d.LocApps, dt, "1")
	if err != nil {
		return res, err
	}

	// Response Builder
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	// Exec - Firebase
	dataMap, err := converter.StructToMap(d)
	firebaseInsert := database.InsertFirebase(id, baseTable, dataMap)

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateCommandMsg(baseTable, "create", int(rowsAffected))
	res.Data = map[string]interface{}{
		"id":             id,
		"data":           d,
		"rows_affected":  rowsAffected,
		"is_realtime_db": firebaseInsert,
	}

	return res, nil
}
