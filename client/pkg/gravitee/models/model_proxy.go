package models

type Proxy struct {
	VirtualHosts     VirtualHosts `json:"virtualHosts,omitempty"`
	Groups           Groups       `json:"groups,omitempty"`
	Failover         Failover     `json:"failover,omitempty"`
	Cors             Cors         `json:"cors,omitempty"`
	Logging          Logging      `json:"logging,omitempty"`
	StripContextPath bool         `json:"stripContextPath,omitempty"`
	PreserveHost     bool         `json:"preserveHost,omitempty"`
	Servers          []string     `json:"servers,omitempty"`
}
