package gravitee

const (
	openapi     = "openapi"
	association = "association.json"
)

type Association struct {
	URL string `json:"url"`
}

type APIDetails string
type Api string

type jobFirstRunDone func() bool

const (
	quotaPolicy  = "Quota"
	apiKeyPolicy = "VerifyAPIKey"
	oauthPolicy  = "OAuthV2"
)

const (
	agentProductTagName  = "AgentCreated"
	agentProductTagValue = "true"
)
