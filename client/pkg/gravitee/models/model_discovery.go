package models

type Discovery struct {
	Provider      string        `json:"provider,omitempty"`
	Configuration Configuration `json:"configuration,omitempty"`
	Enabled       bool          `json:"enabled,omitempty"`
}
