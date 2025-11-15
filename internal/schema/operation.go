package schema

type Operation struct {
	Summary    string      `json:"summary,omitzero"`
	Parameters []Parameter `json:"parameters,omitzero"`

	RequestBody RequestBody `json:"requestBody,omitzero"`
	Responses   Responses   `json:"responses"`
}
