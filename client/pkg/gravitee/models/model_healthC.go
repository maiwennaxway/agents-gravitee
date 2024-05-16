package models

type HealthCheck struct {
	Schedule string  `json:"schedule,omitempty"`
	Steps    []Steps `json:"steps,omitempty"`
	Enabled  bool    `json:"enabled,omitempty"`
	Inherit  bool    `json:"inherit,omitempty"`
}
