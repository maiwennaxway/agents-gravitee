package models

type Quota struct {
	Limit          int    `json:"limit,omitempty"`
	DynamicLimit   string `json:"dynamicLimit,omitempty"`
	PeriodTime     int    `json:"periodTime"`
	PeriodTimeUnit string `json:"periodTimeUnit"`
}
