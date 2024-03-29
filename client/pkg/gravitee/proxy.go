package gravitee

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"github.com/maiwennaxway/agents-gravitee/client/pkg/gravitee/models"
)

// proxyXML
type proxyXML struct {
	XMLName             xml.Name             `xml:"ProxyEndpoint"`
	HTTPProxyConnection *HTTPProxyConnection `xml:"HTTPProxyConnection"`
}

// HTTPProxyConnection
type HTTPProxyConnection struct {
	XMLName     xml.Name `xml:"HTTPProxyConnection"`
	BasePath    string   `xml:"BasePath"`
	VirtualHost string   `xml:"VirtualHost"`
}

// Products
type Proxies []string

// GetAllProxies - get all proxies
func (a *GraviteeClient) GetAllProxies() (Proxies, error) {
	response, err := a.newRequest(http.MethodGet, fmt.Sprintf("%s/apis", a.orgURL),
		WithDefaultHeaders(),
	).Execute()
	if err != nil {
		return nil, err
	}

	proxies := Proxies{}
	err = json.Unmarshal(response.Body, &proxies)
	if err != nil {
		return nil, err
	}

	return proxies, nil
}

// GetProxy - get a proxy with a name
func (a *GraviteeClient) GetProxy(proxyName string) (*models.ApiProxy, error) {
	response, err := a.newRequest(http.MethodGet, fmt.Sprintf("%s/apis/%s", a.orgURL, proxyName),
		WithDefaultHeaders(),
	).Execute()
	if err != nil {
		return nil, err
	}

	proxy := &models.ApiProxy{}
	err = json.Unmarshal(response.Body, proxy)
	if err != nil {
		return nil, err
	}

	return proxy, nil
}

// GetRevision - get a revision of a proxy with a name
func (a *GraviteeClient) GetRevision(proxyName, revision string) (*models.ApiProxyRevision, error) {
	response, err := a.newRequest(http.MethodGet, fmt.Sprintf("%s/apis/%s/revisions/%s", a.orgURL, proxyName, revision),
		WithDefaultHeaders(),
	).Execute()
	if err != nil {
		return nil, err
	}

	proxyRevision := &models.ApiProxyRevision{}
	json.Unmarshal(response.Body, proxyRevision)
	if err != nil {
		return nil, err
	}

	return proxyRevision, nil
}

// GetRevisionConnectionType - get a revision bundle and open the proxy config file
func (a *GraviteeClient) GetRevisionConnectionType(proxyName, revision string) (*HTTPProxyConnection, error) {
	response, err := a.newRequest(http.MethodGet, fmt.Sprintf("%s/apis/%s/revisions/%s", a.orgURL, proxyName, revision),
		WithDefaultHeaders(),
		WithQueryParam("format", "bundle"),
	).Execute()
	if err != nil {
		return nil, err
	}

	// response is a zip file, lets open it and find the proxy config file
	zipReader, err := zip.NewReader(bytes.NewReader(response.Body), int64(len(response.Body)))
	if err != nil {
		return nil, err
	}

	// Read all the files from the zip archive
	var fileBytes []byte
	for _, zipFile := range zipReader.File {
		if zipFile.Name != "apiproxy/proxies/default.xml" {
			continue
		}
		fileBytes, err = readZipFile(zipFile)
		if err != nil {
			return nil, err
		}
		break
	}

	if len(fileBytes) == 0 {
		return nil, fmt.Errorf("could not find the proxy configuration file in the api revision bundle")
	}

	data := &proxyXML{}
	xml.Unmarshal(fileBytes, data)

	return data.HTTPProxyConnection, nil
}

func readZipFile(zf *zip.File) ([]byte, error) {
	f, err := zf.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return io.ReadAll(f)
}

// GetProxy - get a proxy with a name
func (a *GraviteeClient) GetRevisionResourceFile(proxyName, revision, resourceType, resourceName string) ([]byte, error) {
	response, err := a.newRequest(http.MethodGet, fmt.Sprintf("%s/apis/%s/revisions/%s/resourcefiles/%s/%s", a.orgURL, proxyName, revision, resourceType, resourceName),
		WithDefaultHeaders(),
	).Execute()
	if err != nil {
		return nil, err
	}

	return response.Body, nil
}

// GetRevisionPolicyByName - get the details about a named policy on a revision
func (a *GraviteeClient) GetRevisionPolicyByName(proxyName, revision, policyName string) (*PolicyDetail, error) {
	response, err := a.newRequest(http.MethodGet, fmt.Sprintf("%s/apis/%s/revisions/%s/policies/%s", a.orgURL, proxyName, revision, policyName),
		WithDefaultHeaders(),
	).Execute()
	if err != nil {
		return nil, err
	}

	policyDetails := &PolicyDetail{}
	json.Unmarshal(response.Body, policyDetails)
	if err != nil {
		return nil, err
	}

	return policyDetails, nil
}
