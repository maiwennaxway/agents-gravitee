package models

type HCO struct {
	IdleTimeout                   int    `json:"idleTimeout,omitempty"`
	KeepAliveTimeout              int    `json:"keepAliveTimeout,omitempty"`
	ConnectTimeout                int    `json:"connectTimeout,omitempty"`
	KeepAlive                     bool   `json:"keepAlive,omitempty"`
	ReadTimeout                   int    `json:"readTimeout,omitempty"`
	Pipelining                    bool   `json:"pipelining,omitempty"`
	MaxConcurrentConnections      int    `json:"maxConcurrentConnections,omitempty"`
	UseCompression                bool   `json:"useCompression,omitempty"`
	PropagateClientAcceptEncoding bool   `json:"propagateClientAcceptEncoding,omitempty"`
	FollowRedirects               bool   `json:"followRedirects,omitempty"`
	ClearTextUpgrade              bool   `json:"clearTextUpgrade,omitempty"`
	Version                       string `json:"version,omitempty"`
}
