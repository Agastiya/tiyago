package main

import (
	"github.com/agastiya/tiyago/app"
)

// @title Tiyago
// @version 1.0
// @description Tiyago Restfull API
// @BasePath /tiyago
// @securityDefinitions.apikey Bearer
// @in Header
// @name Authorization
// @description Example: Bearer abcdefghijklmnopqrstuvwxyz1234567890
func main() {
	app.AppInit()
}
