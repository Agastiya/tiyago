package utils

import (
	"fmt"

	"github.com/agastiya/tiyago/dto"
)

func SetDefaultParams(params dto.Pagination, defaultParams dto.Pagination) dto.Pagination {
	if params.SortColumn == "" {
		params.SortColumn = defaultParams.SortColumn
	}
	if params.SortOrder == "" {
		params.SortOrder = defaultParams.SortOrder
	}
	if params.PageNumber == 0 {
		params.PageNumber = defaultParams.PageNumber
	}
	if params.PageSize == 0 {
		params.PageSize = defaultParams.PageSize
	}
	return params
}

func CheckExistsFieldName(fieldName, value string, excludeId int64, checkFn func(string, int64) (bool, error)) error {
	exists, err := checkFn(value, excludeId)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("%s already registered", fieldName)
	}
	return nil
}
