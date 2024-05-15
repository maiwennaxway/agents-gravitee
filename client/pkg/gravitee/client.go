package gravitee

import (
	"time"

	coreapi "github.com/Axway/agent-sdk/pkg/api"
	"github.com/Axway/agent-sdk/pkg/jobs"
	"github.com/maiwennaxway/agents-gravitee/client/pkg/config"
)

// GraviteeClient - Represents the Gateway client
type GraviteeClient struct {
	cfg         *config.GraviteeConfig
	apiClient   coreapi.Client
	accessToken string
	EnvId       string
	envToURLs   map[string][]string
	isReady     bool
	orgURL      string
}

// NewClient - Creates a new Gateway Client
func NewClient(graviteeCfg *config.GraviteeConfig) (*GraviteeClient, error) {
	client := &GraviteeClient{
		apiClient:   coreapi.NewClient(nil, ""),
		cfg:         graviteeCfg,
		accessToken: "",
		envToURLs:   make(map[string][]string),
		EnvId:       "DEFAULT",
		isReady:     false,
		orgURL:      graviteeCfg.Auth.GetURL(),
	}
	// create the auth job and register it
	authentication := newAuthJob(
		withAPIClient(client.apiClient),
		withURL(client.orgURL+"/environments/DEFAULT/apis"),
		withToken("Authorization: Bearer "+graviteeCfg.Auth.GetToken()),
	)
	client.isReady = true
	jobs.RegisterIntervalJobWithName(authentication, 10*time.Minute, "Gravitee Auth Token")
	return client, nil
}

// GetEnvId - get the developer id to be used when creating apps
func (a *GraviteeClient) GetEnvId() string {
	return a.EnvId
}

// GetConfig - return the gravitee client config
func (a *GraviteeClient) GetConfig() *config.GraviteeConfig {
	return a.cfg
}

// IsReady - returns true when the gravitee client authenticates
func (a *GraviteeClient) IsReady() bool {
	return a.isReady
}
