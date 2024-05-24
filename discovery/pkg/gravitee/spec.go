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

/*func (a *agentSpec) GetSpecs(apiID string, specPath string) (*specItem, error) {
	data, err := a.
	return &specItem, nil
}*/

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

func (a *agentSpec) GetSpecWithName(id string) (*specItem, error) {
	data, err := a.cache.GetBySecondaryKey(strings.ToLower(id))
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, fmt.Errorf("spec with name %s not found in cache", id)
	}

	specItem := data.(specItem)
	return &specItem, nil
}

func (a *agentSpec) HasSpecChanged(name string, modDate time.Time) bool {
	data, err := a.cache.GetBySecondaryKey(name)
	if err != nil || data == nil {
		// spec not in cache
		return true
	}

	specItem := data.(specItem)
	return modDate.After(specItem.ModDate)
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

func apiPrimaryKey(name string) string {
	return fmt.Sprintf("api-%s", name)
}

func (a *agentSpec) AddApiToCache(name string, modDate time.Time, specHash string) {
	item := ApiCacheItem{
		Name:     strings.ToLower(name),
		ModDate:  modDate,
		SpecHash: specHash,
	}

	a.cache.Set(apiPrimaryKey(name), item)
}
