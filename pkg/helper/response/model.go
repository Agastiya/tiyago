package response

import (
	"database/sql"

	"github.com/agastiya/tiyago/pkg/constant"
)

type APIResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  any    `json:"result,omitempty"`
	Error   any    `json:"error,omitempty"`
}

type ServiceResult struct {
	HasErr       bool
	Err          error
	InternalCode constant.HttpInternalCode
	Tx           *sql.Tx
	Result       any
}
