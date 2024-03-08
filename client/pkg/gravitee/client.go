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
		apiClient: coreapi.NewClient(nil, ""),
		cfg:       graviteeCfg,
		envToURLs: make(map[string][]string),
		EnvId:     "default",
		isReady:   false,
		orgURL:    graviteeCfg.Auth.URL,
	}

	// create the auth job and register it
	authentication := newAuthJob(
		withAPIClient(client.apiClient),
		withUsername(graviteeCfg.Auth.GetUsername()),
		withPassword(graviteeCfg.Auth.GetPassword()),
		withURL(graviteeCfg.Auth.GetURL()),
	)
	jobs.RegisterIntervalJobWithName(authentication, 10*time.Minute, "gravitee Auth Token")

	return client, nil
}

func (a *GraviteeClient) setAccessToken(token string) {
	a.accessToken = token
	a.isReady = true
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
