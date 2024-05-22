package models

type (
	PostLocation struct {
		LocName string `json:"location_name"`
		LocDesc string `json:"location_desc"`
		LocLat  string `json:"location_lat"`
		LocLong string `json:"location_long"`
		LocCat  string `json:"location_category"`
		LocAddr string `json:"location_address"`
		LocApps string `json:"location_apps"`
	}
)
