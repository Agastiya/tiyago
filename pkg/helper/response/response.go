package response

import (
	"compress/gzip"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/agastiya/tiyago/pkg/constant"
)

func ResponseSuccess(w http.ResponseWriter, body any, httpInternalCode constant.HttpInternalCode) {
	result := SuccessResponse{
		RestCode:    int(httpInternalCode),
		RestStatus:  httpInternalCode.Response().HttpTitle,
		RestMessage: httpInternalCode.Response().Description,
		RestResult:  body,
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Encoding", "gzip")
	w.WriteHeader(httpInternalCode.Response().HttpCode)

	gz := gzip.NewWriter(w)
	defer gz.Close()

	if err := json.NewEncoder(gz).Encode(result); err != nil {
		fmt.Println("Failed to encode JSON:", err)
	}
}

func ResponseError(w http.ResponseWriter, err error, httpInternalCode constant.HttpInternalCode) {
	result := ErrorResponse{
		RestCode:    int(httpInternalCode),
		RestStatus:  httpInternalCode.Response().HttpTitle,
		RestMessage: httpInternalCode.Response().Description,
	}

	if err != nil {
		result.RestResult = append(result.RestResult, err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpInternalCode.Response().HttpCode)
	json.NewEncoder(w).Encode(result)
}

func ResponseService(hasErr bool, err error, internalCode constant.HttpInternalCode, tx *sql.Tx, result any) RespResultService {
	return RespResultService{HasErr: hasErr, Err: err, InternalCode: internalCode, Tx: tx, Result: result}
}
