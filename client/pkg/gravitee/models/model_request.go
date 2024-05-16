package models

type Request struct {
	Path     string    `json:"path,omitempty"`
	Method   string    `json:"method,omitempty"`
	Headers  []Headers `json:"headers,omitempty"`
	Body     string    `json:"body,omitempty"`
	FromRoot bool      `json:"fromRoot,omitempty"`
}
