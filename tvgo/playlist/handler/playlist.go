package handler

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/Mirsadikovv/tvgo/playlist/model"
	_ "github.com/Mirsadikovv/tvgo/utils"
	_ "github.com/fobus1289/ufa_shared/http/response"

	"github.com/Mirsadikovv/tvgo/playlist/dto"
	"github.com/Mirsadikovv/tvgo/playlist/service"
	"github.com/fobus1289/ufa_shared/http"
	"github.com/fobus1289/ufa_shared/http/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type playlistHandler struct {
	service service.PlaylistService
}

func NewHandler(router *echo.Group, service service.PlaylistService) {

	group := router.Group("/playlist")
	{
		handler := &playlistHandler{service: service}

		group.POST("", handler.Create)
		group.GET("/page", handler.Page)
		group.GET("/:id", handler.GetById)
		group.GET("/search", handler.Search)
		group.PATCH("/:id", handler.Update)
		group.PUT("/:id", handler.ChangeVisibility)
		group.DELETE("/:id", handler.Delete)

	}
}

// Create godoc
// @Summary      Create a new playlist
// @Description  Create playlist
// @Tags 		 playlist
// @ID           create-playlist
// @Accept       json
// @Produce      json
// @Param        input body dto.CreatePlaylistDto true "playlist information"
// @Success      201 {object} utils.ID "Successful operation"
// @Failure      400 {object} utils.ErrorResponse "Bad request"
// @Failure      500 {object} utils.ErrorResponse "Internal server error"
// @Router       /playlist [post]
func (e *playlistHandler) Create(c echo.Context) error {
	var createDto dto.CreatePlaylistDto
	{
		if err := c.Bind(&createDto); err != nil {
			return http.HTTPError(err).BadRequest()
		}

		if err := validator.Validate(createDto); err != nil {
			return http.HTTPError(err).BadRequest()
		}
	}

	idDto, err := e.service.Create(&createDto)
	{
		if err != nil {
			return http.HTTPError(err).BadRequest()
		}
	}

	return http.Response(c).Created(idDto)
}

// Page godoc
// @Summary      GetContent all playlist with pagination
// @Description  GetContent all playlist with pagination
// @Tags 		 playlist
// @ID           get-all-playlist
// @Accept       json
// @Produce      json
// @Param        page query string false "Page number" default(1)
// @Param        perpage query string false "Number of items per page" default(10)
// @Param        search query string false "Searching by name or description"
// @Success      200 {object} utils.ID "Successful operation"
// @Failure      400 {object} utils.ErrorResponse "Bad request"
// @Failure      500 {object} utils.ErrorResponse "Internal server error"
// @Router       /playlist/page [get]
func (e *playlistHandler) Page(c echo.Context) error {
	var (
		search   = c.QueryParam("search")
		page     = c.QueryParam("page")
		perPage  = c.QueryParam("perpage")
		paginate = http.NewPaginate(page, perPage)
		ctx      = c.Request().Context()
	)

	limitFilter := func(tx *gorm.DB) *gorm.DB {
		return tx.Offset(paginate.Skip()).Limit(paginate.Take()).Order("id ASC")
	}

	filter := func(tx *gorm.DB) *gorm.DB {
		tx = tx.Where("is_visible = ?", true)

		if search != "" {
			search = fmt.Sprintf("%%%s%%", search)
			tx = tx.Where("name ILIKE ?", search)
		}
		return tx
	}

	pageData, err := e.service.Page(ctx, paginate.Take(), filter, limitFilter)
	{
		if err != nil {
			return http.HTTPError(err).BadRequest()
		}
	}
	return http.Response(c).OK(pageData)
}

// Search godoc
// @Summary      GetContent all playlist with pagination
// @Description  GetContent all playlist with pagination
// @Tags 		 playlist
// @ID           search-all-playlist
// @Accept       json
// @Produce      json
// @Param        search query string false "Searching by name or description"
// @Param        limit  query int    false "Limit the number of results" default(20)
// @Success      200 {object} utils.ID "Successful operation"
// @Failure      400 {object} utils.ErrorResponse "Bad request"
// @Failure      500 {object} utils.ErrorResponse "Internal server error"
// @Router       /playlist/search [get]
func (e *playlistHandler) Search(c echo.Context) error {
	const defaultLimit = 15
	const maxLimit = 100

	search := strings.TrimSpace(c.QueryParam("search"))
	ctx := c.Request().Context()

	limitParam := c.QueryParam("limit")
	limit, err := strconv.Atoi(limitParam)
	{
		if err != nil || limit <= 0 {
			limit = defaultLimit
		}
		if limit > maxLimit {
			limit = maxLimit
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {
		tx = tx.Where("is_visible = ?", true)

		if search != "" {
			searchTerm := fmt.Sprintf("%%%s%%", search)
			tx = tx.Where("name ILIKE ?", searchTerm)
		}
		return tx.Limit(limit)
	}

	searchData, err := e.service.Find(ctx, filter)
	if err != nil {
		return http.HTTPError(err).BadRequest()
	}

	return http.Response(c).OK(searchData)
}

// GetById godoc
// @Summary      GetContent playlist by ID
// @Description  GetContent playlist by ID
// @Tags 		 playlist
// @ID           get-playlist-by-id
// @Accept       json
// @Produce      json
// @Param        id path string true "playlist ID"
// @Success      200 {object} model.PlaylistModel "Successful operation"
// @Failure      400 {object} utils.ErrorResponse "Bad request"
// @Failure      500 {object} utils.ErrorResponse "Internal server error"
// @Router       /playlist/{id} [get]
func (e *playlistHandler) GetById(c echo.Context) error {

	var id int64
	{
		if !http.PathValue(c.Param("id")).TryInt64(&id) {
			err := errors.New("error parse id")
			return http.HTTPError(err).BadRequest()
		}
	}

	ctx := c.Request().Context()

	filter := func(tx *gorm.DB) *gorm.DB {
		return tx.Where("id = ?", id).Where("is_visible = ?", true)
	}

	playlist, err := e.service.FindOne(ctx, filter)
	{
		if err != nil {
			return http.HTTPError(err).BadRequest()
		}
	}

	return http.Response(c).OK(playlist)
}

// Update godoc
// @Summary      Update playlist information
// @Description  Update playlist information by ID
// @Tags 		 playlist
// @ID           update-playlist
// @Accept       json
// @Param        id path string true "playlist ID"
// @Param        input body dto.UpdatePlaylistDto true "playlist information"
// @Success      204 "Successful operation"
// @Failure      400 {object} utils.ErrorResponse "Bad request"
// @Failure      500 {object} utils.ErrorResponse "Internal server error"
// @Router       /playlist/{id} [patch]
func (e *playlistHandler) Update(c echo.Context) error {
	var id int64
	{
		if !http.PathValue(c.Param("id")).TryInt64(&id) {
			err := errors.New("parse id error")
			return http.HTTPError(err).BadRequest()
		}
	}

	var updateDto dto.UpdatePlaylistDto
	{
		if err := c.Bind(&updateDto); err != nil {
			return http.HTTPError(err).BadRequest()
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {
		return tx.Where("id", id)
	}

	err := e.service.Update(&updateDto, filter)
	{
		if err != nil {
			return http.HTTPError(err).BadRequest()
		}
	}

	return http.Response(c).NoContent()
}

// Delete godoc
// @Summary      Delete playlist
// @Description  Delete playlist by ID
// @Tags 		 playlist
// @ID           delete-playlist
// @Accept       json
// @Param        id path string true "playlist ID"
// @Success      204 "Successful operation"
// @Failure      400 {object} utils.ErrorResponse "Bad request"
// @Failure      500 {object} utils.ErrorResponse "Internal server error"
// @Router       /playlist/{id} [delete]
func (e *playlistHandler) Delete(c echo.Context) error {
	var id int64
	{
		if !http.PathValue(c.Param("id")).TryInt64(&id) {
			err := errors.New("parse id error")
			return http.HTTPError(err).BadRequest()
		}
	}

	{
		filter := func(tx *gorm.DB) *gorm.DB {
			return tx.Where("id", id)
		}

		if err := e.service.Delete(filter); err != nil {
			return http.HTTPError(err).BadRequest()
		}
	}

	return http.Response(c).NoContent()
}

// PUT godoc
// @Summary      put playlist
// @Description  put playlist by ID
// @Tags 		 playlist
// @ID           change-visibility-playlist
// @Accept       json
// @Param        id path string true "playlist ID"
// @Success      204 "Successful operation"
// @Failure      400 {object} utils.ErrorResponse "Bad request"
// @Failure      500 {object} utils.ErrorResponse "Internal server error"
// @Router       /playlist/{id} [put]
func (e *playlistHandler) ChangeVisibility(c echo.Context) error {
	var id int64
	{
		if !http.PathValue(c.Param("id")).TryInt64(&id) {
			err := errors.New("parse id error")
			return http.HTTPError(err).BadRequest()
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {
		return tx.Where("id", id)
	}

	err := e.service.ChangeVisibility(filter)
	{
		if err != nil {
			return http.HTTPError(err).BadRequest()
		}
	}

	return http.Response(c).NoContent()
}
