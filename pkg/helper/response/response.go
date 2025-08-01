package response

import (
	"compress/gzip"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/agastiya/tiyago/pkg/constant"
)

func JSONResponse(w http.ResponseWriter, body any, err error, httpInternalCode constant.HttpInternalCode) {
	res := APIResponse{
		Code:    int(httpInternalCode),
		Status:  httpInternalCode.Response().HttpTitle,
		Message: httpInternalCode.Response().Description,
	}

	if err != nil {
		res.Error = err.Error()
	} else {
		res.Result = body
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Encoding", "gzip")
	w.WriteHeader(httpInternalCode.Response().HttpCode)

	gz := gzip.NewWriter(w)
	defer gz.Close()

	if encodeErr := json.NewEncoder(gz).Encode(res); encodeErr != nil {
		fmt.Println("Failed to encode JSON:", encodeErr)
	}
}

func NewServiceResult(hasErr bool, err error, internalCode constant.HttpInternalCode, tx *sql.Tx, result any) ServiceResult {
	return ServiceResult{hasErr, err, internalCode, tx, result}
}
