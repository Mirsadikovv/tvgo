package main

import "github.com/Mirsadikovv/tvgo/cmd"

// @title TVGO API Server
// @version 1.0
// @description This is a TVGO API server.
// @BasePath /api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cmd.Exec()
}
