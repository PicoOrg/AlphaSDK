package model

type Response struct {
	StatusCode  int      `json:"status_code"`
	Description string   `json:"description"`
	Schema      Property `json:"schema"`
}
