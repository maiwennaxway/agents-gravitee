/*
 * Deployments
 *
 * Manage API proxy and shared flow deployments.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// ApiProxyRevisionDeploymentDetailsEnv API proxy revision deployment details in an environment.
type ApiProxyRevisionDeploymentDetailsEnv struct {
	// Name of the API proxy.
	APIProxy      string                                   `json:"aPIProxy,omitempty"`
	Configuration ApiProxyRevisionDeploymentsConfiguration `json:"configuration,omitempty"`
	// Name of the environment.
	Environment string `json:"environment,omitempty"`
	// Revision of the API proxy.
	Name string `json:"name,omitempty"`
	// Name of the organization.
	Organization string `json:"organization,omitempty"`
	// Revision of the API proxy.
	Revision string `json:"revision,omitempty"`
	// Used by gravitee support to identify servers that support the API proxy or shared flow deployment.
	Server []ApiProxyRevisionDeploymentsServer `json:"server,omitempty"`
	// Deployment status, such as `deployed` or `undeployed`.
	State string `json:"state,omitempty"`
}
