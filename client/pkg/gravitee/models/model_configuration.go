package models

type Configuration struct {
	Url            string `json:"url,omitempty"`
	Specification  string `json:"specification,omitempty"`
	UseSystemProxy bool   `json:"useSystemProxy,omitempty"`
	Method         string `json:"method,omitempty"`
}
