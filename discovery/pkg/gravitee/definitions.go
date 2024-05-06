package gravitee

//const openapi     = "openapi"

type jobFirstRunDone func() bool

// pour la fonction shouldPublishApi dans pollAPIsJob
const (
	cacheKeyAttribute = "cacheKey"
	agentApiTagName   = "AgentCreated"
	agentApiTagValue  = "true"
)
