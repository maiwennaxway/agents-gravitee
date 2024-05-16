package gravitee

import (
	"encoding/json"
	"fmt"
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
func (a *GraviteeClient) GetApis() ([]models.Api, error) {
	req, err := a.newRequest(http.MethodGet, fmt.Sprintf("%s/environments/%s/apis", a.GetConfig().Auth.GetURL(), a.GetConfig().GetEnv()),
		//WithDefaultHeaders(),
		WithHeader("Content-Type", "application/json"),
		//WithHeader("Accept", "application/json"),
		WithToken(a.GetConfig().Auth.GetToken()),
	).Execute()

	if err != nil {
		return nil, err
	}

	if req.Code != http.StatusOK {
		return nil, fmt.Errorf("received an unexpected response code %d from Gravitee when retrieving the app", req.Code)
	}

	var response AllApis
	err = json.Unmarshal(req.Body, &response.Apis)
	if err != nil {
		return nil, err
	}
	return response.Apis, nil

}

// GetApi - get details of the api
func (a *GraviteeClient) GetApi(apiID string, envID string) (api *models.Api, error error) {
	req, err := a.newRequest(http.MethodGet, fmt.Sprintf("%s/environments/%s/apis/%s", a.cfg.Auth.URL, envID, apiID),
		//WithDefaultHeaders(),
		WithHeader("Content-Type", "application/json"),
		//WithHeader("Accept", "application/json"),
		WithToken(a.GetConfig().Auth.GetToken()),
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
