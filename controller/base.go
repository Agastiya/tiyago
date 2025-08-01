package controller

import (
	"net/http"

	"github.com/agastiya/tiyago/pkg/helper/response"
)

type BaseController struct{}

func NewBaseController() IBaseController {
	return &BaseController{}
}

type IBaseController interface {
	Ping(w http.ResponseWriter, r *http.Request)
}

// @Tags     Ping
// @Accept   json
// @Produce  json
// @Router   /ping [get]
func (bc *BaseController) Ping(w http.ResponseWriter, r *http.Request) {
	response.JSONResponse(w, nil, nil, http.StatusOK)
}
