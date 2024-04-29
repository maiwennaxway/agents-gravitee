package gravitee

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/maiwennaxway/agents-gravitee/client/pkg/gravitee/models"
)

// GetEnvironments - get the list of environments for the org
func (a *GraviteeClient) GetEnvironments() []string {
	// Get the environments
	response, err := a.newRequest(http.MethodGet, fmt.Sprintf("%s/environments", a.cfg.Auth.URL),
		WithDefaultHeaders(),
	).Execute()

	environments := []string{}
	if err == nil {
		json.Unmarshal(response.Body, &environments)
	}

	return environments
}

// GetListAPIs - get the list of APIs
func (a *GraviteeClient) GetApis() (apis Apis, error error) {
	//
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s:8083/management/v2/environments/%s/apis", a.cfg.Auth.URL, a.cfg.EnvName), nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Basic YWRtaW46YWRtaW4=")
	//req.Header.Add("Authorization", "Auth-Graviteeio-APIM=")

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
	payload := strings.NewReader("{\n  \"query\": \"my api\",\n  \"ids\": [\n    \"apiId-1\",\n    \"apiId-2\"\n  ],\n  \"definitionVersion\": \"V4\"\n}")

	req, _ := http.NewRequest("GET", fmt.Sprintf("%s:8083/management/v2/environments/%s/api/%s", a.cfg.Auth.URL, envID, apiID), payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Basic YWRtaW46YWRtaW4=")
	//req.Header.Add("Authorization", "Auth-Graviteeio-APIM=")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

	apitry := models.Api{}
	err = json.Unmarshal(body, &apitry)
	if err != nil {
		return nil, err
	}
	return &apitry, nil
}
