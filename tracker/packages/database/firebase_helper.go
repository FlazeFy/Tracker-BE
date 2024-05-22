package database

import (
	"fmt"
)

type GroupedResult struct {
	Context string
	Total   int
}

func GroupAndCount[T any, U comparable](items []T, getField func(T) U, limit int) ([]GroupedResult, error) {
	grouped := make(map[U]int)
	distinctContexts := 0

	for _, item := range items {
		fieldValue := getField(item)
		grouped[fieldValue]++
		if grouped[fieldValue] == 1 {
			distinctContexts++
		}
	}

	var results []GroupedResult
	for context, total := range grouped {
		if distinctContexts <= limit {
			results = append(results, GroupedResult{
				Context: fmt.Sprintf("%v", context),
				Total:   total,
			})
		}
	}

	return results, nil
}
