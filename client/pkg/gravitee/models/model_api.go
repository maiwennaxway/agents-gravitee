package models

type Api struct {
	//The environment uuId
	EnvironmentId string

	//The execution mode of the API.
	ExecutionMode string

	ContextPath string
	//The context path of the API.

	Proxy_virtualHosts_host string
	//The host of the virtual host

	Proxy_virtualHosts_path string
	//The path of the virtual host

	Proxy_virtualHosts_overrideEntrypoint bool
	//Override the entrypoint or not

	Proxy_groups_name string
	//Endpoint group's name.

	//proxy_groups_endpoints []HttpEndpointV2
	//The list of endpoints associated with this endpoint group.

	Proxy_groups_loadBalancer_type string

	Proxy_groups_Services_discovery_provider string
	//The provider of the service

	Proxy_groups_Services_discovery_configuration interface{}
	//The configuration of the service

	Proxy_groups_Services_discovery_enabled bool
	//Is the service enabled or not.

	Proxy_groups_httpProxy_enabled bool
	//Is the proxy enabled or not

	Proxy_groups_httpProxy_useSystemProxy bool
	//Use the system proxy or not

	Proxy_groups_httpProxy_host string
	//The host of the proxy

	Proxy_groups_httpProxy_port int
	//The port of the proxy

	Proxy_groups_httpProxy_username string
	//The username used to connect to the proxy

	Proxy_groups_httpProxy_password string
	//The password used to connect to the proxy

	Proxy_groups_httpProxy_type string
	//The type of the proxy.

	/*Allowed values:
	*HTTP
	*SOCKS4
	*SOCKS5
	*Example:
	*HTTP
	 */

	Proxy_groups_httpClientOptions_idleTimeout int
	//The idle timeout of the http client in ms

	Proxy_groups_httpClientOptions_connectTimeout int
	//The connect timeout of the http client in ms

	Proxy_groups_httpClientOptions_keepAlive bool
	//The keep alive parameter of the http client

	Proxy_groups_httpClientOptions_readTimeout int
	//The read timeout of the http client in ms

	Proxy_groups_httpClientOptions_pipelining bool
	//The pipelining parameter of the http client

	Proxy_groups_httpClientOptions_maxConcurrentConnections int
	//The max connections of the http client

	Proxy_groups_httpClientOptions_useCompression bool
	//Use compression or not

	Proxy_groups_httpClientOptions_propagateClientAcceptEncoding bool
	//Propagate the client accept encoding or not

	Proxy_groups_httpClientOptions_followRedirects bool
	//Follow redirects or not

	Proxy_groups_httpClientOptions_clearTextUpgrade bool
	//Clear text upgrade or not

	Proxy_groups_httpClientOptions_version string
	//The protocol version.

	/*Allowed values:
	HTTP_1_1
	HTTP_2
	Default:
	HTTP_1_1
	Example:
	HTTP_1_1
	*/

	Proxy_groups_httpClientSslOptions_trustAll bool
	//Trust all certificates or not

	Proxy_groups_httpClientSslOptions_hostnameVerifier bool
	//Should the hostname be verified or not

	//trustStore one of: JKSTrustStore
	Proxy_groups_httpClientSslOptions_trustStore_type string
	//The type of the trust store.

	/*Allowed values:
	JKS
	PEM
	PKCS12
	NONE
	Example:
	JKS
	*/
	Proxy_groups_httpClientSslOptions_trustStore_alias string
	//The alias of the trust store

	Proxy_groups_httpClientSslOptions_trustStore_path string
	//The path of the trust store

	Proxy_groups_httpClientSslOptions_trustStore_content string
	//The content of the trust store

	Proxy_groups_httpClientSslOptions_trustStore_password string
	//The password of the trust store

	// keyStore one of: JKSKeyStore
	Proxy_groups_httpClientOptions_keyStore_type string
	//required The type of the key store.

	/*Allowed values:
	JKS
	PEM
	PKCS12
	NONE
	Example:
	JKS*/

	Proxy_groups_httpClientOptions_keyStore_alias string
	// The alias of the key store

	Proxy_groups_httpClientOptions_keyStore_path string
	// The path of the key store

	Proxy_groups_httpClientOptions_keyStore_content string
	// The content of the key store

	Proxy_groups_httpClientOptions_keyStore_password string
	// The password of the key store

	Proxy_groups_httpClientOptions_keyStore_keyPassword string
	// The key password of the key store

	Proxy_groups_httpClientOptions_headers_name string
	// The name of the header

	Proxy_groups_httpClientOptions_headers_value string
	// The value of the header

	//The list of headers associated with this endpoint group.

	Proxy_groups_headers_name string
	// The name of the header

	Proxy_groups_headers_value string
	// The value of the header

	Proxy_failover_maxAttempts int
	// The maximum number of attempts.

	Proxy_failover_retryTimeout int
	// The retry timeout in ms

	Proxy_failover_cases []string
	// The list of cases associated with this failover.

	/*Allowed value:
	TIMEOUT
	Default:
	TIMEOUT
	*/
	Proxy_cors_allowCredentials bool
	Proxy_cors_allowHeaders     []string
	Proxy_cors_allowMethods     []string
	Proxy_cors_allowOrigin      []string
	Proxy_cors_enabled          bool
	Proxy_cors_exposeHeaders    []string
	Proxy_cors_maxAge           int
	Proxy_cors_runPolicies      bool

	Proxy_logging_mode string
	// The mode of the logging.

	/*Allowed values:
	NONE
	CLIENT
	PROXY
	CLIENT_PROXY
	Example:
	NONE
	*/
	Proxy_logging_scope string
	// The scope of the logging.

	/*Allowed values:
	NONE
	REQUEST
	RESPONSE
	REQUEST_RESPONSE
	Example:
	REQUEST
	*/
	Proxy_logging_content string
	// The content of the logging.

	/*Allowed values:
	NONE
	HEADERS
	PAYLOADS
	HEADERS_PAYLOADS
	Example:
	NONE
	*/
	Proxy_logging_condition string
	//The condition of the logging

	Proxy_stripContextPath bool
	Proxy_preserveHost     bool
	Proxy_servers          []string
	//proxy_paths object
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
