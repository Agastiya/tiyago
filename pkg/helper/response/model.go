package response

import (
	"database/sql"

	"github.com/agastiya/tiyago/pkg/constant"
)

type SuccessResponse struct {
	RestCode    int    `json:"code"`
	RestStatus  string `json:"status"`
	RestMessage string `json:"message"`
	RestResult  any    `json:"result" swaggertype:"object,string" example:"key:value,key2:value2"`
}

type ErrorResponse struct {
	RestCode    int    `json:"code"`
	RestStatus  string `json:"status"`
	RestMessage string `json:"message"`
	RestResult  []any  `json:"result" swaggertype:"object,string" example:"key:value,key2:value2"`
}

type RespResultService struct {
	HasErr       bool
	Err          error
	InternalCode constant.HttpInternalCode
	Tx           *sql.Tx
	Result       any
}
