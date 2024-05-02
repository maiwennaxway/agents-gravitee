package gravitee

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/maiwennaxway/agents-gravitee/client/pkg/gravitee/models"
)

// GetEnvironments - get the list of environments for the org
/*func (a *GraviteeClient) GetEnvironments() []string {
	// Get the environments
	response, err := a.newRequest(http.MethodGet, fmt.Sprintf("%s/environments", a.cfg.Auth.URL),
		WithDefaultHeaders(),
	).Execute()

	environments := []string{}
	if err == nil {
		json.Unmarshal(response.Body, &environments)
	}

	return environments
}*/

// GetListAPIs - get the list of APIs
func (a *GraviteeClient) GetApis() (apis Apis, error error) {
	//
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/environments/%s/apis", a.cfg.Auth.URL, a.cfg.EnvName), nil)

	req.Header.Add("Accept", "application/json")
	//req.Header.Add("Authorization", "Basic YWRtaW46YWRtaW4=")
	req.Header.Add("Authorization", "Bearer 8f734df7-a350-44c3-b34d-f7a350c4c37a")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

	apis = Apis{}
	json.Unmarshal(body, &apis)

	return apis, nil

}

// GetApi - get details of the api
func (a *GraviteeClient) GetApi(apiID string, envID string) (api *models.Api, error error) {
	req, err := a.newRequest(http.MethodGet, fmt.Sprintf("%s/environments/%s/api/%s", a.cfg.Auth.URL, envID, apiID),
		WithDefaultHeaders(),
		WithHeader("Content-Type", "application/json"),
		WithHeader("Accept", "application/json"),
		WithToken("8f734df7-a350-44c3-b34d-f7a350c4c37a"),
	).Execute()

	if err != nil {
		return nil, err
	}

	if req.Code != http.StatusOK {
		return nil, fmt.Errorf("received an unexpected response code %d from Gravitee when retrieving the app", req.Code)
	}

	apitry := models.Api{}
	err = json.Unmarshal(req.Body, &apitry)
	if err != nil {
		return nil, err
	}
	return &apitry, nil
}
