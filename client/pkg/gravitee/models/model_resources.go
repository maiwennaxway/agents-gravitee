package models

type Resources struct {
	Resources_name          string        `json:"name"`
	Resources_type          string        `json:"type"`
	Resources_configuration Configuration `json:"configuration"`
	Resources_enabled       bool          `json:"enabled"`
}
