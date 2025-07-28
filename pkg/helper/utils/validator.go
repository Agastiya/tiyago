package utils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func Validate(input any) error {
	err := validate.Struct(input)
	if err == nil {
		return nil
	}

	var sb strings.Builder
	if _, ok := err.(*validator.InvalidValidationError); ok {
		return fmt.Errorf("invalid validation input: %w", err)
	}

	validationErrors := err.(validator.ValidationErrors)

	for i, fieldErr := range validationErrors {
		if fieldErr.Param() != "" {
			sb.WriteString(fmt.Sprintf("%s %s:%s", fieldErr.Field(), fieldErr.Tag(), fieldErr.Param()))
		} else {
			sb.WriteString(fmt.Sprintf("%s %s", fieldErr.Field(), fieldErr.Tag()))
		}
		if i < len(validationErrors)-1 {
			sb.WriteString(", ")
		}
	}

	return errors.New(sb.String())
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
