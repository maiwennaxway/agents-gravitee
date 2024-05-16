package models

type Cors struct {
	AllowCredentials bool     `json:"allowCredentials,omitempty"`
	AllowHeaders     []string `json:"allowHeaders,omitempty"`
	AllowMethods     []string `json:"allowMethods,omitempty"`
	AllowOrigin      []string `json:"allowOrigin,omitempty"`
	Enabled          bool     `json:"enabled,omitempty"`
	ExposeHeaders    []string `json:"exposeHeaders,omitempty"`
	MaxAge           int      `json:"maxAge,omitempty"`
	RunPolicies      bool     `json:"runPolicies,omitempty"`
}
