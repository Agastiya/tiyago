package response

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/agastiya/tiyago/pkg/constant"
)

func ResponseSuccess(w http.ResponseWriter, body any, httpInternalCode constant.HttpInternalCode) {
	result := ResponseSuccessStruct{
		RestCode:    int(httpInternalCode),
		RestStatus:  httpInternalCode.Response().HttpTitle,
		RestMessage: httpInternalCode.Response().Description,
		RestResult:  body,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpInternalCode.Response().HttpCode)
	json.NewEncoder(w).Encode(result)
}

func ResponseError(w http.ResponseWriter, err error, httpInternalCode constant.HttpInternalCode) {
	result := ResponseErrorStruct{
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

func ResponseService(HasErr bool, Err error, InternalCode constant.HttpInternalCode, tx *sql.Tx, Result any) (result RespResultService) {
	result = RespResultService{
		HasErr:       HasErr,
		Err:          Err,
		InternalCode: InternalCode,
		Tx:           tx,
		Result:       Result,
	}
	return
}
