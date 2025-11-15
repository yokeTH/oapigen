package schema

type Paths map[string]*PathItem

type PathItem struct {
	Get     Operation `json:"get,omitzero"`
	Put     Operation `json:"put,omitzero"`
	Post    Operation `json:"post,omitzero"`
	Delete  Operation `json:"delete,omitzero"`
	Options Operation `json:"options,omitzero"`
	Head    Operation `json:"head,omitzero"`
	Patch   Operation `json:"patch,omitzero"`
	Trace   Operation `json:"trace,omitzero"`
}
