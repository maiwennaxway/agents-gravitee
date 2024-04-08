package gravitee

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
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
func (a *GraviteeClient) GetApis() {
	//
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s:8083/management/v2/environments/%s/apis", a.cfg.Auth.URL, a.cfg.EnvName), nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Basic YWRtaW46YWRtaW4=")
}

type API string

// GetApi - get details of the api
func (a *GraviteeClient) GetApibyApiId(apiID string) (API, error) {
	payload := strings.NewReader("{\n  \"query\": \"my api\",\n  \"ids\": [\n    \"apiId-1\",\n    \"apiId-2\"\n  ],\n  \"definitionVersion\": \"V4\"\n}")

	req, err := http.NewRequest("POST", fmt.Sprintf("%s:8083/management/v2/environments/%s/api/%s", a.cfg.Auth.URL, a.cfg.EnvName, apiID), payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Basic YWRtaW46YWRtaW4=")

	res, err := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

	return API(apiID), err
}
