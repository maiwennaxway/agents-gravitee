package models

type Properties struct {
	Properties_key string
	//required
	Properties_value string
	//required
	Properties_encrypted   bool
	Properties_dynamic     bool
	Properties_encryptable bool
}
