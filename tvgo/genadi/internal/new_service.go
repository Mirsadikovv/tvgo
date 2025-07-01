package internal

import (
	"bytes"
	"log"
	"path"

	v2 "git.sriss.uz/shared/genadi/swagger/v2"
	"github.com/Mirsadikovv/tvgo/genadi/dto"
	"github.com/Mirsadikovv/tvgo/genadi/service"
	"github.com/Mirsadikovv/tvgo/genadi/stuble"
	"github.com/iancoleman/strcase"
)

func NewService(serviceName string, jsonData v2.SwaggerObject) []dto.VirtualFile {
	serviceNameUc := strcase.ToCamel(serviceName)

	// dirs := []string{
	// 	path.Join("modules", serviceNameUc, "service"),
	// 	path.Join("modules", serviceNameUc, "pages"),
	// 	path.Join("modules", serviceNameUc, "router"),
	// }

	files := map[string]string{
		path.Join("modules", serviceNameUc, "service", "index.ts"): stuble.Service,
		path.Join("modules", serviceNameUc, "router", "index.ts"):  stuble.Router,
		path.Join("modules", serviceNameUc, "pages", "Create.vue"): stuble.Create,
		path.Join("modules", serviceNameUc, "pages", "Edit.vue"):   stuble.Edit,
		path.Join("modules", serviceNameUc, "pages", "Page.vue"):   stuble.Page,
		path.Join("modules", serviceNameUc, "pages", "View.vue"):   stuble.View,
	}

	var result []dto.VirtualFile

	for filePath, content := range files {
		buf := new(bytes.Buffer)

		m := map[string]any{
			"ServiceName": serviceNameUc,
			"JsonData":    jsonData,
		}

		if err := service.Tmp(content).Execute(buf, m); err != nil {
			log.Fatalln("template execute error:", err)
		}

		result = append(result, dto.VirtualFile{
			Path:    filePath,
			Content: buf,
		})
	}

	return result
}
