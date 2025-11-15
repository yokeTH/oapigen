package shared

type Route struct {
	Method   string
	Path     string
	Handler  string
	BodyType string
	RespType string
	Params   []string
}
