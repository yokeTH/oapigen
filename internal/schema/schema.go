package schema

type SchemaRef struct {
	Ref    string  `json:"$ref,omitzero"`
	Schema *Schema `json:"-"`
}

type Schema struct {
	Type string `json:"type,omitzero"` // object, array, string, integer, number, boolean

	Name        string `json:"name,omitzero"`
	Description string `json:"description,omitzero"`
	Required    bool   `json:"required,omitzero"`

	// type = object
	Properties map[string]SchemaRef `json:"properties,omitzero"`

	// type = array
	Items SchemaRef   `json:"items,omitzero"`
	AllOf []SchemaRef `json:"allOf,omitzero"`
	OneOf []SchemaRef `json:"oneOf,omitzero"`
	AnyOf []SchemaRef `json:"anyOf,omitzero"`
}
