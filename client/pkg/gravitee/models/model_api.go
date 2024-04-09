package models

type Api struct {
	//The environment uuId
	EnvironmentId string

	//The execution mode of the API.
	ExecutionMode string

	ContextPath string

	Attributes []Attributes
	//The context path of the API.
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

	PathMappings []string
	// The list of path mappings associated with this API.

	// The list of entrypoints associated with this API.

	Entrypoints_target string
	// The target of the entrypoint.

	Entrypoints_host string
	// The host of the entrypoint.

	Entrypoints_tags []string
	// The list of sharding tags associated with this entrypoint.

	Id string
	// API's uuid.
	Name string
	// API's name. Duplicate names can exists.
	Description string
	// API's description. A short description of your API.

	CrossId string
	// API's crossId. Identifies API across environments.

	ApiVersion string
	// APIs version. Its a simple string only used in the portal.

	DefinitionVersion string
	// required APIs gravitee definition version.

	/*Allowed values:
	V1
	V2
	V4
	Example:
	V4*/

	DeployedAt string
	/*<date-time>
	The last date (as timestamp) when the API was deployed.*/

	CreatedAt string
	/*<date-time>
	The date (as timestamp) when the API was created.*/

	UpdatedAt string
	/*<date-time>
	The last date (as timestamp) when the API was updated.*/

	DisableMembershipNotifications bool
	//Disable membership notifications.

	Groups []string
	// APIs groups. Used to add team in your API.

	State string
	// The state of the API regarding the gateway(s).

	/*Allowed values:
	CLOSED
	INITIALIZED
	STARTED
	STOPPED
	STOPPING
	Example:
	STARTED*/
	DeploymentState string
	// The deployment state of the API regarding the gateway(s).

	/*Allowed values:
	NEED_REDEPLOY
	DEPLOYED
	Example:
	DEPLOYED*/

	Visibility string
	// The visibility of the resource regarding the portal.

	/*Allowed values:
	PUBLIC
	PRIVATE
	Example:
	PUBLIC
	*/
	Labels []string
	// The free list of labels associated with this API.

	LifecycleState string
	// The status of the API regarding the console.

	/*Allowed values:
	ARCHIVED
	CREATED
	DEPRECATED
	PUBLISHED
	UNPUBLISHED
	Example:
	CREATED*/

	Tags []string
	// The list of sharding tags associated with this API.
	PrimaryOwner_id string
	// Owners uuid.

	PrimaryOwner_email string
	// Owner's email. Can be null if owner is a group.

	PrimaryOwner_displayName string
	// Owners name.

	PrimaryOwner_type string
	// The type of membership

	/*Allowed values:
	USER
	GROUP*/

	Categories []string
	//The list of category ids associated with this API.

	//read-only the context where the api definition was created

	DefinitionContext_origin string
	// The origin of the API.

	/*Allowed values:
	MANAGEMENT
	KUBERNETES
	Example:
	MANAGEMENT*/
	DefinitionContext_mode string
	//The mode of the API. fully_managed: Mode indicating the api is fully managed by the origin and so, only the origin should be able to manage the api. api_definition_only: Mode indicating the api is partially managed by the origin and so, only the origin should be able to manage the api definition part of the api. This includes everything regarding the definition of the apis (plans, flows, metadata, ...)

	/*Allowed values:
	FULLY_MANAGED
	API_DEFINITION_ONLY
	Example:
	FULLY_MANAGED*/

	WorkflowState string
	//read-only The status of the API regarding the review feature.

	/*Allowed values:
	DRAFT
	IN_REVIEW
	REQUEST_FOR_CHANGES
	REVIEW_OK
	Example:
	DRAFT*/

	ResponseTemplates interface{}

	Resources_name string
	//required
	Resources_type string
	//required
	Resources_configuration interface{}
	//required
	Resources_enabled bool

	Properties_key string
	//required
	Properties_value string
	//required
	Properties_encrypted   bool
	Properties_dynamic     bool
	Properties_encryptable bool
	//write-only
	Lnks_pictureUrl string
	//The URL to the API's picture.

	Links_backgroundUrl string
	// The URL to the API's background.
}
