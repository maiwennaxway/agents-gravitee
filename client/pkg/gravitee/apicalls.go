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
	req, err := a.newRequest(http.MethodGet, fmt.Sprintf("%s/v2/organizations/%s/environments/%s/apis", a.GetConfig().GetURL(), a.OrgId, a.GetConfig().GetEnv()),
		WithHeader("Content-Type", "application/json"),
		WithToken(a.GetConfig().Auth.GetToken()),
	).Execute()

	if err != nil {
		return nil, err
	}

	var Apis AllApis
	err = json.Unmarshal(req.Body, &Apis)
	if err != nil {
		return nil, err
	}

	if Apis.Pagination.TotalCount >= int64(Apis.Pagination.PerPage) {
		req, err := a.newRequest(http.MethodGet, fmt.Sprintf("%s/v2/organizations/%s/environments/%s/apis?perPage=%d", a.GetConfig().GetURL(), a.OrgId, a.GetConfig().GetEnv(), Apis.Pagination.TotalCount),
			WithHeader("Content-Type", "application/json"),
			WithToken(a.GetConfig().Auth.GetToken()),
		).Execute()

		if err != nil {
			return nil, err
		}

		var Apis AllApis
		err = json.Unmarshal(req.Body, &Apis)
		if err != nil {
			return nil, err
		}
		return Apis.Apis, nil
	}
	return Apis.Apis, nil
}

// GetApi - get details of the api
func (a *GraviteeClient) GetApi(apiID string) (api *models.Api, error error) {
	req, err := a.newRequest(http.MethodGet, fmt.Sprintf("%s/v2/organizations/%s/environments/%s/apis/%s", a.cfg.URL, a.OrgId, a.EnvId, apiID),
		WithHeader("Content-Type", "application/json"),
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

func (a *GraviteeClient) GetSpecs(apiID string) (specs []models.Spec, error error) {
	req, err := a.newRequest(http.MethodGet, fmt.Sprintf("%s/v2/organizations/%s/environments/%s/apis/%s/pages", a.cfg.URL, a.OrgId, a.EnvId, apiID),
		WithHeader("Content-Type", "application/json"),
		WithToken(a.GetConfig().Auth.GetToken()),
	).Execute()

	if err != nil {
		return nil, err
	}

	if req.Code != http.StatusOK {
		return nil, fmt.Errorf("received an unexpected response code %d from Gravitee when retrieving the app", req.Code)
	}
	var response AllSpecs
	err = json.Unmarshal(req.Body, &response)
	if err != nil {
		return nil, err
	}
	return response.Specs, nil
}

// GetApps - Get all Applications
func (a *GraviteeClient) GetApps() ([]models.App, error) {
	req, err := a.newRequest(http.MethodGet, fmt.Sprintf("%s/organizations/%s/environments/%s/applications", a.cfg.URL, a.OrgId, a.EnvId),
		WithHeader("Content-Type", "application/json"),
		WithToken(a.GetConfig().Auth.GetToken()),
	).Execute()

	if err != nil {
		return nil, err
	}

	if req.Code != http.StatusOK {
		return nil, fmt.Errorf("received an unexpected response code %d from Gravitee when retrieving the app", req.Code)
	}

	applis := []models.App{}
	err = json.Unmarshal(req.Body, &applis)
	if err != nil {
		return nil, err
	}

	return applis, nil
}

// GetApp - Get an Application by id
func (a *GraviteeClient) GetApp(id string) (app *models.App, err error) {
	req, err := a.newRequest(http.MethodGet, fmt.Sprintf("%s/organizations/%s/environments/%s/applications/%s", a.cfg.URL, a.OrgId, a.EnvId, id),
		WithHeader("Content-Type", "application/json"),
		WithToken(a.GetConfig().Auth.GetToken()),
	).Execute()

	if err != nil {
		return nil, err
	}

	if req.Code != http.StatusOK {
		return nil, fmt.Errorf("received an unexpected response code %d from Gravitee when retrieving the app", req.Code)
	}

	application := models.App{}
	err = json.Unmarshal(req.Body, &application)
	if err != nil {
		return nil, err
	}

	return &application, nil
}

func (a *GraviteeClient) DeployApi(apiID string) error {
	body, _ := json.Marshal("deploymentLabel: 'deploy'")
	req, err := a.newRequest(http.MethodPost, fmt.Sprintf("%s/v2/organizations/%s/environments/%s/apis/%s/deployments", a.cfg.URL, a.OrgId, a.EnvId, apiID),
		WithHeader("Content-Type", "application/json"),
		WithToken(a.GetConfig().Auth.GetToken()),
		WithBody(body),
	).Execute()

	if err != nil {
		return err
	}

	if req.Code != http.StatusOK {
		return fmt.Errorf("received an unexpected response code %d from Gravitee when retrieving the app", req.Code)
	}

	return nil
}
