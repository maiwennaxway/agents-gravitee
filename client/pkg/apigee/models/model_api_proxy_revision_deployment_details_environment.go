/*
 * Deployments
 *
 * Manage API proxy and shared flow deployments.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// ApiProxyRevisionDeploymentDetailsEnvironment struct for ApiProxyRevisionDeploymentDetailsEnvironment
type ApiProxyRevisionDeploymentDetailsEnvironment struct {
	Configuration ApiProxyRevisionDeploymentsConfiguration `json:"configuration,omitempty"`
	// Revision of the API proxy.
	Name string `json:"name,omitempty"`
	// Used by gravitee support to identify servers that support the API proxy or shared flow deployment.
	Server []ApiProxyRevisionDeploymentsServer `json:"server,omitempty"`
	// Deployment status, such as `deployed` or `undeployed`.
	State string `json:"state,omitempty"`
}
