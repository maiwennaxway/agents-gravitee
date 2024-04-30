package gravitee

import (
	"fmt"
	"testing"
	"time"

	"github.com/Axway/agent-sdk/pkg/apic"
	"github.com/maiwennaxway/agents-gravitee/client/pkg/config"

	//"github.com/maiwennaxway/agents-gravitee/client/pkg/gravitee"
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
		config         config.GraviteeConfig
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
			name:  "The API PetStore was found",
			ApiID: "f2e12fc3-fdff-4f8b-a12f-c3fdffef8b17",
		},
		{
			name:  "The API API_Ms was found",
			ApiID: "c6f8c1c6-f530-46ed-b8c1-c6f530f6ed37",
		},
		{
			name:  "The API API_Ms2 was found",
			ApiID: "285cde3d-4340-44d2-9cde-3d4340e4d22a",
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
			client := mockAPIClient{
				t:            t,
				cfg:          &tc.config,
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

			filterFunc := func(map[string]string) bool {
				return !tc.filterFailed
			}

			ApiJob := newPollAPIsJob(client, cache, readyFunc, 10, filterFunc)
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

type mockAPIClient struct {
	t            *testing.T
	cfg          *config.GraviteeConfig
	ApiId        string
	allApiErr    bool
	getApiErr    bool
	specNotFound bool
}

func (m mockAPIClient) GetConfig() *config.GraviteeConfig {
	return m.cfg
}

func (m mockAPIClient) GetApis() (apis Apis, err error) {
	ApiId := m.ApiId
	if ApiId == "" {
		ApiId = "f2e12fc3-fdff-4f8b-a12f-c3fdffef8b17"
	}

	apis = []string{ApiId}
	if m.allApiErr {
		apis = nil
		err = fmt.Errorf("error get all apis")
	}
	return
}

func (m mockAPIClient) GetApi(apiId string) (api *models.Api, err error) {
	apis := map[string]*models.Api{
		"c6f8c1c6-f530-46ed-b8c1-c6f530f6ed37": {
			Id:                "c6f8c1c6-f530-46ed-b8c1-c6f530f6ed37",
			Name:              "API_Ms",
			ApiVersion:        "1.2",
			Description:       "first API, made by maiwenn",
			CrossId:           "",
			DefinitionVersion: "",
			DeployedAt:        "",
			CreatedAt:         "Apr 17, 2024, 10:46:03 AM",
			UpdatedAt:         "",
			EnvironmentId:     "DEFAULT",
			ExecutionMode:     "V2",
			ContextPath:       "/api_ms",
		},
		/* je suis pas sure que celle là marche car elle est stopped "f2e12fc3-fdff-4f8b-a12f-c3fdffef8b17": {
			Id:            "f2e12fc3-fdff-4f8b-a12f-c3fdffef8b17",
			Name:          "petstore",
			ApiVersion:    "1.0.7",
			Description:   "This is a sample server Petstore server.  You can find out more about Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/).  For this sample, you can use the api key `special-key` to test the authorization filters.",
			CreatedAt:     "Apr 16, 2024, 9:52:11 AM",
			EnvironmentId: "DEFAULT",
			ExecutionMode: "V2",
			ContextPath:   "/v2",
		},*/
		"285cde3d-4340-44d2-9cde-3d4340e4d22a": {
			Id:            "285cde3d-4340-44d2-9cde-3d4340e4d22a",
			Name:          "API_Msdeux",
			ApiVersion:    "1.3",
			Description:   "second API, made by maiwenn",
			CreatedAt:     "Apr 25, 2024, 1:06:49 PM",
			EnvironmentId: "DEFAULT",
			ExecutionMode: "V2",
			ContextPath:   "/api_msdeux",
		},
	}
	if m.getApiErr {
		return nil, fmt.Errorf("error get api")
	}
	return apis[apiId], nil
}

func (m mockAPIClient) GetSpecFile(path string) ([]byte, error) {
	assert.Equal(m.t, specPath, path)
	return []byte("spec"), nil
}

func (m mockAPIClient) IsReady() bool { return true }

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
