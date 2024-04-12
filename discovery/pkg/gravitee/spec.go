package gravitee

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/Axway/agent-sdk/pkg/cache"
)

type agentSpec struct {
	cache              cache.Cache
	specEndpointToKeys map[string][]specItem
	mutex              *sync.Mutex
}

type specItem struct {
	ID          string
	Name        string
	ContentPath string
	ModDate     time.Time
}

func newAgentSpec() *agentSpec {
	return &agentSpec{
		cache:              cache.New(),
		specEndpointToKeys: make(map[string][]specItem),
		mutex:              &sync.Mutex{},
	}
}

func specPrimaryKey(name string) string {
	return fmt.Sprintf("spec-%s", name)
}

func (a *agentSpec) AddSpecToCache(id, path, name string, modDate time.Time, endpoints ...string) {
	item := specItem{
		ID:          id,
		Name:        strings.ToLower(name),
		ContentPath: path,
		ModDate:     modDate,
	}

	a.cache.SetWithSecondaryKey(specPrimaryKey(name), path, item)
	a.cache.SetSecondaryKey(specPrimaryKey(name), strings.ToLower(name))
	a.cache.SetSecondaryKey(specPrimaryKey(name), id)
	a.mutex.Lock()
	defer a.mutex.Unlock()
	for _, ep := range endpoints {
		if _, found := a.specEndpointToKeys[ep]; !found {
			a.specEndpointToKeys[ep] = []specItem{}
		}
		a.specEndpointToKeys[ep] = append(a.specEndpointToKeys[ep], item)
	}
}

func (a *agentSpec) GetSpecWithPath(path string) (*specItem, error) {
	data, err := a.cache.GetBySecondaryKey(path)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, fmt.Errorf("spec path name %s not found in cache", path)
	}

	specItem := data.(specItem)
	return &specItem, nil
}

func (a *agentSpec) GetSpecWithName(name string) (*specItem, error) {
	data, err := a.cache.GetBySecondaryKey(strings.ToLower(name))
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, fmt.Errorf("spec with name %s not found in cache", name)
	}

	specItem := data.(specItem)
	return &specItem, nil
}

// GetSpecPathWithEndpoint - returns the lat modified spec found with this endpoint
func (a *agentSpec) GetSpecPathWithEndpoint(endpoint string) (string, error) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	items, found := a.specEndpointToKeys[endpoint]
	if !found {
		return "", fmt.Errorf("no spec found for endpoint: %s", endpoint)
	}

	latest := specItem{}
	for _, item := range items {
		if item.ModDate.After(latest.ModDate) {
			latest = item
		}
	}

	return latest.ContentPath, nil
}
