package main

import (
	"github.com/agastiya/tiyago/app"
)

// @title Tiyago REST API
// @version 1.0
// @description Tiyago RESTful API
// @basePath /tiyago
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Bearer token authentication. Example: Bearer abcdefghijklmnopqrstuvwxyz1234567890
func main() {
	app.AppInit()
}
