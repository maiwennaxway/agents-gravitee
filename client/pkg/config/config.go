package config

import (
	"errors"
	"strings"
	"time"

	"github.com/Axway/agent-sdk/pkg/cmd/properties"
)

type props interface {
	AddStringProperty(name string, defaultVal string, description string)
	AddIntProperty(name string, defaultVal int, description string)
	AddBoolProperty(name string, defaultVal bool, description string)
	AddDurationProperty(name string, defaultVal time.Duration, description string, opts ...properties.DurationOpt)
	StringPropertyValue(name string) string
	IntPropertyValue(name string) int
	BoolPropertyValue(name string) bool
	DurationPropertyValue(name string) time.Duration
}

// GraviteeWorkers - number of workers for the gravitee agent to use
type GraviteeSpecConfig struct {
	DisablePollForSpecs bool   `config:"disablePollForSpecs"`
	Unstructured        bool   `config:"unstructured"`
	MatchOnURL          bool   `config:"matchOnURL"`
	LocalPath           string `config:"localDirectory"`
	SpecExtensions      string `config:"extensions"`
	Extensions          []string
}

// GraviteeIntervals - intervals for the gravitee agent to use
type GraviteeIntervals struct {
	Spec  time.Duration `config:"spec"`
	Api   time.Duration `config:"product"`
	Stats time.Duration `config:"stats"`
}

// GraviteeWorkers - number of workers for the gravitee agent to use
type GraviteeWorkers struct {
	Spec int `config:"spec"`
	Api  int `config:"product"`
}

// GraviteeConfig - represents the config for gateway
type GraviteeConfig struct {
	EnvName         string              `config:"environment id"`
	OrgName         string              `config:"organization id"`
	URL             string              `config:"api url"`
	Auth            *AuthConfig         `config:"auth"`
	CloneAttributes bool                `config:"cloneAttributes"`
	Specs           *GraviteeSpecConfig `config:"specs"`
	Intervals       *GraviteeIntervals  `config:"interval"`
	Workers         *GraviteeWorkers    `config:"workers"`
	Filter          string              `config:"filter"`
}

func NewGraviteeConfig() *GraviteeConfig {
	return &GraviteeConfig{
		Auth:      &AuthConfig{},
		Intervals: &GraviteeIntervals{},
		Workers:   &GraviteeWorkers{},
		Specs:     &GraviteeSpecConfig{},
	}
}

const (
	pathAPIURL                  = "gravitee.url"
	pathAuthToken               = "gravitee.auth.token"
	pathSpecInterval            = "gravitee.interval.spec"
	pathApiInterval             = "gravitee.interval.product"
	pathStatsInterval           = "gravitee.interval.stats"
	pathorg                     = "gravitee.orgID"
	pathenv                     = "gravitee.envID"
	pathFilter                  = "gravitee.filter"
	pathSpecWorkers             = "gravitee.workers.spec"
	pathMode                    = "gravitee.DiscoveryMode"
	pathCloneAttributes         = "gravitee.cloneAttributes"
	pathApiWorkers              = "gravitee.workers.product"
	pathSpecMatchOnURL          = "gravitee.specConfig.matchOnURL"
	pathSpecLocalPath           = "gravitee.specConfig.localPath"
	pathSpecExtensions          = "gravitee.specConfig.extensions"
	pathSpecUnstructured        = "gravitee.specConfig.unstructured"
	pathSpecDisablePollForSpecs = "gravitee.specConfig.disablePollForSpecs"
)

// AddProperties - adds config needed for gravitee client
func AddProperties(rootProps properties.Properties) {
	rootProps.AddStringProperty(pathAPIURL, "http://sl1csoapp7131.pcloud.axway.int:8083/management", "URL to use when authenticating to gravitee")
	rootProps.AddStringProperty(pathAuthToken, "", "Token for the user to authenticate to gravitee")
	rootProps.AddStringProperty(pathorg, "DEFAULT", "Organization Name to use")
	rootProps.AddStringProperty(pathenv, "DEFAULT", "Environment name to use")
	rootProps.AddStringProperty(pathFilter, "", "Filter used on discovering Gravitee apis")
	rootProps.AddBoolProperty(pathCloneAttributes, false, "Set to true to copy the tags when provisioning a Api")
	rootProps.AddDurationProperty(pathSpecInterval, 30*time.Minute, "The time interval between checking for updated specs", properties.WithLowerLimit(1*time.Minute))
	rootProps.AddDurationProperty(pathApiInterval, 30*time.Second, "The time interval between checking for updated products", properties.WithUpperLimit(5*time.Minute))
	rootProps.AddDurationProperty(pathStatsInterval, 5*time.Minute, "The time interval between checking for updated stats", properties.WithLowerLimit(1*time.Minute), properties.WithUpperLimit(15*time.Minute))
	rootProps.AddIntProperty(pathSpecWorkers, 20, "Max number of workers discovering specs")
	rootProps.AddBoolProperty(pathSpecMatchOnURL, true, "Set to false to skip matching spec URLs to proxy URLs")
	rootProps.AddStringProperty(pathSpecLocalPath, "", "Path to a local directory that contains the spec files")
	rootProps.AddStringProperty(pathSpecExtensions, "json,yaml,yml", "Comma separated list of spec file extensions")
	rootProps.AddBoolProperty(pathSpecUnstructured, false, "Set to true to enable discovering apis that have no associated spec")
	rootProps.AddBoolProperty(pathSpecDisablePollForSpecs, false, "Set to true to disable polling gravitee for specs, rely on the local directory or spec URLs")

}

// ParseConfig - parse the config on startup
func ParseConfig(rootProps props) *GraviteeConfig {
	specExtensions := rootProps.StringPropertyValue(pathSpecExtensions)
	extensions := []string{}
	for _, e := range strings.Split(specExtensions, ",") {
		extensions = append(extensions, strings.TrimSpace(e))
	}
	return &GraviteeConfig{
		EnvName:         rootProps.StringPropertyValue(pathenv),
		OrgName:         rootProps.StringPropertyValue(pathorg),
		Filter:          rootProps.StringPropertyValue(pathFilter),
		CloneAttributes: rootProps.BoolPropertyValue(pathCloneAttributes),
		URL:             rootProps.StringPropertyValue(pathAPIURL),
		Intervals: &GraviteeIntervals{
			Stats: rootProps.DurationPropertyValue(pathStatsInterval),

			Spec: rootProps.DurationPropertyValue(pathSpecInterval),
			Api:  rootProps.DurationPropertyValue(pathApiInterval),
		},
		Workers: &GraviteeWorkers{
			Spec: rootProps.IntPropertyValue(pathSpecWorkers),
		},
		Auth: &AuthConfig{
			Token: rootProps.StringPropertyValue(pathAuthToken),
		},
		Specs: &GraviteeSpecConfig{
			MatchOnURL:          rootProps.BoolPropertyValue(pathSpecMatchOnURL),
			LocalPath:           rootProps.StringPropertyValue(pathSpecLocalPath),
			DisablePollForSpecs: rootProps.BoolPropertyValue(pathSpecDisablePollForSpecs),
			Unstructured:        rootProps.BoolPropertyValue(pathSpecUnstructured),
			SpecExtensions:      specExtensions,
			Extensions:          extensions,
		},
	}
}

// ValidateCfg - Validates the gateway config
func (a *GraviteeConfig) ValidateCfg() (err error) {
	if a.Auth == nil || a.Auth.Token == "" {
		return errors.New("configuration gravitee non valide: le nom d'utilisateur n'est pas configuré")
	}

	if a.Workers == nil || a.Workers.Spec < 1 {
		return errors.New("configuration gravitee non valide: les workers spec doivent être supérieurs à 0")
	}

	if a.EnvName == "" {
		return errors.New("configuration gravitee non valide: environnement invalide")
	}

	return

}

// Get Env
func (a *GraviteeConfig) GetEnv() string {
	return a.EnvName
}

func (a *GraviteeConfig) GetOrg() string {
	return a.OrgName
}

func (a *GraviteeConfig) GetURL() string {
	return a.URL
}

// GetAuth - Returns the Auth Config
func (a *GraviteeConfig) GetAuth() *AuthConfig {
	return a.Auth
}

// GetIntervals - Returns the Intervals
func (a *GraviteeConfig) GetIntervals() *GraviteeIntervals {
	return a.Intervals
}

// GetWorkers - Returns the number of Workers
func (a *GraviteeConfig) GetWorkers() *GraviteeWorkers {
	return a.Workers
}

func (a *GraviteeConfig) ShouldCloneAttributes() bool {
	return a.CloneAttributes
}
