package schema

type OpenAPI struct {
	OpenAPI string   `json:"openapi"`
	Info    Info     `json:"info"`
	Servers []Server `json:"servers"`
	Paths   Paths    `json:"paths"`
}
