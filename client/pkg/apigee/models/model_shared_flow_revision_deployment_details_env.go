/*
 * Deployments
 *
 * Manage API proxy and shared flow deployments.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// SharedFlowRevisionDeploymentDetailsEnv Shared flow revision deployed environment details.
type SharedFlowRevisionDeploymentDetailsEnv struct {
	// Environment details.
	Environment []SharedFlowRevisionDeploymentDetailsEnvEnvironment `json:"environment,omitempty"`
	// Revision of the shared flow.
	Name string `json:"name,omitempty"`
	// Name of the organization.
	Organization string `json:"organization,omitempty"`
	// Name of the shared flow.
	SharedFlow string `json:"sharedFlow,omitempty"`
}
