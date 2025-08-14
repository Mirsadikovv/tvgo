package handler

import (
	"errors"
	"strconv"
	"strings"

	"github.com/Mirsadikovv/tvgo/file_service/dto"
	"github.com/Mirsadikovv/tvgo/file_service/service"
	_ "github.com/Mirsadikovv/tvgo/utils"
	"github.com/Mirsadikovv/tvgo/utils/middleware"
	_ "github.com/fobus1289/ufa_shared/http/response"

	"github.com/fobus1289/ufa_shared/http"
	"github.com/fobus1289/ufa_shared/http/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type fileHandler struct {
	service service.FileService
}

func NewHandler(router *echo.Group, service service.FileService) {
	group := router.Group("/file", middleware.XKeyMiddleware)
	{
		handler := &fileHandler{service: service}

		group.POST("", handler.Create)
		router.GET("/file/page", handler.Page)
		router.GET("/file/:id", handler.GetById)
		router.GET("/file/download/:id", handler.Download)
		router.GET("/file/download/:category/:owner", handler.DownloadByOwner)
		group.GET("/search", handler.Search)
		group.PATCH("/:id", handler.Update)
		group.PATCH("/:id/replace", handler.Replace)
		group.PUT("/:id", handler.ChangeVisibility)
		group.DELETE("/:id", handler.Delete)
	}
}

// Create godoc
// @Summary      Create a new file
// @Description  Create file
// @Tags 		 file
// @ID           create-file
// @Accept       json
// @Produce      json
// @Param        input formData dto.CreateFileDto true "file information"
// @Param 		 file formData file true "File Document"
// @Param        X-Key header string true "API Key for authentication"
// @Success      201 {object} utils.ID "Successful operation"
// @Failure      400 {object} utils.ErrorResponse "Bad request"
// @Failure      500 {object} utils.ErrorResponse "Internal server error"
// @Router       /file [post]
func (e *fileHandler) Create(c echo.Context) error {
	var createDto dto.CreateFileDto
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
// @Summary      Get all files with pagination
// @Description  Get all files with pagination
// @Tags 		 file
// @ID           get-all-files
// @Accept       json
// @Produce      json
// @Param        page query string false "Page number" default(1)
// @Param        perpage query string false "Number of items per page" default(10)
// @Param        file_query_params query dto.FileQueryParams false "Searching by params"
// @Param        X-Key header string false "API Key for authentication"
// @Success      200 {object} utils.ID "Successful operation"
// @Failure      400 {object} utils.ErrorResponse "Bad request"
// @Failure      500 {object} utils.ErrorResponse "Internal server error"
// @Router       /file/page [get]
func (e *fileHandler) Page(c echo.Context) error {
	var (
		page     = c.QueryParam("page")
		perPage  = c.QueryParam("perpage")
		paginate = http.NewPaginate(page, perPage)
		ctx      = c.Request().Context()
		params   dto.FileQueryParams
	)

	if err := c.Bind(&params); err != nil {
		return http.HTTPError(err).BadRequest()
	}

	limitFilter := func(tx *gorm.DB) *gorm.DB {
		return tx.Offset(paginate.Skip()).Limit(paginate.Take()).Order("id ASC")
	}

	filter := func(tx *gorm.DB) *gorm.DB {
		tx = tx.Where("is_visible = ?", true)

		if params.Category != nil {
			tx = tx.Where("category = ?", *params.Category)
		}

		if params.Owner != nil {
			tx = tx.Where("owner = ?", *params.Owner)
		}

		if params.Sign != nil {
			tx = tx.Where("sign = ?", *params.Sign)
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
// @Summary      Search all files
// @Description  Search all files
// @Tags 		 file
// @ID           search-all-files
// @Accept       json
// @Produce      json
// @Param        file_query_params query dto.FileQueryParams false "Searching by params"
// @Param        limit  query int    false "Limit the number of results" default(20)
// @Param        X-Key header string true "API Key for authentication"
// @Success      200 {object} utils.ID "Successful operation"
// @Failure      400 {object} utils.ErrorResponse "Bad request"
// @Failure      500 {object} utils.ErrorResponse "Internal server error"
// @Router       /file/search [get]
func (e *fileHandler) Search(c echo.Context) error {
	const (
		defaultLimit = 15
		maxLimit     = 100
	)

	var (
		ctx        = c.Request().Context()
		limitParam = c.QueryParam("limit")
		params     dto.FileQueryParams
	)

	if err := c.Bind(&params); err != nil {
		return http.HTTPError(err).BadRequest()
	}

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

		if params.Category != nil {
			tx = tx.Where("category = ?", *params.Category)
		}

		if params.Owner != nil {
			tx = tx.Where("owner = ?", *params.Owner)
		}

		if params.Sign != nil {
			tx = tx.Where("sign = ?", *params.Sign)
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
// @Summary      Get file by ID
// @Description  Get file by ID
// @Tags 		 file
// @ID           get-file-by-id
// @Accept       json
// @Produce      json
// @Param        id path string true "file ID"
// @Param        X-Key header string false "API Key for authentication"
// @Success      200 {object} model.FileModel "Successful operation"
// @Failure      400 {object} utils.ErrorResponse "Bad request"
// @Failure      500 {object} utils.ErrorResponse "Internal server error"
// @Router       /file/{id} [get]
func (e *fileHandler) GetById(c echo.Context) error {
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

	file, err := e.service.FindOne(ctx, filter)
	{
		if err != nil {
			return http.HTTPError(err).BadRequest()
		}
	}

	return http.Response(c).OK(file)
}

// Download godoc
// @Summary      Download file by ID
// @Description  Download file by ID
// @Tags 		 file
// @ID           download-file-by-id
// @Accept       json
// @Produce      json
// @Param        id path string true "file ID"
// @Param        X-Key header string false "API Key for authentication"
// @Success      200 {object} utils.ID "Successful operation"
// @Failure      400 {object} utils.ErrorResponse "Bad request"
// @Failure      500 {object} utils.ErrorResponse "Internal server error"
// @Router       /file/download/{id} [get]
func (e *fileHandler) Download(c echo.Context) error {
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

	file, err := e.service.FindOne(ctx, filter)
	{
		if err != nil {
			return http.HTTPError(err).BadRequest()
		}
	}

	return c.File(file.BuildFilePath())
}

// DownloadByOwner godoc
// @Summary      Download file by owner
// @Description  Download file by category and owner
// @Tags 		 file
// @ID           download-file-by-owner
// @Accept       json
// @Produce      json
// @Param        category path string true "category"
// @Param        owner path string true "owner"
// @Param        sign query string false "sign"
// @Param        X-Key header string false "API Key for authentication"
// @Success      200 {object} utils.ID "Successful operation"
// @Failure      400 {object} utils.ErrorResponse "Bad request"
// @Failure      500 {object} utils.ErrorResponse "Internal server error"
// @Router       /file/download/{category}/{owner} [get]
func (e *fileHandler) DownloadByOwner(c echo.Context) error {
	category := strings.TrimSpace(c.Param("category"))
	owner := strings.TrimSpace(c.Param("owner"))

	if category == "" {
		err := errors.New("category is required")
		return http.HTTPError(err).BadRequest()
	}

	if owner == "" {
		err := errors.New("owner is required")
		return http.HTTPError(err).BadRequest()
	}

	ctx := c.Request().Context()

	filter := func(tx *gorm.DB) *gorm.DB {
		tx = tx.Where("category = ?", category).Where("owner = ?", owner).Where("is_visible = ?", true)

		if sign := c.QueryParam("sign"); sign != "" {
			tx = tx.Where("sign = ?", sign)
		}

		return tx
	}

	file, err := e.service.FindOne(ctx, filter)
	{
		if err != nil {
			return http.HTTPError(err).BadRequest()
		}
	}

	return c.File(file.BuildFilePath())
}

// Update godoc
// @Summary      Update file information
// @Description  Update file information by ID
// @Tags 		 file
// @ID           update-file
// @Accept       json
// @Param        id path string true "file ID"
// @Param        input body dto.UpdateFileDto true "file information"
// @Param        X-Key header string true "API Key for authentication"
// @Success      204 "Successful operation"
// @Failure      400 {object} utils.ErrorResponse "Bad request"
// @Failure      500 {object} utils.ErrorResponse "Internal server error"
// @Router       /file/{id} [patch]
func (e *fileHandler) Update(c echo.Context) error {
	var id int64
	{
		if !http.PathValue(c.Param("id")).TryInt64(&id) {
			err := errors.New("parse id error")
			return http.HTTPError(err).BadRequest()
		}
	}

	var updateDto dto.UpdateFileDto
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

// Replace godoc
// @Summary      Replace file
// @Description  Replace file by ID
// @Tags 		 file
// @ID           replace-file
// @Accept       json
// @Produce      json
// @Param        id path string true "file ID"
// @Param        input formData dto.ReplaceFileDto true "file information"
// @Param 		 file formData file true "File Document"
// @Param        X-Key header string true "API Key for authentication"
// @Success      204 "Successful operation"
// @Failure      400 {object} utils.ErrorResponse "Bad request"
// @Failure      500 {object} utils.ErrorResponse "Internal server error"
// @Router       /file/{id}/replace [patch]
func (e *fileHandler) Replace(c echo.Context) error {
	var id int64
	{
		if !http.PathValue(c.Param("id")).TryInt64(&id) {
			err := errors.New("parse id error")
			return http.HTTPError(err).BadRequest()
		}
	}

	var replaceDto dto.ReplaceFileDto
	{
		if err := c.Bind(&replaceDto); err != nil {
			return http.HTTPError(err).BadRequest()
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {
		return tx.Where("id", id)
	}

	err := e.service.Replace(&replaceDto, filter)
	{
		if err != nil {
			return http.HTTPError(err).BadRequest()
		}
	}

	return http.Response(c).NoContent()
}

// Delete godoc
// @Summary      Delete file
// @Description  Delete file by ID
// @Tags 		 file
// @ID           delete-file
// @Accept       json
// @Param        id path string true "file ID"
// @Param        X-Key header string true "API Key for authentication"
// @Success      204 "Successful operation"
// @Failure      400 {object} utils.ErrorResponse "Bad request"
// @Failure      500 {object} utils.ErrorResponse "Internal server error"
// @Router       /file/{id} [delete]
func (e *fileHandler) Delete(c echo.Context) error {
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

	if err := e.service.Delete(filter); err != nil {
		return http.HTTPError(err).BadRequest()
	}

	return http.Response(c).NoContent()
}

// ChangeVisibility godoc
// @Summary      Change file visibility
// @Description  Change file visibility by ID
// @Tags 		 file
// @ID           change-visibility-file
// @Accept       json
// @Param        id path string true "file ID"
// @Param        X-Key header string true "API Key for authentication"
// @Success      204 "Successful operation"
// @Failure      400 {object} utils.ErrorResponse "Bad request"
// @Failure      500 {object} utils.ErrorResponse "Internal server error"
// @Router       /file/{id} [put]
func (e *fileHandler) ChangeVisibility(c echo.Context) error {
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
