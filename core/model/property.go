package model

type Property struct {
	Type       string     `json:"type"` // ["integer", "string", "boolean", "object"]
	Example    any        `json:"example"`
	Default    any        `json:"default"`
	Properties []Property `json:"properties"`
	Schema     string     `json:"schema"`
}
