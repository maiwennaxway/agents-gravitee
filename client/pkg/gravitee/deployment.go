package gravitee

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/maiwennaxway/agents-gravitee/client/pkg/gravitee/models"
)

// GetDeployments - get a deployments for a proxy
func (a *config.graviteeClient) GetDeployments(proxyName string) (*models.DeploymentDetails, error) {
	response, err := a.newRequest(http.MethodGet, fmt.Sprintf("%s/apis/%s/deployments", a.orgURL, proxyName),
		WithDefaultHeaders(),
	).Execute()

	if err != nil {
		return nil, err
	}

	details := &models.DeploymentDetails{}
	json.Unmarshal(response.Body, details)
	if err != nil {
		return nil, err
	}

	return details, nil
}
