package models

type Properties struct {
	Properties_key         string `json:"key"`
	Properties_value       string `json:"value"`
	Properties_encrypted   bool   `json:"encrypted"`
	Properties_dynamic     bool   `json:"dynamic"`
	Properties_encryptable bool   `json:"encryptable"`
}
