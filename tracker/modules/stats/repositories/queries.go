package repositories

import (
	"context"
	"log"
	"net/http"
	"tracker/modules/locations/models"
	"tracker/packages/database"
	"tracker/packages/helpers/generator"
	"tracker/packages/helpers/response"
)

func GetTotalLocationByCategoryFirebase() (response.Response, error) {
	// Declaration
	var obj map[string]models.Location
	var res response.Response
	var baseTable = "locations"
	limit := 2

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

	// Set grouping & count
	getCategory := func(l models.Location) string {
		return l.LocCat
	}

	results, err := database.GroupAndCount(arrobj, getCategory, limit)
	if err != nil {
		log.Fatalln("error in grouping and counting:", err)
		return res, err
	}

	// Build response
	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg(baseTable, len(arrobj))
	res.Data = results

	return res, nil
}

func GetTotalLocationByAppsFirebase() (response.Response, error) {
	// Declaration
	var obj map[string]models.Location
	var res response.Response
	var baseTable = "locations"
	limit := 2

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

	// Set grouping & count
	getCategory := func(l models.Location) string {
		return l.LocApps
	}

	results, err := database.GroupAndCount(arrobj, getCategory, limit)
	if err != nil {
		log.Fatalln("error in grouping and counting:", err)
		return res, err
	}

	// Build response
	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg(baseTable, len(arrobj))
	res.Data = results

	return res, nil
}
