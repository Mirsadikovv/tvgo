package genadi_front

import (
	"encoding/json"
	"fmt"

	v2 "github.com/Mirsadikovv/tvgo/genadi/dto/v2"
	"github.com/Mirsadikovv/tvgo/genadi/internal"
	_ "github.com/Mirsadikovv/tvgo/utils"
	"github.com/Mirsadikovv/tvgo/utils/request"
	_ "github.com/fobus1289/ufa_shared/http/response"

	"net/http"

	"github.com/Mirsadikovv/tvgo/genadi/dto"
	"github.com/Mirsadikovv/tvgo/genadi/service"
	"github.com/labstack/echo/v4"
)

type genadiHandler struct {
}

func NewHandler(router *echo.Group) {

	group := router.Group("/genadi")
	{
		handler := &genadiHandler{}

		group.POST("", handler.Create)
	}
}

// Create 		 godoc
// @Summary      Create a new geandi
// @Description  Create geandi
// @Tags 		 geandi
// @ID           create-geandi
// @Accept       json
// @Produce      json
// @Param        input body dto.CreateGenadiDto true "geandi information"
// @Success      201 {object} utils.ID "Successful operation"
// @Failure      400 {object} utils.ErrorResponse "Bad request"
// @Failure      500 {object} utils.ErrorResponse "Internal server error"
// @Router       /genadi [post]
func (e *genadiHandler) Create(c echo.Context) error {
	req := request.Request(c)

	var invitationDto dto.CreateGenadiDto
	{
		if err := req.BindBody(&invitationDto); err != nil {
			return req.BadRequest(err)
		}
	}
	swagger := httpGet[v2.SwaggerObject](invitationDto.Swagger)
	files := generateFiles(swagger)

	zipBuffer, err := service.ArchiveFilesToBuffer(files)
	if err != nil {
		return req.BadRequest(err)
	}

	w := c.Response()
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", `attachment; filename="output.zip"`)

	return c.Blob(http.StatusOK, "application/zip", zipBuffer.Bytes())
}

func generateFiles(swagger v2.SwaggerObject) []dto.VirtualFile {
	var files []dto.VirtualFile

	for _, service := range swagger.Paths.Tags() {
		serviceFiles := internal.NewService(service, swagger)
		files = append(files, serviceFiles...)
	}

	return files
}

func httpGet[T any](url string) T {
	var none T
	if url == "" {
		return none
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err, 1)
		return none
	}

	req.SetBasicAuth("", "")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err, 2)
		return none
	}

	defer response.Body.Close()

	var result T

	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		fmt.Println(err, 3)
		return none
	}

	return result
}
