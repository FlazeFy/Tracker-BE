package converter

import (
	"database/sql"
	"reflect"
	"strconv"
	"strings"
)

func CheckNullString(data sql.NullString) string {
	var res string
	if data.Valid {
		res = data.String
	} else {
		res = ""
	}

	return res
}

func TotalChar(val string) int {
	trimed := strings.TrimSpace(val)
	return len(trimed)
}

func StructToMap(data interface{}) (map[string]interface{}, error) {
	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	out := make(map[string]interface{})
	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		fieldValue := val.Field(i).Interface()
		fieldName := typ.Field(i).Tag.Get("json")

		// Custom conversion logic
		switch fieldValue.(type) {
		case string:
			// Example of converting string to int
			if intValue, err := strconv.Atoi(fieldValue.(string)); err == nil {
				out[fieldName] = intValue
			} else {
				out[fieldName] = fieldValue
			}
		default:
			out[fieldName] = fieldValue
		}
	}

	return out, nil
}
