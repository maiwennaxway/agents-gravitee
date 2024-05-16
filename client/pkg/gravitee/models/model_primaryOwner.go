package models

type PrimaryOwner struct {
	PrimaryOwner_id string `json:"id,omitempty"`
	// Owners uuid.

	PrimaryOwner_email string `json:"email,omitempty"`
	// Owner's email. Can be null if owner is a group.

	PrimaryOwner_displayName string `json:"name,omitempty"`
	// Owners name.

	PrimaryOwner_type string `json:"type,omitempty"`
	// The type of membership

	/*Allowed values:
	USER
	GROUP*/
}
