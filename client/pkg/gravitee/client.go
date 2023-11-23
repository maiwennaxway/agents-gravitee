package gravitee

import (
	"fmt"
	"time"

	coreapi "github.com/Axway/agent-sdk/pkg/api"
	"github.com/Axway/agent-sdk/pkg/jobs"
)

// NewClient - Creates a new Gateway Client
func NewClient(graviteeCfg *graviteeConfig) (*graviteeClient, error) {
	client := &graviteeClient{
		apiClient:   coreapi.NewClient(nil, ""),
		cfg:         graviteeCfg,
		envToURLs:   make(map[string][]string),
		isReady:     false,
		developerID: graviteeCfg.DeveloperID,
		orgURL:      fmt.Sprintf("%s/%s/organizations/%s", graviteeCfg.URL, graviteeCfg.APIVersion, graviteeCfg.Organization),
		dataURL:     graviteeCfg.DataURL,
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

// GetDeveloperID - get the developer id to be used when creating apps
func (a *graviteeClient) GetDeveloperID() string {
	return a.developerID
}

// GetConfig - return the gravitee client config
func (a *graviteeClient) GetConfig() *config.graviteeConfig {
	return a.cfg
}

// IsReady - returns true when the gravitee client authenticates
func (a *graviteeClient) IsReady() bool {
	return a.isReady
}
