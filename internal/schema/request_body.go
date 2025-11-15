package schema

type RequestBody struct {
	Description string               `json:"description,omitzero"`
	Content     map[string]MediaType `json:"content,omitzero"`
	Required    bool                 `json:"required,omitzero"`
	Ref         string               `json:"$ref,omitzero"`
}

type MediaType struct {
	Schema  *SchemaRef `json:"schema,omitzero"`
	Example any        `json:"example,omitzero"`
}
