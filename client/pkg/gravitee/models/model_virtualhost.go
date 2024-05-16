package models

type VirtualHosts struct {
	Host               string `json:"host,omitempty"`
	Path               string `json:"path,omitempty"`
	OverrideEntrypoint bool   `json:"overrideEntrypoint"`
}
