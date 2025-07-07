package controller

import (
	"net/http"

	"github.com/agastiya/tiyago/pkg/constant"
	"github.com/agastiya/tiyago/pkg/helper/response"
)

type Base interface {
	Ping(w http.ResponseWriter, r *http.Request)
}

// @Tags     Ping
// @Accept   json
// @Produce  json
// @Router   /ping [get]
func (c Controller) Ping(w http.ResponseWriter, r *http.Request) {
	response.ResponseSuccess(w, nil, constant.StatusOKJson)
}
