package service

import (
	"bytes"
	"fmt"

	"os"

	v2 "github.com/Mirsadikovv/tvgo/genadi/dto/v2"
)

func CreateFiles(serviceName string, jsonData v2.SwaggerObject, files map[string]string) (bytes.Buffer, error) {

	var buffer bytes.Buffer

	for filePath, content := range files {

		file, err := os.Create(filePath)
		{
			if err != nil {
				return buffer, fmt.Errorf("create file error %v", err)
			}
		}
		defer file.Close()

		m := map[string]any{
			"ServiceName": serviceName,
			"JsonData":    jsonData,
		}

		{
			if err := Tmp(content).Execute(&buffer, m); err != nil {
				return buffer, fmt.Errorf("content copy error %v", err)
			}

			if _, err := file.Write(buffer.Bytes()); err != nil {
				return buffer, fmt.Errorf("write content error %v", err)
			}
		}
	}

	return buffer, nil
}
