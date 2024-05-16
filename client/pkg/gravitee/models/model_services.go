package models

type Services struct {
	Services_dynamicProperty_schedule string
	// The schedule of the service

	Services_dynamicProperty_provider string
	// The type of the dynamic property provider.

	/*Allowed value:
	HTTP
	Example:
	HTTP*/

	//Services_dynamicProperty_configuration HttpDynamicPropertyProviderConfiguration
	Services_dynamicProperty_configuration_url string
	// The url of the dynamic property provider

	Services_dynamicProperty_configuration_specification interface{}
	// The specification of the dynamic property provider

	Services_dynamicProperty_configuration_useSystemProxy bool
	// Use the system proxy or not

	Services_dynamicProperty_configuration_method string
	// The method of the selector
	Services_dynamicProperty_enabled bool
	// Is the service enabled or not.

	Services_healthCheck_schedule string
	// The schedule of the service

	Services_healthCheck_steps_name string
	// The name of the step

	Services_healthCheck_steps_request_path string
	// The path of the request

	Services_healthCheck_steps_request_method string
	// The method of the selector

	/*Allowed values:
	CONNECT
	DELETE
	GET
	HEAD
	OPTIONS
	PATCH
	POST
	PUT
	TRACE
	OTHER
	Example:
	GET
	*/

	Services_healthCheck_steps_request_headers_name string
	// The name of the header
	Services_healthCheck_steps_request_headers_value string
	// The value of the header

	Services_healthCheck_steps_request_body string
	// The body of the request

	Services_healthCheck_steps_request_fromRoot bool
	// Is the request from the root or not

	Services_healthCheck_steps_response_assertions []string

	Services_healthCheck_steps_enabled bool
	// Is the service enabled or not.
}
