package config

import (
	"errors"
	"strings"
	"time"

	coreapi "github.com/Axway/agent-sdk/pkg/api"
	"github.com/Axway/agent-sdk/pkg/cmd/properties"
	corecfg "github.com/Axway/agent-sdk/pkg/config"
)

// graviteeConfig - represents the config for gateway
type graviteeConfig struct {
	corecfg.IConfigValidator
	Organization    string             `config:"organization"`
	URL             string             `config:"url"`
	DataURL         string             `config:"dataURL"`
	APIVersion      string             `config:"apiVersion"`
	Filter          string             `config:"filter"`
	DeveloperID     string             `config:"developerID"`
	Auth            *AuthConfig        `config:"auth"`
	Intervals       *graviteeIntervals `config:"interval"`
	Workers         *graviteeWorkers   `config:"workers"`
	CloneAttributes bool               `config:"cloneAttributes"`
	AllTraffic      bool               `config:"allTraffic"`
	NotSetTraffic   bool               `config:"notSetTraffic"`
	mode            discoveryMode
}

// graviteeClient - Represents the Gateway client
type graviteeClient struct {
	cfg         *graviteeConfig
	apiClient   coreapi.Client
	accessToken string
	developerID string
	envToURLs   map[string][]string
	isReady     bool
	orgURL      string
	dataURL     string
}

// graviteeIntervals - intervals for the gravitee agent to use
type graviteeIntervals struct {
	Proxy   time.Duration `config:"proxy"`
	Spec    time.Duration `config:"spec"`
	Product time.Duration `config:"product"`
	Stats   time.Duration `config:"stats"`
}

// graviteeWorkers - number of workers for the gravitee agent to use
type graviteeWorkers struct {
	Proxy   int `config:"proxy"`
	Spec    int `config:"spec"`
	Product int `config:"product"`
}

type discoveryMode int

const (
	discoveryModeProxy = iota + 1
	discoveryModeProduct
)

const (
	discoveryModeProxyString   = "proxy"
	discoveryModeProductString = "product"
)

func (m discoveryMode) String() string {
	return map[discoveryMode]string{
		discoveryModeProxy:   discoveryModeProductString,
		discoveryModeProduct: discoveryModeProductString,
	}[m]
}

func stringToDiscoveryMode(s string) discoveryMode {
	if mode, ok := map[string]discoveryMode{
		discoveryModeProxyString:   discoveryModeProxy,
		discoveryModeProductString: discoveryModeProduct,
	}[strings.ToLower(s)]; ok {
		return mode
	}
	return 0
}

const (
	pathURL                = "gravitee.url"
	pathDataURL            = "gravitee.dataURL"
	pathAPIVersion         = "gravitee.apiVersion"
	pathOrganization       = "gravitee.organization"
	pathMode               = "gravitee.discoveryMode"
	pathFilter             = "gravitee.filter"
	pathCloneAttributes    = "gravitee.cloneAttributes"
	pathAllTraffic         = "gravitee.allTraffic"
	pathNotSetTraffic      = "gravitee.notSetTraffic"
	pathAuthURL            = "gravitee.auth.url"
	pathAuthServerUsername = "gravitee.auth.serverUsername"
	pathAuthServerPassword = "gravitee.auth.serverPassword"
	pathAuthUsername       = "gravitee.auth.username"
	pathAuthPassword       = "gravitee.auth.password"
	pathSpecInterval       = "gravitee.interval.spec"
	pathProxyInterval      = "gravitee.interval.proxy"
	pathProductInterval    = "gravitee.interval.product"
	pathStatsInterval      = "gravitee.interval.stats"
	pathDeveloper          = "gravitee.developerID"
	pathSpecWorkers        = "gravitee.workers.spec"
	pathProxyWorkers       = "gravitee.workers.proxy"
	pathProductWorkers     = "gravitee.workers.product"
)

// AddProperties - adds config needed for gravitee client
func AddProperties(rootProps properties.Properties) {
	rootProps.AddStringProperty(pathMode, "proxy", "gravitee Organization")
	rootProps.AddStringProperty(pathOrganization, "", "gravitee Organization")
	rootProps.AddStringProperty(pathURL, "https://api.enterprise.gravitee.com", "gravitee Base URL")
	rootProps.AddStringProperty(pathAPIVersion, "v1", "gravitee API Version")
	rootProps.AddStringProperty(pathFilter, "", "Filter used on discovering gravitee products")
	rootProps.AddStringProperty(pathDataURL, "https://gravitee.com/dapi/api", "gravitee Data API URL")
	rootProps.AddStringProperty(pathAuthURL, "https://login.gravitee.com", "URL to use when authenticating to gravitee")
	rootProps.AddStringProperty(pathAuthServerUsername, "edgecli", "Username to use to when requesting gravitee token")
	rootProps.AddStringProperty(pathAuthServerPassword, "edgeclisecret", "Password to use to when requesting gravitee token")
	rootProps.AddBoolProperty(pathCloneAttributes, false, "Set to true to copy the tags when provisioning a Product in product mode")
	rootProps.AddBoolProperty(pathAllTraffic, false, "Set to true to report metrics for all traffic for the selected mode")
	rootProps.AddBoolProperty(pathNotSetTraffic, false, "Set to true to report metrics for values reported with (not set) ast the name")
	rootProps.AddStringProperty(pathAuthUsername, "", "Username to use to authenticate to gravitee")
	rootProps.AddStringProperty(pathAuthPassword, "", "Password for the user to authenticate to gravitee")
	rootProps.AddDurationProperty(pathSpecInterval, 30*time.Minute, "The time interval between checking for updated specs", properties.WithLowerLimit(1*time.Minute))
	rootProps.AddDurationProperty(pathProxyInterval, 30*time.Second, "The time interval between checking for updated proxies", properties.WithUpperLimit(5*time.Minute))
	rootProps.AddDurationProperty(pathProductInterval, 30*time.Second, "The time interval between checking for updated products", properties.WithUpperLimit(5*time.Minute))
	rootProps.AddDurationProperty(pathStatsInterval, 5*time.Minute, "The time interval between checking for updated stats", properties.WithLowerLimit(1*time.Minute), properties.WithUpperLimit(15*time.Minute))
	rootProps.AddStringProperty(pathDeveloper, "", "Developer ID used to create applications")
	rootProps.AddIntProperty(pathProxyWorkers, 10, "Max number of workers discovering proxies")
	rootProps.AddIntProperty(pathSpecWorkers, 20, "Max number of workers discovering specs")
	rootProps.AddIntProperty(pathProductWorkers, 10, "Max number of workers discovering products")
}

// ParseConfig - parse the config on startup
func ParseConfig(rootProps properties.Properties) *graviteeConfig {
	return &graviteeConfig{
		Organization:    rootProps.StringPropertyValue(pathOrganization),
		URL:             strings.TrimSuffix(rootProps.StringPropertyValue(pathURL), "/"),
		APIVersion:      rootProps.StringPropertyValue(pathAPIVersion),
		DataURL:         strings.TrimSuffix(rootProps.StringPropertyValue(pathDataURL), "/"),
		DeveloperID:     rootProps.StringPropertyValue(pathDeveloper),
		mode:            stringToDiscoveryMode(rootProps.StringPropertyValue(pathMode)),
		Filter:          rootProps.StringPropertyValue(pathFilter),
		CloneAttributes: rootProps.BoolPropertyValue(pathCloneAttributes),
		AllTraffic:      rootProps.BoolPropertyValue(pathAllTraffic),
		NotSetTraffic:   rootProps.BoolPropertyValue(pathNotSetTraffic),
		Intervals: &graviteeIntervals{
			Stats:   rootProps.DurationPropertyValue(pathStatsInterval),
			Proxy:   rootProps.DurationPropertyValue(pathProxyInterval),
			Spec:    rootProps.DurationPropertyValue(pathSpecInterval),
			Product: rootProps.DurationPropertyValue(pathProductInterval),
		},
		Workers: &graviteeWorkers{
			Proxy:   rootProps.IntPropertyValue(pathProxyWorkers),
			Spec:    rootProps.IntPropertyValue(pathSpecWorkers),
			Product: rootProps.IntPropertyValue(pathProductWorkers),
		},
		Auth: &AuthConfig{
			Username:       rootProps.StringPropertyValue(pathAuthUsername),
			Password:       rootProps.StringPropertyValue(pathAuthPassword),
			ServerUsername: rootProps.StringPropertyValue(pathAuthServerUsername),
			ServerPassword: rootProps.StringPropertyValue(pathAuthServerPassword),
			URL:            rootProps.StringPropertyValue(pathAuthURL),
		},
	}
}

// ValidateCfg - Validates the gateway config
func (a *graviteeConfig) ValidateCfg() (err error) {
	if a.mode == 0 {
		return errors.New("invalid gravitee configuration: discoveryMode must be proxy or product")
	}

	if a.URL == "" {
		return errors.New("invalid gravitee configuration: url is not configured")
	}

	if a.APIVersion == "" {
		return errors.New("invalid gravitee configuration: api version is not configured")
	}

	if a.DataURL == "" {
		return errors.New("invalid gravitee configuration: data url is not configured")
	}

	if a.Auth.Username == "" {
		return errors.New("invalid gravitee configuration: username is not configured")
	}

	if a.Auth.Password == "" {
		return errors.New("invalid gravitee configuration: password is not configured")
	}

	if a.DeveloperID == "" {
		return errors.New("invalid gravitee configuration: developer ID must be configured")
	}

	if a.Workers.Proxy < 1 {
		return errors.New("invalid gravitee configuration: proxy workers must be greater than 0")
	}

	if a.Workers.Spec < 1 {
		return errors.New("invalid gravitee configuration: spec workers must be greater than 0")
	}

	return
}

// GetAuth - Returns the Auth Config
func (a *graviteeConfig) GetAuth() *AuthConfig {
	return a.Auth
}

// GetIntervals - Returns the Intervals
func (a *graviteeConfig) GetIntervals() *graviteeIntervals {
	return a.Intervals
}

// GetWorkers - Returns the number of Workers
func (a *graviteeConfig) GetWorkers() *graviteeWorkers {
	return a.Workers
}

func (a *graviteeConfig) IsProxyMode() bool {
	return a.mode == discoveryModeProxy
}

func (a *graviteeConfig) IsProductMode() bool {
	return a.mode == discoveryModeProduct
}

func (a *graviteeConfig) ShouldCloneAttributes() bool {
	return a.CloneAttributes
}

func (a *graviteeConfig) ShouldReportAllTraffic() bool {
	return a.AllTraffic
}

func (a *graviteeConfig) ShouldReportNotSetTraffic() bool {
	return a.NotSetTraffic
}
