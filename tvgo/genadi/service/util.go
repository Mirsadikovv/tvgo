package service

import (
	"archive/zip"
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/Mirsadikovv/tvgo/genadi/dto"
	"github.com/iancoleman/strcase"
)

func ToLowerCamel(s string) string {
	return strcase.ToLowerCamel(s)
}

func ToCamel(s string) string {
	return strcase.ToCamel(s)
}

func ToSnake(s string) string {
	return strcase.ToSnake(s)
}

func ToKebab(s string) string {
	return strcase.ToKebab(s)
}

func WithSpace(s string) string {
	return strcase.ToDelimited(s, ' ')
}

func Contains(str, substr string) bool {
	return strings.Contains(str, substr)
}

func Default(value interface{}, fallback interface{}) interface{} {
	if value == nil {
		return fallback
	}
	return value
}

func SwaggerTypeToTSType(swaggerType string) string {
	switch swaggerType {
	case "string":
		return "string"
	case "integer":
		return "number"
	case "number":
		return "number"
	case "boolean":
		return "boolean"
	case "array":
		return "any[]"
	case "object":
		return "{ [key: string]: any }"
	default:
		return "any"
	}
}

func ExtractDtoName(ref string) string {
	const prefix = "#/definitions/"

	startIndex := strings.Index(ref, prefix)
	if startIndex == -1 {
		return ""
	}

	startIndex += len(prefix)
	dtoName := ref[startIndex:]
	dtoName = strings.TrimSpace(dtoName)

	return dtoName
}

func ReplacePathParams(path string) string {
	path = strings.ReplaceAll(path, "{", "${")
	return path
}

func NormalizeModelName(name string) string {
	return strings.TrimPrefix(name, "model.")
}

func hasPrefix(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

func dict(values ...interface{}) (map[string]interface{}, error) {
	if len(values)%2 != 0 {
		return nil, fmt.Errorf("dict requires an even number of arguments")
	}
	d := make(map[string]interface{})
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, fmt.Errorf("dict keys must be strings")
		}
		d[key] = values[i+1]
	}
	return d, nil
}

func set(m map[string]interface{}, key string, value interface{}) map[string]interface{} {
	m[key] = value
	return m
}

func Tmp(data string) *template.Template {
	return template.Must(template.New("").Delims("<<", ">>").Funcs(template.FuncMap{
		"toCamel":            ToCamel,
		"toLowerCamel":       ToLowerCamel,
		"toSnake":            ToSnake,
		"toKebab":            ToKebab,
		"extractDtoName":     ExtractDtoName,
		"withSpace":          WithSpace,
		"swaggerType":        SwaggerTypeToTSType,
		"toLower":            strings.ToLower,
		"toUpper":            strings.ToUpper,
		"contains":           Contains,
		"replacePath":        ReplacePathParams,
		"default":            Default,
		"normalizeModelName": NormalizeModelName,
		"hasPrefix":          hasPrefix,
		"dict":               dict,
		"set":                set,
	}).Parse(data))
}

func ArchiveFilesToBuffer(files []dto.VirtualFile) (*bytes.Buffer, error) {
	var zipBuffer bytes.Buffer
	zipWriter := zip.NewWriter(&zipBuffer)

	for _, file := range files {
		f, err := zipWriter.Create(file.Path)
		if err != nil {
			return nil, err
		}
		_, err = f.Write(file.Content.Bytes())
		if err != nil {
			return nil, err
		}
	}

	err := zipWriter.Close()
	if err != nil {
		return nil, err
	}

	return &zipBuffer, nil
}
