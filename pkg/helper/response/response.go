package response

import (
	"compress/gzip"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func JSONResponse(w http.ResponseWriter, body any, err error, httpCode int) {
	res := APIResponse{
		Code:   httpCode,
		Status: http.StatusText(httpCode),
	}

	if err != nil {
		res.Error = err.Error()
	} else {
		res.Result = body
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Encoding", "gzip")
	w.WriteHeader(httpCode)

	gz := gzip.NewWriter(w)
	defer gz.Close()

	if encodeErr := json.NewEncoder(gz).Encode(res); encodeErr != nil {
		fmt.Println("Failed to encode JSON:", encodeErr)
	}
}

func NewServiceResult(hasErr bool, err error, httpCode int, tx *sql.Tx, result any) ServiceResult {
	return ServiceResult{hasErr, err, httpCode, tx, result}
}
