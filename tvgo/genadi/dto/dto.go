package dto

import "bytes"

type VirtualFile struct {
	Path    string
	Content *bytes.Buffer
} // @name VirtualFile

type CreateGenadiDto struct {
	Swagger string `json:"swagger"`
} // @name CreateGenadiDto
