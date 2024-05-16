package models

type HttpEndpointV2 struct {
	Name                 string      `json:"name,omitempty"`
	Target               string      `json:"target,omitempty"`
	Weight               int         `json:"weight,omitempty"`
	Backup               bool        `json:"backup,omitempty"`
	Status               string      `json:"status,omitempty"`
	Tenants              []string    `json:"tenants,omitempty"`
	Type                 string      `json:"type,omitempty"`
	Inherit              bool        `json:"inherit,omitempty"`
	HealthCheck          HealthCheck `json:"healthCheck,omitempty"`
	HttpProxy            HP          `json:"httpProxy,omitempty"`
	HttpClientOptions    HCO         `json:"httpClientOptions,omitempty"`
	HttpClientSslOptions HCSO        `json:"httpClientSslOptions,omitempty"`
	Headers              []Headers   `json:"headers,omitempty"`
}
