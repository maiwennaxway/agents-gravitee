package models

type DynamicProperty struct {
	Services_dynamicProperty_schedule      string        `json:"schedule"`
	Services_dynamicProperty_provider      string        `json:"provider"`
	Services_dynamicProperty_configuration Configuration `json:"confiration"`
	Services_dynamicProperty_enabled       bool          `json:"enabled"`
}
