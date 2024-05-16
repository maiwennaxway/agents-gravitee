package models

type HCSO struct {
	TrustAll         bool    `json:"trustAll,omitempty"`
	HostnameVerifier bool    `json:"hostnameVerifiers,omitempty"`
	Headers          Headers `json:"headers,omitempty"`
}
