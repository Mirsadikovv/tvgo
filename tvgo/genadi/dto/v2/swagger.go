package v2

type SwaggerObject struct {
	Swagger     string                      `json:"swagger" yaml:"swagger"`
	Info        InfoObject                  `json:"info" yaml:"info"`
	Host        string                      `json:"host,omitempty" yaml:"host"`
	BasePath    string                      `json:"basePath" yaml:"basePath"`
	Schemes     []string                    `json:"schemes,omitempty" yaml:"schemes"`
	Consumes    []string                    `json:"consumes,omitempty" yaml:"consumes"`
	Produces    []string                    `json:"produces,omitempty" yaml:"produces"`
	Paths       PathItemObject              `json:"paths" yaml:"paths"`
	Definitions map[string]DefinitionObject `json:"definitions" yaml:"definitions"`
}

type InfoObject struct {
	Title          string         `json:"title" yaml:"title"`
	Version        string         `json:"version" yaml:"version"`
	Description    string         `json:"description" yaml:"description"`
	TermsOfService string         `json:"termsOfService,omitempty" yaml:"termsOfService"`
	Contact        *ContactObject `json:"contact,omitempty" yaml:"contact"`
	License        *LicenseObject `json:"license,omitempty" yaml:"license"`
}

type ContactObject struct {
	Name  string `json:"name,omitempty" yaml:"name"`
	Url   string `json:"url,omitempty" yaml:"url"`
	Email string `json:"email,omitempty" yaml:"email"`
}

type LicenseObject struct {
	Name string `json:"name" yaml:"name"`
	Url  string `json:"url" yaml:"url"`
}

type DefinitionObject struct {
	Type       string                    `json:"type" yaml:"type"`
	Properties map[string]PropertyObject `json:"properties" yaml:"properties"`
}

type PropertyObject struct {
	Type  string         `json:"type,omitempty" yaml:"type"`
	Ref   string         `json:"$ref,omitempty" yaml:"$ref"`
	Items map[string]any `json:"items,omitempty"`
}

type PathItemObject map[string]map[string]OperationObject

type OperationObject struct {
	Tags        []string                  `json:"tags,omitempty"`        // Теги
	Summary     string                    `json:"summary,omitempty"`     // Краткое описание
	Description string                    `json:"description,omitempty"` // Подробное описание
	OperationID string                    `json:"operationId,omitempty"` // Уникальный ID операции
	Consumes    []string                  `json:"consumes,omitempty"`    // Поддерживаемые MIME-типы (вход)
	Produces    []string                  `json:"produces,omitempty"`    // Поддерживаемые MIME-типы (выход)
	Parameters  []ParameterObject         `json:"parameters,omitempty"`  // Параметры запроса
	Responses   map[string]ResponseObject `json:"responses"`             // Описания ответов
	Schemes     []string                  `json:"schemes,omitempty"`     // Схемы (http, https, ws, wss)
	Deprecated  bool                      `json:"deprecated,omitempty"`  // Устаревшая операция
	Security    []map[string][]string     `json:"security,omitempty"`    // Требования безопасности
}

type ParameterObject struct {
	Name        string     `json:"name"` // Имя параметра
	In          string     `json:"in"`   // Где находится параметр (query, header, path, body, formData)
	Description string     `json:"description,omitempty"`
	Required    bool       `json:"required,omitempty"` // Обязателен ли параметр
	Type        string     `json:"type,omitempty"`     // Тип параметра
	Enum        []string   `json:"enum,omitempty"`
	Schema      *SchemaRef `json:"schema,omitempty"` // Ссылка на объект схемы
}

// ResponseObject описывает ответ API
type ResponseObject struct {
	Description string `json:"description"`      // Описание ответа
	Schema      any    `json:"schema,omitempty"` // Схема данных ответа
}

type SchemaRef struct {
	Ref string `json:"$ref,omitempty"` // Ссылка на схему
}
