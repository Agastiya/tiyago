package utils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func Validate(mystruct any) error {
	validate = validator.New()
	err := validate.Struct(mystruct)
	if err != nil {
		var errorValidate = "error: "
		if _, ok := err.(*validator.InvalidValidationError); ok {
			errorValidate += fmt.Sprintf("%s ,", err.Error())
		}
		errValidator := err.(validator.ValidationErrors)
		for index, err := range errValidator {
			if index+1 == len(errValidator) {
				errorValidate += err.Field() + " " + err.Tag()
			} else {
				errorValidate += err.Field() + " " + err.Tag() + ","
			}
		}
		return errors.New(errorValidate)
	}
	return nil
}

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
