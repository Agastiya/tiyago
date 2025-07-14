package utils

import "strings"

func ValidateSortColumn(allowedFields map[string]string, sortColumn string, defaultSortColumn string) string {
	column := allowedFields[sortColumn]
	if column == "" {
		column = defaultSortColumn
	}
	return column
}

func ValidateSortOrder(sortOrder string, defaultSortOrder string) string {
	var orders = []string{"ASC", "DESC"}
	for _, ordering := range orders {
		if strings.EqualFold(sortOrder, ordering) {
			return sortOrder
		}
	}
	return defaultSortOrder
}
