package gravitee

import (
	"fmt"
	"testing"
	"time"

	"github.com/Axway/agent-sdk/pkg/apic"
	"github.com/maiwennaxway/agents-gravitee/client/pkg/config"
	"github.com/maiwennaxway/agents-gravitee/client/pkg/gravitee"
	"github.com/maiwennaxway/agents-gravitee/client/pkg/gravitee/models"
	"github.com/stretchr/testify/assert"
)

const (
	specPath = "/path/to/spec"
)

func Test_pollAPIsJob(t *testing.T) {
	tests := []struct {
		name           string
		ApiID          string
		config         *config.GraviteeConfig
		allApiErr      bool
		getApiErr      bool
		specNotFound   bool
		filterFailed   bool
		specNotInCache bool
		apiPublished   bool
	}{
		{
			name: "All apis were found",
		},
		{
			name:         "All specs were found",
			specNotFound: false,
		},
		{
			name:  "The API 10101 was found",
			ApiID: "10101",
		},
		{
			name:      "should stop when getting all apis details fails",
			getApiErr: true,
		},
		{
			name:         "do not publish when should publish check fails",
			filterFailed: true,
		},
		{
			name:      "should stop when getting all apis fails",
			allApiErr: true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			client := mockApiClient{
				t:            t,
				cfg:          tc.config,
				ApiId:        tc.ApiID,
				allApiErr:    tc.allApiErr,
				getApiErr:    tc.getApiErr,
				specNotFound: tc.specNotFound,
			}

			cache := mockApiCache{
				specNotInCache: tc.specNotInCache,
			}

			readyFunc := func() bool {
				return true
			}

			ApiJob := newPollAPIsJob(client, cache, readyFunc, 10)
			assert.False(t, ApiJob.FirstRunComplete())

			ApiJob.isPublishedFunc = func(id string) bool {
				return tc.apiPublished
			}

			err := ApiJob.Execute()
			if tc.allApiErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}

		})
	}
}

type mockApiClient struct {
	t            *testing.T
	cfg          *config.GraviteeConfig
	ApiId        string
	allApiErr    bool
	getApiErr    bool
	specNotFound bool
}

func (m mockApiClient) GetConfig() *config.GraviteeConfig {
	return m.cfg
}

func (m mockApiClient) GetApis() (apis gravitee.Apis, err error) {
	ApiId := m.ApiId
	if ApiId == "" {
		ApiId = "RTE"
	}

	apis = []string{ApiId}
	if m.allApiErr {
		apis = nil
		err = fmt.Errorf("error get all apis")
	}
	return
}

func (m mockApiClient) GetApi(apiId string) (api *models.Api, err error) {
	apis := map[string]*models.Api{
		"api1": {},
	}
	if m.getApiErr {
		return nil, fmt.Errorf("error get api")
	}
	return apis[apiId], nil
}

func (m mockApiClient) GetSpecFile(path string) ([]byte, error) {
	assert.Equal(m.t, specPath, path)
	return []byte("spec"), nil
}

func (m mockApiClient) IsReady() bool { return false }

type mockApiCache struct {
	specNotInCache bool
}

func (m mockApiCache) GetSpecWithName(name string) (*specItem, error) {
	if m.specNotInCache {
		return nil, fmt.Errorf("spec not in cache")
	}
	return &specItem{
		ID:          "id",
		Name:        "name",
		ContentPath: "/path/to/spec",
		ModDate:     time.Now(),
	}, nil
}

func (m mockApiCache) AddPublishedServiceToCache(cacheKey string, serviceBody *apic.ServiceBody) {
}

func (m mockApiCache) AddProductToCache(name string, modDate time.Time, specHash string) {
}

func (m mockApiCache) HasProductChanged(name string, modDate time.Time, specHash string) bool {
	return true
}

func (m mockApiCache) GetProductWithName(name string) (*ApiCacheItem, error) {
	return nil, nil
}
