package models

type Content struct {
	Logging_headers                 bool `json:"header"`
	Logging_content_messageHeaders  bool `json:"messageHeaders"`
	Logging_content_payload         bool `json:"payload"`
	Logging_content_messagePayload  bool `json:"messagePayload"`
	Logging_content_messageMetadata bool `json:"messageMetadata"`
}
