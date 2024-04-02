package config

import (
	"errors"
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

// GraviteeConfig - represents the config for gateway
type GraviteeConfig struct {
	EnvName         string             `config:"environment id"`
	Auth            *AuthConfig        `config:"auth"`
	CloneAttributes bool               `config:"cloneAttributes"`
	Intervals       *GraviteeIntervals `config:"interval"`
	Workers         *GraviteeWorkers   `config:"workers"`
	mode            DiscoveryMode
}

// GraviteeIntervals - intervals for the gravitee agent to use
type GraviteeIntervals struct {
	Spec    time.Duration `config:"spec"`
	Product time.Duration `config:"product"`
	Stats   time.Duration `config:"stats"`
}

// GraviteeWorkers - number of workers for the gravitee agent to use
type GraviteeWorkers struct {
	Spec    int `config:"spec"`
	Product int `config:"product"`
}

type DiscoveryMode int

const (
	pathAuthURL         = "gravitee.auth.url"
	pathAuthUsername    = "gravitee.auth.username"
	pathAuthPassword    = "gravitee.auth.password"
	pathSpecInterval    = "gravitee.interval.spec"
	pathProxyInterval   = "gravitee.interval.proxy"
	pathProductInterval = "gravitee.interval.product"
	pathStatsInterval   = "gravitee.interval.stats"
	pathenv             = "gravitee.envID"
	pathSpecWorkers     = "gravitee.workers.spec"
	pathMode            = "gravitee.DiscoveryMode"
	pathCloneAttributes = "gravitee.cloneAttributes"
)

// AddProperties - adds config needed for gravitee client
func AddProperties(rootProps properties.Properties) {
	rootProps.AddStringProperty(pathAuthURL, "Https://login.gravitee.com", "URL to use when authenticating to gravitee")
	rootProps.AddStringProperty(pathAuthUsername, "", "Username to use to authenticate to gravitee")
	rootProps.AddStringProperty(pathAuthPassword, "", "Password for the user to authenticate to gravitee")
	rootProps.AddStringProperty(pathenv, "Default Environment", "Environment name to use")
	rootProps.AddBoolProperty(pathCloneAttributes, false, "Set to true to copy the tags when provisioning a Product in product mode")
	rootProps.AddDurationProperty(pathSpecInterval, 30*time.Minute, "The time interval between checking for updated specs", properties.WithLowerLimit(1*time.Minute))
	rootProps.AddDurationProperty(pathProxyInterval, 30*time.Second, "The time interval between checking for updated proxies", properties.WithUpperLimit(5*time.Minute))
	rootProps.AddDurationProperty(pathProductInterval, 30*time.Second, "The time interval between checking for updated products", properties.WithUpperLimit(5*time.Minute))
	rootProps.AddDurationProperty(pathStatsInterval, 5*time.Minute, "The time interval between checking for updated stats", properties.WithLowerLimit(1*time.Minute), properties.WithUpperLimit(15*time.Minute))
	rootProps.AddIntProperty(pathSpecWorkers, 20, "Max number of workers discovering specs")
}

// ParseConfig - parse the config on startup
func ParseConfig(rootProps props) *GraviteeConfig {
	return &GraviteeConfig{
		EnvName:         rootProps.StringPropertyValue(pathenv),
		CloneAttributes: rootProps.BoolPropertyValue(pathCloneAttributes),
		Intervals: &GraviteeIntervals{
			Stats: rootProps.DurationPropertyValue(pathStatsInterval),

			Spec:    rootProps.DurationPropertyValue(pathSpecInterval),
			Product: rootProps.DurationPropertyValue(pathProductInterval),
		},
		Workers: &GraviteeWorkers{
			Spec: rootProps.IntPropertyValue(pathSpecWorkers),
		},
		Auth: &AuthConfig{
			Username: rootProps.StringPropertyValue(pathAuthUsername),
			Password: rootProps.StringPropertyValue(pathAuthPassword),
			URL:      rootProps.StringPropertyValue(pathAuthURL),
		},
	}
}

// ValidateCfg - Validates the gateway config
func (a *GraviteeConfig) ValidateCfg() (err error) {
	if a.mode == 0 {
		return errors.New("configuration gravitee non valide: DiscoveryMode doit être proxy ou product")
	}

	if a.Auth.Username == "" {
		return errors.New("configuration gravitee non valide: le nom d'utilisateur n'est pas configuré")
	}

	if a.Auth.Password == "" {
		return errors.New("configuration gravitee non valide: le mot de passe n'est pas configuré")
	}

	if a.Workers.Spec < 1 {
		return errors.New("configuration gravitee non valide: les workers spec doivent être supérieurs à 0")
	}

	if a.EnvName == "" {
		return errors.New("configuration gravitee non valide: environnement invalide")
	}

	return

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
