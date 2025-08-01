package response

import (
	"database/sql"
)

type APIResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result any    `json:"result,omitempty"`
	Error  any    `json:"error,omitempty"`
}

type ServiceResult struct {
	HasErr   bool
	Err      error
	HttpCode int
	Tx       *sql.Tx
	Result   any
}
