/*
 * Shared flows and flow hooks API
 *
 * Manage shared flows and flow hooks. For more information, see: * <a href=\"https://docs.gravitee.com/api-platform/fundamentals/shared-flows\">Reusable shared flows</a> * <a href=\"https://docs.gravitee.com/api-platform/fundamentals/flow-hooks\">Attaching a shared flow using a flow hook</a>.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// SharedFlowMetadata Shared flow metadata.
type SharedFlowMetadata struct {
	// Time when the shared flow was created in milliseconds since epoch.
	CreatedAt int `json:"createdAt,omitempty"`
	// Email address of developer that created the shared flow.
	CreatedBy string `json:"createdBy,omitempty"`
	// Time when the shared flow was last modified in milliseconds since epoch.
	LastModifiedAt int `json:"lastModifiedAt,omitempty"`
	// Email address of developer that last modified the shared flow.
	LastModifiedBy string `json:"lastModifiedBy,omitempty"`
}
