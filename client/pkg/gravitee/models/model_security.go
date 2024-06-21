package models

type Security struct {
	Type          string        `json:"type"`
	Configuration Configuration `json:"configuration"`
}
