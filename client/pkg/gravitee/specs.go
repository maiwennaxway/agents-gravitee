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
		WithDefaultHeaders(),
		WithHeader("Authorization", "Bearer 8f734df7-a350-44c3-b34d-f7a350c4c37a"),
	).Execute()

	if err != nil {
		return nil, err
	}

	return response.Body, nil
}

// GetSpecFromURL - downloads the specfile from a URL outside of gravitee
func (a *GraviteeClient) GetSpecFromURL(url string, options ...RequestOption) ([]byte, error) {
	// Get the spec file
	response, err := a.newRequest(http.MethodGet, url, options...).Execute()

	if err != nil {
		return nil, err
	}

	return response.Body, nil
}

// GetAllSpecs - downloads the specfile from gravitee given the path of its location ()
func (a *GraviteeClient) GetAllSpecs() ([]SpecDetails, error) {
	// Get the spec file
	response, err := a.newRequest(http.MethodGet, fmt.Sprintf("%s/environments/%s/apis", a.orgURL, a.EnvId),
		WithDefaultHeaders(),
		WithHeader("Authorization", "Bearer 8f734df7-a350-44c3-b34d-f7a350c4c37a"),
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
