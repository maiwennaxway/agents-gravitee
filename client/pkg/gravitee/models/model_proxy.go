package models

type Proxy struct {
	VirtualHosts       string
	Host               string
	Path               string
	OverrideEntrypoint bool
	Groups             Groups
	Failover           Failover
	Cors               Cors
	Logging            Logging
	StripContextPath   bool
	PreserveHost       bool
	Servers            []string
}
