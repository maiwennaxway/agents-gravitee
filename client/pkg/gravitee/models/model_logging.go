package models

type Logging struct {
	Mode      string `json:"mode,omitempty"`
	Scope     string `json:"scope,omitempty"`
	Content   string `json:"content,omitempty"`
	Condition string `json:"condition,omitempty"`
}
