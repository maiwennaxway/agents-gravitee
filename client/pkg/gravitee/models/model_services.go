package models

type Services struct {
	Services_dynamicProperty DynamicProperty `json:"dynamicProperty"`
	Services_healthCheck     HealthCheck     `json:"healthCheck"`
}
