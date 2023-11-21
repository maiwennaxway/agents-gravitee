/*
 * API Proxies API
 *
 * Manage API proxies. You expose APIs on gravitee Edge by implementing API proxies.  API proxies decouple the app-facing API from your backend services, shielding those apps from backend code changes. As you make backend changes to your services, apps continue to call the same API without any interruption. For more information, see <a href=\"https://docs.gravitee.com/api-platform/fundamentals/understanding-apis-and-api-proxies\">Understanding APIs and API proxies</a>.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// ApiProxyRevisionResourceFilesResourceFile Resource filename.
type ApiProxyRevisionResourceFilesResourceFile struct {
	// Name of the resource file.
	Name string `json:"name,omitempty"`
	// Type of resource file.
	Type string `json:"type,omitempty"`
}
