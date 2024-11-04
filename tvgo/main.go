package main

import "github.com/Mirsadikovv/tvgo/cmd"

// @title Tournament service API
// @version 1.0
// @description This is a Tournament service API.
// @BasePath /api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cmd.Exec()
}
