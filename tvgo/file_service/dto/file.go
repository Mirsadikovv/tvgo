package dto

import (
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/Mirsadikovv/tvgo/file_service/model"
	"github.com/fobus1289/ufa_shared/http/response"
)

type PageFileResponseType = response.PaginateResponse[*model.FileModel] // @name PageFileResponseType

type CreateFileDto struct {
	FileHeader *multipart.FileHeader `form:"file" swaggerignore:"true"`
	Category   string                `form:"category"`
	Owner      string                `form:"owner"`
	Sign       string                `form:"sign"`
} // @name CreateFileDto

type UpdateFileDto struct {
	Owner     *string `json:"owner"`
	Category  *string `json:"category"`
	Sign      *string `json:"sign"`
	IsVisible *bool   `json:"isVisible"`
} // @name UpdateFileDto

type ReplaceFileDto struct {
	FileHeader *multipart.FileHeader `form:"file" swaggerignore:"true"`
} // @name ReplaceFileDto

type FileQueryParams struct {
	Sign     *string `query:"sign"`
	Category *string `query:"category"`
	Owner    *string `query:"owner"`
} // @name FileQueryParams

type EmployeeFileParams struct {
	Sign       *string `query:"sign"`
	EmployeeId *int64  `query:"employeeId"`
} // @name EmployeeFileParams

type FileDto struct {
	Id        int64      `json:"id"`
	DirName   string     `json:"dirName"`
	Category  string     `json:"category"`
	Owner     string     `json:"owner"`
	Sign      string     `json:"sign"`
	Filename  string     `json:"filename"`
	MimeType  string     `json:"mimeType"`
	Size      int64      `json:"size"`
	IsVisible bool       `json:"isVisible"`
	CreatedAt *time.Time `json:"createdAt"`
} // @name FileDto

func (f *FileDto) BuildFilePath() string {
	return filepath.Join("./", f.DirName, f.Category, f.Owner, f.Filename)
}
