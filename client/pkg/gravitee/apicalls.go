package gravitee

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/maiwennaxway/agents-gravitee/client/pkg/gravitee/models"
)

const (
	developerAppsURL = "%s/developers/%s/apps/%s"
)

// GetEnvironments - get the list of environments for the org
func (a *GraviteeClient) GetEnvironments() []string {
	// Get the developers
	response, err := a.newRequest(http.MethodGet, fmt.Sprintf("%s/environments", a.cfg.Auth.URL),
		WithDefaultHeaders(),
	).Execute()

	environments := []string{}
	if err == nil {
		json.Unmarshal(response.Body, &environments)
	}

	return environments
}

func (a *GraviteeClient) GetApis() (Apis, error) {
	response, err := a.newRequest(http.MethodGet, fmt.Sprintf("%s/apis", a.orgURL),
		WithDefaultHeaders(),
	).Execute()
	if err != nil {
		return nil, err
	}

	apis:= Apis{}
	if err != nil {
		json.Unmarshal(response.Body, &apis)
	}

	return apis, nil
}

// GetProduct - get details of the product
func (a *GraviteeClient) GetApi(apiID string) (models.Api.Id, error) {
	// Get the product
	response, err := a.newRequest(http.MethodGet, fmt.Sprintf("%s/apis/%s", a.orgURL, apiID),
		WithDefaultHeaders(),
	).Execute()
	if err != nil {
		return nil, err
	}

	if response.Code != http.StatusOK {
		return nil, fmt.Errorf("received an unexpected response code %d from gravitee when retrieving the app", response.Code)
	}

	api := &models.ApiProduct{}
	json.Unmarshal(response.Body, api)

	return api, nil
}

// GetRevisionSpec - gets the resource file of type openapi for the org, api, revision, and spec file specified
func (a *GraviteeClient) GetRevisionSpec(apiName, revisionNumber, specFile string) []byte {
	// Get the openapi resource file
	response, err := a.newRequest(http.MethodGet, fmt.Sprintf("%s/apis/%s/revisions/%s/resourcefiles/openapi/%s", a.orgURL, apiName, revisionNumber, specFile),
		WithDefaultHeaders(),
	).Execute()

	if err != nil {
		return []byte{}
	}

	return response.Body
}

// GetStats - get the api stats for a specific environment
func (a *GraviteeClient) GetStats(env, dimension, metricSelect string, start, end time.Time) (*models.Metrics, error) {
	// Get the spec content file
	const format = "01/02/2006 15:04"

	response, err := a.newRequest(http.MethodGet, fmt.Sprintf("%s/environments/%v/stats/%s", a.orgURL, env, dimension),
		WithQueryParams(map[string]string{
			"select":    metricSelect,
			"timeUnit":  "minute",
			"timeRange": fmt.Sprintf("%s~%s", time.Time.UTC(start).Format(format), time.Time.UTC(end).Format(format)),
			"sortby":    "sum(message_count)",
			"sort":      "ASC",
		}),
		WithDefaultHeaders(),
	).Execute()

	if err != nil {
		return nil, err
	}

	stats := &models.Metrics{}
	err = json.Unmarshal(response.Body, stats)
	if err != nil {
		return nil, err
	}

	return stats, nil
}

func (a *GraviteeClient) CreateAPIProduct(product *models.ApiProduct) (*models.ApiProduct, error) {
	// create a new developer app
	data, err := json.Marshal(product)
	if err != nil {
		return nil, err
	}

	u := fmt.Sprintf("%s/apiproducts", a.orgURL)
	response, err := a.newRequest(http.MethodPost, u,
		WithDefaultHeaders(),
		WithBody(data),
	).Execute()

	if err != nil {
		return nil, err
	}

	if response.Code != http.StatusCreated {
		return nil, fmt.Errorf("received an unexpected response code %d from gravitee when creating the api product", response.Code)
	}

	newProduct := models.ApiProduct{}
	err = json.Unmarshal(response.Body, &newProduct)
	if err != nil {
		return nil, err
	}

	return &newProduct, err

}
