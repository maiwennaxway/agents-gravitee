package gravitee

import (
	"time"

	coreapi "github.com/Axway/agent-sdk/pkg/api"
	"github.com/Axway/agent-sdk/pkg/jobs"
)

// graviteeClient - Represents the Gateway client
type graviteeClient struct {
	cfg         *config.graviteeConfig
	apiClient   coreapi.Client
	accessToken string
	envToURLs   map[string][]string
	isReady     bool
}

// NewClient - Creates a new Gateway Client
func NewClient(graviteeCfg *config.graviteeConfig) (*graviteeClient, error) {
	client := &graviteeClient{
		apiClient: coreapi.NewClient(nil, ""),
		cfg:       graviteeCfg,
		envToURLs: make(map[string][]string),
		isReady:   false,
	}

	// create the auth job and register it
	authentication := newAuthJob(
		withAPIClient(client.apiClient),
		withUsername(graviteeCfg.Auth.GetUsername()),
		withPassword(graviteeCfg.Auth.GetPassword()),
		withURL(graviteeCfg.Auth.GetURL()),
		withAuthServerUsername(graviteeCfg.Auth.GetServerUsername()),
		withAuthServerPassword(graviteeCfg.Auth.GetServerPassword()),
		withTokenSetter(client.setAccessToken),
	)
	jobs.RegisterIntervalJobWithName(authentication, 10*time.Minute, "gravitee Auth Token")

	return client, nil
}

func (a *graviteeClient) setAccessToken(token string) {
	a.accessToken = token
	a.isReady = true
}

// GetConfig - return the gravitee client config
func (a *graviteeClient) GetConfig() *config.graviteeConfig {
	return a.cfg
}

// IsReady - returns true when the gravitee client authenticates
func (a *graviteeClient) IsReady() bool {
	return a.isReady
}
