package model

type Parameter struct {
	Description string   `json:"description"`
	Name        string   `json:"name"`
	In          string   `json:"in"` // ["body", "query"]
	Required    bool     `json:"required"`
	Type        string   `json:"type"` // ["integer", "string", "boolean"]
	Example     any      `json:"example"`
	Default     any      `json:"default"`
	Schema      Property `json:"schema"`
}
