package gravitee

import (
	"fmt"
	"testing"

	"github.com/maiwennaxway/agents-gravitee/client/pkg/config"
	"github.com/maiwennaxway/agents-gravitee/client/pkg/gravitee"
	"github.com/stretchr/testify/assert"
)

func Test_pollAPIsJob(t *testing.T) {
	tests := []struct {
		name         string
		ApiID        string
		config       *config.GraviteeConfig
		allApiErr    bool
		getApiErr    bool
		specNotFound bool
		filterFailed bool
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

			readyFunc := func() bool {
				return true
			}

			ApiJob := newPollAPIsJob(client, readyFunc, 10)
			assert.False(t, ApiJob.FirstRunComplete())

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

func (m mockApiClient) GetApi(apiId string) (api gravitee.Api, err error) {
	apis = map[string]*gravitee.APIs{
		"RTE": {ApiResources: []string{},
			ApprovalType: "auto",
			Attributes: []gravitee.Attribute{
				{
					Name:  "access",
					Value: "public",
				},
			},
			CreatedAt:      1665416157626,
			CreatedBy:      "cicd_technical_user@engie.com",
			Description:    "Generated Api",
			DisplayName:    "RTE",
			Environments:   []string{"acc", "int", "itt", "ppd"},
			LastModifiedAt: 1665758109625,
			LastModifiedBy: "cicd_technical_user@engie.com",
			Name:           "RTE",
			Proxies:        []string{"public-apiset-protected-ecowatt-v10"},
			Quota:          "10000",
			QuotaInterval:  "1",
			QuotaTimeUnit:  "minute",
			Scopes:         []string{"apihour:read", "apihour:write"},
		},
		"cell": {ApiResources: []string{"/"},
			ApprovalType: "auto",
			Attributes: []gravitee.Attribute{
				{
					Name:  "access",
					Value: "public",
				},
			},
			CreatedAt:      1632752367332,
			CreatedBy:      "cicd_technical_user@engie.com",
			Description:    "Generated Api",
			DisplayName:    "Cell",
			Environments:   []string{"acc", "int", "itt", "ppd"},
			LastModifiedAt: 1665758109625,
			LastModifiedBy: "cicd_technical_user@engie.com",
			Name:           "cell",
			Quota:          "10000",
			QuotaInterval:  "1",
			QuotaTimeUnit:  "minute",
			Scopes:         []string{"apihour:read", "apihour:write"},
		},
		"priv-PushNotif": {ApiResources: []string{"/"},
			ApprovalType: "auto",
			Attributes: []gravitee.Attribute{
				{
					Name:  "access",
					Value: "public",
				},
			},
			CreatedAt:      1632752359124,
			CreatedBy:      "cicd_technical_user@engie.com",
			Description:    "Generated Api",
			DisplayName:    "Private-PushNotif",
			Environments:   []string{"acc", "int", "itt", "ppd"},
			LastModifiedAt: 1665758129808,
			LastModifiedBy: "cicd_technical_user@engie.com",
			Name:           "priv-PushNotif",
			Quota:          "10000",
			QuotaInterval:  "1",
			QuotaTimeUnit:  "minute",
			Scopes:         []string{"apihour:read", "apihour:write"},
		},
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
