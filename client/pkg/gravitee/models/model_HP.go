package models

type HP struct {
	Enabled        bool   `json:"enable,omitempty"`
	UseSystemProxy bool   `json:"useSystemProxy,omitempty"`
	Host           string `json:"host,omitempty"`
	Port           int    `json:"port,omitempty"`
	Username       string `json:"username,omitempty"`
	Password       string `json:"password,omitempty"`
	Type           string `json:"type,omitempty"`
}
