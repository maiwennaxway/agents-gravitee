package models

type PathOperator struct {
	Path     string `json:"path,omitempty"`
	Operator string `json:"operator,omitempty"`
}
