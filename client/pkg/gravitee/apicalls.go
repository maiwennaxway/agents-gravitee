package gravitee

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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
func (a *GraviteeClient) GetApis() ([]string, error) {
	req, err := a.newRequest(http.MethodGet, fmt.Sprintf("%s/environments/%s/apis", a.GetConfig().Auth.GetURL(), a.GetConfig().GetEnv()),
		//WithDefaultHeaders(),
		WithHeader("Content-Type", "application/json"),
		//WithHeader("Accept", "application/json"),
		WithToken(a.GetConfig().Auth.GetToken()),
	).Execute()

	if err != nil {
		return nil, err
	}

	var data []map[string]interface{}

	// Décodez les données JSON dans la structure
	err = json.Unmarshal(req.Body, &data)
	if err != nil {
		return nil, fmt.Errorf("erreur lors du décodage des données JSON : %v", err)
	}

	// Initialisez une variable pour stocker les lignes formatées
	var formattedData []string

	// Ajoutez l'en-tête du tableau
	header := "| API Name | Environment | Execution Mode | Context Path | Target | State |"
	dashes := strings.Repeat("-", len(header))
	formattedData = append(formattedData, header, dashes)

	// Parcourez les données et ajoutez chaque ligne formatée
	for _, item := range data {
		apiName := item["name"].(string)
		environment := item["environmentId"].(string)
		executionMode := item["executionMode"].(string)
		contextPath := item["contextPath"].(string)
		target := item["proxy"].(map[string]interface{})["groups"].([]interface{})[0].(map[string]interface{})["endpoints"].([]interface{})[0].(map[string]interface{})["target"].(string)
		state := item["state"].(string)

		line := fmt.Sprintf("| %-8s | %-11s | %-14s | %-12s | %-6s | %-5s |", apiName, environment, executionMode, contextPath, target, state)
		formattedData = append(formattedData, line)
	}

	return formattedData, nil
	/*err = json.Unmarshal(req.Body, &response.Apis)
	if err != nil {
		return "", err
	}*/
	//return data, nil
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
