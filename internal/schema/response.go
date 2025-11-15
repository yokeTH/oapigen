package schema

type Responses map[string]Response

type Response struct {
	Description string                `json:"description"`
	Content     map[string]*MediaType `json:"content,omitzero"`
	Ref         string                `json:"$ref,omitzero"`
}
