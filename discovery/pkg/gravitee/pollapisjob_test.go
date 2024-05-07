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
			assert.True(t, ApiJob.FirstRunComplete())

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

func (m mockAPIClient) GetApis() (apis []string, err error) {
	ApiId := m.ApiId
	if ApiId == "" {
		ApiId = "f2e12fc3-fdff-4f8b-a12f-c3fdffef8b17"
	}
	apis = []string{ApiId}
	if m.allApiErr {
		apis = nil
		err = fmt.Errorf("error get all apis")
	}
	return apis, err
}

func (m mockAPIClient) GetApi(apiId, envId string) (api *models.Api, err error) {
	apis := map[string]*models.Api{
		"c6f8c1c6-f530-46ed-b8c1-c6f530f6ed37": {
			DefinitionVersion:              "V2",
			EnvironmentId:                  envId,
			ExecutionMode:                  "V4_EMULATION_ENGINE",
			ContextPath:                    "/api_ms",
			Id:                             "c6f8c1c6-f530-46ed-b8c1-c6f530f6ed37",
			Name:                           "API_Ms",
			Description:                    "first API, made by maiwenn",
			ApiVersion:                     "1.2",
			DeployedAt:                     "2024-04-17T08:46:05.217Z",
			CreatedAt:                      "2024-04-17T08:46:03.658Z",
			UpdatedAt:                      "2024-04-17T08:46:06.064Z",
			DisableMembershipNotifications: false,
			Visibility:                     "PRIVATE",
			LifecycleState:                 "CREATED",
			DefinitionContext_origin:       "MANAGEMENT",
			DefinitionContext_mode:         "FULLY_MANAGED",
			Lnks_pictureUrl:                "http://sl2csoapp1490.pcloud.axway.int:8083/management/v2/environments/DEFAULT/apis/c6f8c1c6-f530-46ed-b8c1-c6f530f6ed37/picture?hash=1713343566064",
			Links_backgroundUrl:            "http://sl2csoapp1490.pcloud.axway.int:8083/management/v2/environments/DEFAULT/apis/c6f8c1c6-f530-46ed-b8c1-c6f530f6ed37/background?hash=1713343566064",
		},
		"f2e12fc3-fdff-4f8b-a12f-c3fdffef8b17": {
			Id:                             "f2e12fc3-fdff-4f8b-a12f-c3fdffef8b17",
			Name:                           "petstore",
			ApiVersion:                     "1.0.7",
			Description:                    "This is a sample server Petstore server.  You can find out more about Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/).  For this sample, you can use the api key `special-key` to test the authorization filters.",
			DeployedAt:                     "2024-04-16T07:54:04.866Z",
			CreatedAt:                      "2024-04-16T07:52:11.643Z",
			UpdatedAt:                      "2024-04-16T07:54:04.866Z",
			DefinitionVersion:              "V2",
			EnvironmentId:                  envId,
			ExecutionMode:                  "V4_EMULATION_ENGINE",
			ContextPath:                    "/v2",
			Visibility:                     "PRIVATE",
			LifecycleState:                 "CREATED",
			DefinitionContext_origin:       "MANAGEMENT",
			DefinitionContext_mode:         "FULLY_MANAGED",
			DisableMembershipNotifications: false,
			Lnks_pictureUrl:                "http://sl2csoapp1490.pcloud.axway.int:8083/management/v2/environments/DEFAULT/apis/f2e12fc3-fdff-4f8b-a12f-c3fdffef8b17/picture?hash=1713254044866",
			Links_backgroundUrl:            "http://sl2csoapp1490.pcloud.axway.int:8083/management/v2/environments/DEFAULT/apis/f2e12fc3-fdff-4f8b-a12f-c3fdffef8b17/background?hash=1713254044866",
		},
		"285cde3d-4340-44d2-9cde-3d4340e4d22a": {
			Id:                             "285cde3d-4340-44d2-9cde-3d4340e4d22a",
			Name:                           "API_Msdeux",
			ApiVersion:                     "1.3",
			Description:                    "second API, made by maiwenn",
			DeployedAt:                     "2024-04-25T11:06:54.748Z",
			CreatedAt:                      "2024-04-25T11:06:49.799Z",
			UpdatedAt:                      "2024-04-25T11:06:56.48Z",
			DefinitionVersion:              "V2",
			EnvironmentId:                  envId,
			ExecutionMode:                  "V4_EMULATION_ENGINE",
			Visibility:                     "PRIVATE",
			LifecycleState:                 "CREATED",
			ContextPath:                    "/api_msdeux",
			DefinitionContext_origin:       "MANAGEMENT",
			DefinitionContext_mode:         "FULLY_MANAGED",
			Lnks_pictureUrl:                "http://sl2csoapp1490.pcloud.axway.int:8083/management/v2/environments/DEFAULT/apis/285cde3d-4340-44d2-9cde-3d4340e4d22a/picture?hash=1714043216480",
			Links_backgroundUrl:            "http://sl2csoapp1490.pcloud.axway.int:8083/management/v2/environments/DEFAULT/apis/285cde3d-4340-44d2-9cde-3d4340e4d22a/background?hash=1714043216480",
			DisableMembershipNotifications: false,
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

func (m mockAPIClient) IsReady() bool { return false }

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

func (m mockApiCache) AddApiToCache(name string, modDate time.Time, specHash string) {
}

func (m mockApiCache) HasProductChanged(name string, modDate time.Time, specHash string) bool {
	return true
}

func (m mockApiCache) GetProductWithName(name string) (*ApiCacheItem, error) {
	return nil, nil
}
