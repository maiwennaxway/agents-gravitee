/*
 * Developer apps API
 *
 * Manage developers that register apps.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// DeveloperApp Developer app details.
type DeveloperApp struct {
	// Output only. App family.
	AppFamily string `json:"appFamily,omitempty"`
	// List of API products to which the app is associated (when creating or updating an app). The name of the API product is the name returned when you <a href=\"/docs/api-products/1/routes/organizations/%7Borg_name%7D/apiproducts/get\">list API products</a>. For example, if the Display Name of the API product in the Edge UI is `Premium API Product`, then the name is `premium-api-product` (all lowercase with spaces replaced by hyphens). You can add API products later when you <a href=\"/docs/developer-apps/1/routes/organizations/%7Borg_name%7D/developers/%7Bdeveloper_email%7D/apps/%7Bapp_name%7D/put\">update the developer app</a>. Existing API products are listed in the `credentials` array.
	ApiProducts []string `json:"apiProducts,omitempty"`
	// Output only. ID of the app.
	AppId string `json:"appId,omitempty"`
	// List of attributes used for customizing profile information or for app-specific processing. With gravitee Edge for Public Cloud, the custom attribute limit is 18. Note the folowing:  * `DisplayName` is an attribute that provides the app name in the Edge UI. This can be different from the name (unique ID) of the app. * `Notes` attribute lets you add notes about the developer app. * Any other arbitrary name/value pairs you create are included as custom attributes.
	Attributes []Attribute `json:"attributes,omitempty"`
	// Callback URL used by OAuth 2.0 authorization servers to communicate authorization codes back to apps. See the OAuth 2.0 documentation for more details.
	CallbackUrl string `json:"callbackUrl,omitempty"`
	// Output only. Time the app was created in milliseconds since epoch.
	CreatedAt int `json:"createdAt,omitempty"`
	// Output only. Email address of the developer that created the app.
	CreatedBy string `json:"createdBy,omitempty"`
	// Output only. Set of credentials for the app. Credentials are API key/secret pairs associated with API products.
	Credentials []DeveloperAppCredentials `json:"credentials,omitempty"`
	// Output only. ID of the developer.
	EnvId string `json:"EnvId,omitempty"`
	// Lifetime of the consumer key that will be generated for the developer app, in milliseconds. The default value, `-1`, indicates an infinite validity period. Once set, the expiration can't be updated.
	KeyExpiresIn int `json:"keyExpiresIn,omitempty"`
	// Output only. Last modified time as milliseconds since epoch.
	LastModifiedAt int `json:"lastModifiedAt,omitempty"`
	// Output only. Email of developer that last modified the app.
	LastModifiedBy string `json:"lastModifiedBy,omitempty"`
	// Name of the developer app. Required when creating a developer app; not required when updating a developer app.   The name is used to uniquely identify the app for this organization and developer. Names must begin with an alphanumeric character and can contain letters, numbers, spaces, and the following characters: `. _ # - $ %`. While you can use spaces in the name, we recommend that you use camel case, underscores, or hyphens instead. Otherwise, you will have to URL-encode the app name when you need to include it in the URL of other Edge API calls. See the <a href=\"https://docs.gravitee.com/api-platform/reference/naming-guidelines\">naming restrictions</a>.
	Name string `json:"name"`
	// Scopes to apply to the app. The specified scope names must already exist on the API product that you associate with the app.
	Scopes []string `json:"scopes,omitempty"`
	// Status of the credential.
	Status string `json:"status,omitempty"`
}
