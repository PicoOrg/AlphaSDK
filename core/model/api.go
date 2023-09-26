package model

type API struct {
	Path        string      `json:"path"`
	Method      string      `json:"method"`
	Description string      `json:"description"`
	Consumes    []string    `json:"consumes"`
	Produces    []string    `json:"produces"`
	Tags        []string    `json:"tags"`
	Summary     string      `json:"summary"`
	Parameters  []Parameter `json:"parameters"`
	Responses   []Response  `json:"responses"`
}
