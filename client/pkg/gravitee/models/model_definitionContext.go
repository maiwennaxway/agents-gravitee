package models

type DefinitionContext struct {
	DefinitionContext_origin string `json:"origin,omitempty"`
	DefinitionContext_mode   string `json:"mode,omitempty"`
	SyncFrom                 string `json:"syncFrom,omitempty"`
}
