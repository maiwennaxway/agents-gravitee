package gravitee

import (
	"testing"

	"github.com/Axway/agent-sdk/pkg/api"
	"github.com/maiwennaxway/agents-gravitee/client/pkg/config"

	"github.com/stretchr/testify/assert"
)

func createTestClient(t *testing.T, mockClient *api.MockHTTPClient) *GraviteeClient {
	cfg := &config.GraviteeConfig{
		EnvName: "DEFAULT",
		URL:     "http://test.com",
		Auth: &config.AuthConfig{
			Token: "",
		},
	}

	c, err := NewClient(cfg)
	assert.Nil(t, err)
	assert.NotNil(t, c)
	assert.Equal(t, c.GetConfig(), cfg)
	c.apiClient = mockClient

	return c
}

func TestNewClient(t *testing.T) {
	c := createTestClient(t, nil)
	assert.True(t, c.IsReady())
}
