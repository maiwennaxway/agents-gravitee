/*
 * Environments API
 *
 * By default, gravitee organizations are provisioned with two environments: `test` and `prod`. An environment provides a runtime execution context for APIs. An API revision must be deployed to an environment before it can be accessed at runtime. No constraints are placed on the usage between different environments (`test` versus `prod`, for example). Developers are free to implement and enforce any type or testing, promotion, and deployment procedures that suit their development lifecycle.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// EnvironmentProperties List of properties that can be used to customize the environment.
type EnvironmentProperties struct {
	// Environment property names and values.
	Property []EnvironmentPropertiesProperty `json:"property,omitempty"`
}
