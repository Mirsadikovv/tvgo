package service

import (
	playlistHandler "github.com/Mirsadikovv/tvgo/playlist/handler"
	playlistService "github.com/Mirsadikovv/tvgo/playlist/service"
	userHandler "github.com/Mirsadikovv/tvgo/user/handler"
	userService "github.com/Mirsadikovv/tvgo/user/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewService(routerGroup *echo.Group, db *gorm.DB) {
	userHandler.NewHandler(routerGroup, userService.NewService(db))
	playlistHandler.
		NewHandler(routerGroup, playlistService.NewService(db))

}
