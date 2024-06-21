package gravitee

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/Axway/agent-sdk/pkg/api"
	"github.com/maiwennaxway/agents-gravitee/client/pkg/gravitee/models"
	"github.com/stretchr/testify/assert"
)

/*func TestGetEnvironments(t *testing.T) {
	cases := map[string]struct {
		responses    []api.MockResponse
		expectedEnvs int
	}{
		"environments returned": {
			responses: []api.MockResponse{
				{
					RespData: `["env1","env2"]`,
					RespCode: http.StatusOK,
				},
			},
			expectedEnvs: 2,
		},
		"error getting environments": {
			responses: []api.MockResponse{
				{
					ErrString: "error",
				},
			},
			expectedEnvs: 0,
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			c := createTestClient(t, &api.MockHTTPClient{Responses: tc.responses})

			data := c.GetEnvironments()
			assert.Len(t, data, tc.expectedEnvs)
		})
	}
}*/

func TestGetApis(t *testing.T) {
	cases := map[string]struct {
		responses    []api.MockResponse
		expectedEnvs int
		expectErr    bool
	}{
		"apis returned": {
			responses: []api.MockResponse{
				{
					RespData: `["api1","api2"]`,
					RespCode: http.StatusOK,
				},
			},
			expectedEnvs: 2,
		},
		"error getting apis": {
			responses: []api.MockResponse{
				{
					ErrString: "error",
				},
			},
			expectErr: true,
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			c := createTestClient(t, &api.MockHTTPClient{Responses: tc.responses})

			data, err := c.GetApis()
			if tc.expectErr {
				assert.NotNil(t, err)
			}
			assert.Len(t, data, tc.expectedEnvs)
		})
	}
}

func TestGetApi(t *testing.T) {
	expectedApi := models.Api{
		Name: "API_Ms",
	}
	expectedAppData, _ := json.Marshal(expectedApi)
	cases := map[string]struct {
		responses []api.MockResponse
		apiId     string
		expectErr bool
	}{
		"error making http call": {
			responses: []api.MockResponse{
				{
					ErrString: "error",
				},
			},
			expectErr: true,
		},
		"error, unexpected response code": {
			responses: []api.MockResponse{
				{
					RespCode: http.StatusAccepted,
				},
			},
			expectErr: true,
		},
		"error, data returned not a api": {
			responses: []api.MockResponse{
				{
					RespCode: http.StatusOK,
					RespData: `"data":"aaaa"`,
				},
			},
			expectErr: true,
		},
		"success getting api": {
			responses: []api.MockResponse{
				{
					RespCode: http.StatusOK,
					RespData: string(expectedAppData),
				},
			},
			apiId: expectedApi.Name,
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			c := createTestClient(t, &api.MockHTTPClient{Responses: tc.responses})

			apiOut, err := c.GetApi(tc.apiId)
			if tc.expectErr {
				assert.NotNil(t, err)
				return
			}
			assert.Nil(t, err)
			assert.NotNil(t, apiOut)
			assert.Equal(t, expectedApi.Name, apiOut.Name)
		})
	}
}
