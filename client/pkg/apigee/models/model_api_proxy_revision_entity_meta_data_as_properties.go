/*
 * API Proxies API
 *
 * Manage API proxies. You expose APIs on gravitee Edge by implementing API proxies.  API proxies decouple the app-facing API from your backend services, shielding those apps from backend code changes. As you make backend changes to your services, apps continue to call the same API without any interruption. For more information, see <a href=\"https://docs.gravitee.com/api-platform/fundamentals/understanding-apis-and-api-proxies\">Understanding APIs and API proxies</a>.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// ApiProxyRevisionEntityMetaDataAsProperties Kev-value map of metadata describing the API proxy revision.
type ApiProxyRevisionEntityMetaDataAsProperties struct {
	// Type of bundle. Set to `zip`.
	BundleType string `json:"bundle_type,omitempty"`
	// Time when the API proxy revision was created in milliseconds since epoch.
	CreatedAt string `json:"createdAt,omitempty"`
	// Email address of developer that created the API proxy.
	CreatedBy string `json:"createdBy,omitempty"`
	// Time when the API proxy version was last modified in milliseconds since epoch.
	LastModifiedAt string `json:"lastModifiedAt,omitempty"`
	// Email address of developer that last modified the API proxy.
	LastModifiedBy string `json:"lastModifiedBy,omitempty"`
	// Set to `null`.
	SubType string `json:"subType,omitempty"`
}
