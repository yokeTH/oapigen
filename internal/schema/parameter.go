package schema

type Parameter struct {
	Name        string `json:"name"`
	In          string `json:"in"` // query | header | path | cookie
	Description string `json:"description"`
	Required    bool   `json:"required"`
}
