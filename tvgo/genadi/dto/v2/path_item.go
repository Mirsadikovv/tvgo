package v2

type Void struct{}

type PathItemTag struct {
	Path        string            `json:"path"`
	Method      string            `json:"method"`
	Summary     string            `json:"summary,omitempty"`
	Description string            `json:"description,omitempty"`
	Consumes    []string          `json:"consumes,omitempty"`
	Produces    []string          `json:"produces,omitempty"`
	Parameters  []ParameterObject `json:"parameters,omitempty"`
}

func (p PathItemObject) SortByTags(names ...string) map[string][]PathItemTag {
	operationObject := map[string][]PathItemTag{}

	var tagName string

	if len(names) > 0 {
		tagName = names[0]
	}

	for url, mob := range p {
		for method, v := range mob {
			tag := v.Tags[0]

			if tagName != "" && tagName != tag {
				continue
			}

			operationObject[tag] = append(operationObject[tag], PathItemTag{
				Path:        url,
				Method:      method,
				Summary:     v.Summary,
				Description: v.Description,
				Consumes:    v.Consumes,
				Produces:    v.Produces,
				Parameters:  v.Parameters,
			})
		}
	}

	return operationObject
}

func (p PathItemObject) Tags() []string {
	tmp := map[string]Void{}

	for _, mob := range p {
		for _, v := range mob {
			for _, tag := range v.Tags {
				tmp[tag] = Void{}
			}
		}
	}

	var tags []string

	for tag := range tmp {
		tags = append(tags, tag)
	}

	return tags
}
