package constant

import "net/http"

type (
	HttpInternalCode int

	Response struct {
		HttpCode    int
		HttpTitle   string
		Description string
	}
)

const (

	// 2xx
	StatusOKJson HttpInternalCode = 2000

	// 4xx
	StatusDataBadRequest  HttpInternalCode = 4000
	StatusUnauthorized    HttpInternalCode = 4010
	StatusForbidden       HttpInternalCode = 4030
	StatusDataNotFound    HttpInternalCode = 4040
	StatusDataConflict    HttpInternalCode = 4090
	StatusTooManyRequests HttpInternalCode = 4290

	// 5xx
	StatusInternalServerError HttpInternalCode = 5000
)

func (i HttpInternalCode) is2xx() bool {
	return i >= 2000 && i < 3000
}

func (i HttpInternalCode) is4xx() bool {
	return i >= 4000 && i < 5000
}

func (i HttpInternalCode) is5xx() bool {
	return i >= 5000 && i < 6000
}

func (i HttpInternalCode) Response() Response {

	if i.is2xx() {
		return response2xx[i]
	} else if i.is4xx() {
		return response4xx[i]
	} else if i.is5xx() {
		return response5xx[i]
	} else {
		return Response{}
	}
}

var response2xx = map[HttpInternalCode]Response{
	StatusOKJson: {
		HttpCode:    http.StatusOK,
		HttpTitle:   http.StatusText(http.StatusOK),
		Description: "Success",
	},
}

var response4xx = map[HttpInternalCode]Response{
	StatusDataBadRequest: {
		HttpCode:    http.StatusBadRequest,
		HttpTitle:   http.StatusText(http.StatusBadRequest),
		Description: "Invalid Data",
	},
	StatusUnauthorized: {
		HttpCode:    http.StatusUnauthorized,
		HttpTitle:   http.StatusText(http.StatusUnauthorized),
		Description: "User Unauthorized",
	},
	StatusForbidden: {
		HttpCode:    http.StatusForbidden,
		HttpTitle:   http.StatusText(http.StatusForbidden),
		Description: "Forbidden",
	},
	StatusDataNotFound: {
		HttpCode:    http.StatusNotFound,
		HttpTitle:   http.StatusText(http.StatusNotFound),
		Description: "Data Not Found",
	},
	StatusDataConflict: {
		HttpCode:    http.StatusConflict,
		HttpTitle:   http.StatusText(http.StatusConflict),
		Description: "Data Conflict",
	},
	StatusTooManyRequests: {
		HttpCode:    http.StatusTooManyRequests,
		HttpTitle:   http.StatusText(http.StatusTooManyRequests),
		Description: "To Many Request",
	},
}

var response5xx = map[HttpInternalCode]Response{
	StatusInternalServerError: {
		HttpCode:    http.StatusInternalServerError,
		HttpTitle:   http.StatusText(http.StatusInternalServerError),
		Description: "internal server error",
	},
}
