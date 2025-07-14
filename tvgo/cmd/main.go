package cmd

import (
	"log"
	"os"
	"strconv"

	_ "github.com/Mirsadikovv/tvgo/docs"
	echoSwagger "github.com/swaggo/echo-swagger"

	userHandler "github.com/Mirsadikovv/tvgo/user/handler"
	userModel "github.com/Mirsadikovv/tvgo/user/model"
	userService "github.com/Mirsadikovv/tvgo/user/service"
	"gorm.io/gorm"

	playlistHandler "github.com/Mirsadikovv/tvgo/playlist/handler"
	playlistModel "github.com/Mirsadikovv/tvgo/playlist/model"
	playlistService "github.com/Mirsadikovv/tvgo/playlist/service"

	genadiHandler "github.com/Mirsadikovv/tvgo/genadi/handler"

	loader "github.com/fobus1289/ufa_shared/config"
	"github.com/fobus1289/ufa_shared/pg"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Exec() {

	projectEnv := loader.ProjectEnv()
	log.Println(os.Getenv("POSTGRES_PASSWORD"), "-----------------------")
	pgConfig := pg.NewConfigEmpty()
	{
		pgConfig.SetHost(projectEnv.PgHost).
			SetPort(projectEnv.PgPort).
			SetDbname(projectEnv.PgDB).
			SetUser(projectEnv.PgUser).
			SetPassword(os.Getenv("POSTGRES_PASSWORD"))
	}

	db, err := pg.NewGorm(&pgConfig)
	{
		if err != nil {
			log.Panicln(err)
		}

		db.AutoMigrate(
			userModel.UserModel{},
			playlistModel.PlaylistModel{},
		)
	}

	router := echo.New()
	{
		setMiddlewares(router)
		createHandler(router, db)
		runHTTPServerOnAddr(router, projectEnv.HttpPort)
	}

}

func runHTTPServerOnAddr(handler *echo.Echo, port int) {
	url := strconv.FormatInt(int64(port), 10)
	{
		log.Panicln(handler.Start(":" + url))
	}
}

func setMiddlewares(router *echo.Echo) {
	router.Use(middleware.RemoveTrailingSlash())
	router.Use(middleware.RequestID())
	// router.Use(middleware.Recover())
	router.Use(middleware.CORS())
}

func createHandler(router *echo.Echo, db *gorm.DB) {
	group := router.Group("/api/v1")
	{
		group.GET("/swagger/*", echoSwagger.WrapHandler)

		userHandler.NewHandler(group, userService.NewService(db))
		playlistHandler.NewHandler(group, playlistService.NewService(db))
		genadiHandler.NewHandler(group)
	}
}
