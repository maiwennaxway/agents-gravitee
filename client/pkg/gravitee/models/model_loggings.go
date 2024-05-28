package models

type Loggings struct {
	Logging_condition        string  `json:"condition"`
	Logging_messageCondition string  `json:"messageCondition"`
	Logging_content          Content `json:"content"`
	Logging_Phase            Phase   `json:"phase"`
	Logging_Mode             Mode    `json:"mode"`
}
