package handler

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/Mirsadikovv/tvgo/user/model"
	_ "github.com/Mirsadikovv/tvgo/utils"
	_ "github.com/fobus1289/ufa_shared/http/response"

	"github.com/Mirsadikovv/tvgo/user/dto"
	"github.com/Mirsadikovv/tvgo/user/service"
	"github.com/fobus1289/ufa_shared/http"
	"github.com/fobus1289/ufa_shared/http/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type userHandler struct {
	service service.UserService
}

func NewHandler(router *echo.Group, service service.UserService) {

	group := router.Group("/user")
	{
		handler := &userHandler{service: service}

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
// @Summary      Create a new user
// @Description  Create user
// @Tags 		 user
// @ID           create-user
// @Accept       json
// @Produce      json
// @Param        input body dto.CreateUserDto true "user information"
// @Success      201 {object} utils.ID "Successful operation"
// @Failure      400 {object} utils.ErrorResponse "Bad request"
// @Failure      500 {object} utils.ErrorResponse "Internal server error"
// @Router       /user [post]
func (e *userHandler) Create(c echo.Context) error {
	var createDto dto.CreateUserDto
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
// @Summary      GetContent all user with pagination
// @Description  GetContent all user with pagination
// @Tags 		 user
// @ID           get-all-user
// @Accept       json
// @Produce      json
// @Param        page query string false "Page number" default(1)
// @Param        perpage query string false "Number of items per page" default(10)
// @Param        search query string false "Searching by name or description"
// @Success      200 {object} utils.ID "Successful operation"
// @Failure      400 {object} utils.ErrorResponse "Bad request"
// @Failure      500 {object} utils.ErrorResponse "Internal server error"
// @Router       /user/page [get]
func (e *userHandler) Page(c echo.Context) error {
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
// @Summary      GetContent all user with pagination
// @Description  GetContent all user with pagination
// @Tags 		 user
// @ID           search-all-user
// @Accept       json
// @Produce      json
// @Param        search query string false "Searching by name or description"
// @Param        limit  query int    false "Limit the number of results" default(20)
// @Success      200 {object} utils.ID "Successful operation"
// @Failure      400 {object} utils.ErrorResponse "Bad request"
// @Failure      500 {object} utils.ErrorResponse "Internal server error"
// @Router       /user/search [get]
func (e *userHandler) Search(c echo.Context) error {
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
// @Summary      GetContent user by ID
// @Description  GetContent user by ID
// @Tags 		 user
// @ID           get-user-by-id
// @Accept       json
// @Produce      json
// @Param        id path string true "user ID"
// @Success      200 {object} model.UserModel "Successful operation"
// @Failure      400 {object} utils.ErrorResponse "Bad request"
// @Failure      500 {object} utils.ErrorResponse "Internal server error"
// @Router       /user/{id} [get]
func (e *userHandler) GetById(c echo.Context) error {

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

	user, err := e.service.FindOne(ctx, filter)
	{
		if err != nil {
			return http.HTTPError(err).BadRequest()
		}
	}

	return http.Response(c).OK(user)
}

// Update godoc
// @Summary      Update user information
// @Description  Update user information by ID
// @Tags 		 user
// @ID           update-user
// @Accept       json
// @Param        id path string true "user ID"
// @Param        input body dto.UpdateUserDto true "user information"
// @Success      204 "Successful operation"
// @Failure      400 {object} utils.ErrorResponse "Bad request"
// @Failure      500 {object} utils.ErrorResponse "Internal server error"
// @Router       /user/{id} [patch]
func (e *userHandler) Update(c echo.Context) error {
	var id int64
	{
		if !http.PathValue(c.Param("id")).TryInt64(&id) {
			err := errors.New("parse id error")
			return http.HTTPError(err).BadRequest()
		}
	}

	var updateDto dto.UpdateUserDto
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
// @Summary      Delete user
// @Description  Delete user by ID
// @Tags 		 user
// @ID           delete-user
// @Accept       json
// @Param        id path string true "user ID"
// @Success      204 "Successful operation"
// @Failure      400 {object} utils.ErrorResponse "Bad request"
// @Failure      500 {object} utils.ErrorResponse "Internal server error"
// @Router       /user/{id} [delete]
func (e *userHandler) Delete(c echo.Context) error {
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
// @Summary      put user
// @Description  put user by ID
// @Tags 		 user
// @ID           change-visibility-user
// @Accept       json
// @Param        id path string true "user ID"
// @Success      204 "Successful operation"
// @Failure      400 {object} utils.ErrorResponse "Bad request"
// @Failure      500 {object} utils.ErrorResponse "Internal server error"
// @Router       /user/{id} [put]
func (e *userHandler) ChangeVisibility(c echo.Context) error {
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
