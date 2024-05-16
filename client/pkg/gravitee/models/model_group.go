package models

type Groups struct {
	Name                 string           `json:"name,omitempty"`
	Endpoints            []HttpEndpointV2 `json:"endpoints,omitempty"`
	TypeofloadBalancer   string           `json:"type,omitempty"`
	Services             Services         `json:"services,omitempty"`
	HttpProxy            HP               `json:"httpProxy,omitempty"`
	HttpClientOptions    HCO              `json:"httpClientOptions,omitempty"`
	HttpClientSslOptions HCSO             `json:"httpClientSslOptions,omitempty"`
	Headers              Headers          `json:"headers,omitempty"`
}
