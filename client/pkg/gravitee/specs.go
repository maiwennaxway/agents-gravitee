package gravitee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetSpecFile - downloads the specfile from gravitee given the path of its location
func (a *GraviteeClient) GetSpecFile(specPath string) ([]byte, error) {
	// Get the spec file
	response, err := a.newRequest(http.MethodGet, fmt.Sprintf("%s%s", a.orgURL, specPath),
		//WithDefaultHeaders(),
		WithToken(a.GetConfig().Auth.GetToken()),
	).Execute()

	if err != nil {
		return nil, err
	}

	return response.Body, nil
}

// GetAllSpecs - downloads the specfile from gravitee given the path of its location ()
func (a *GraviteeClient) GetAllSpecs() ([]SpecDetails, error) {
	// Get the spec file
	response, err := a.newRequest(http.MethodGet, fmt.Sprintf("%s/environments/%s/apis", a.GetConfig().Auth.GetURL(), a.GetConfig().GetEnv()),
		//WithDefaultHeaders(),
		WithToken(a.GetConfig().Auth.GetToken()),
		//WithHeader("Authorization", fmt.Sprintf("Bearer %s", a.GetConfig().Auth.GetToken())),
	).Execute()

	if err != nil {
		return nil, err
	}

	details := SpecDetails{}
	err = json.Unmarshal(response.Body, &details)
	if err != nil {
		return nil, err
	}

	return details.Contents, nil
}
