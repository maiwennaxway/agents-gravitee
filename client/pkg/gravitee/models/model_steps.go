package models

type Steps struct {
	Name     string   `json:"name,omitempty"`
	Request  Request  `json:"request,omitempty"`
	Response Response `json:"response,omitempty"`
}
