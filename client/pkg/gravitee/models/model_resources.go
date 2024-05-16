package models

type Resources struct {
	Resources_name string
	//required
	Resources_type string
	//required
	Resources_configuration interface{}
	//required
	Resources_enabled bool
}
