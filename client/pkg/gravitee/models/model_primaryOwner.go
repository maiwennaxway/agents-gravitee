package models

type PrimaryOwner struct {
	PrimaryOwner_id          string `json:"id,omitempty"`
	PrimaryOwner_email       string `json:"email,omitempty"`
	PrimaryOwner_displayName string `json:"name,omitempty"`
	PrimaryOwner_type        string `json:"type,omitempty"`
}
