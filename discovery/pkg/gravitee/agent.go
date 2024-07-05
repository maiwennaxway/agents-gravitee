package gravitee

import (
	"github.com/Axway/agent-sdk/pkg/agent"
	"github.com/Axway/agent-sdk/pkg/apic/provisioning"
	corecfg "github.com/Axway/agent-sdk/pkg/config"
	"github.com/Axway/agent-sdk/pkg/filter"
	"github.com/Axway/agent-sdk/pkg/jobs"
	"github.com/maiwennaxway/agents-gravitee/client/pkg/config"
	"github.com/maiwennaxway/agents-gravitee/client/pkg/gravitee"
	"github.com/sirupsen/logrus"
)

const ApiKeyName = provisioning.APIKeyARD

// AgentConfig - represents the config for agent
type AgentConfig struct {
	CentralCfg  corecfg.CentralConfig  `config:"central"`
	GraviteeCfg *config.GraviteeConfig `config:"gravitee"`
}

// Agent - Represents the Gateway client
type Agent struct {
	cfg             *AgentConfig
	GraviteeClient  *gravitee.GraviteeClient
	discoveryFilter filter.Filter
	stopChan        chan struct{}
	agentCache      *agentSpec
}

// NewAgent - Creates a new Agent
func NewAgent(agentCfg *AgentConfig) (*Agent, error) {
	GraviteeClient, err := gravitee.NewClient(agentCfg.GraviteeCfg)
	if err != nil {
		return nil, err
	}

	discoveryFilter, err := filter.NewFilter(agentCfg.GraviteeCfg.Filter)
	if err != nil {
		return nil, err
	}

	newAgent := &Agent{
		GraviteeClient:  GraviteeClient,
		cfg:             agentCfg,
		discoveryFilter: discoveryFilter,
		stopChan:        make(chan struct{}),
		agentCache:      newAgentSpec(),
	}

	// newAgent.handleSubscriptions()
	provisioner := NewProvisioner(
		newAgent.GraviteeClient,
		agentCfg.CentralCfg.GetCredentialConfig().GetExpirationDays(),
		agent.GetCacheManager(),
		agentCfg.GraviteeCfg.ShouldCloneAttributes(),
	)
	agent.RegisterProvisioner(provisioner)
	registerKeyAuth()

	return newAgent, nil
}

func (a *Agent) Run() error {
	// Start the agent jobs
	err := a.registerJobs()
	if err != nil {
		return err
	}
	a.running()
	return nil
}

func (a *Agent) shouldPushAPI(attributes map[string]string) bool {
	// Evaluate the filter condition
	return a.discoveryFilter.Evaluate(attributes)
}

func registerKeyAuth() {
	//"The api key. Leave empty for autogeneration"
	_, err := agent.NewAPIKeyAccessRequestBuilder().SetName(ApiKeyName).Register()
	if err != nil {
		logrus.Error("Error registering API key Access Request")
	}
	_, err = agent.NewAPIKeyCredentialRequestBuilder(agent.WithCRDIsSuspendable()).IsRenewable().Register()
	if err != nil {
		logrus.Error("Error registering API Credential Access Request")
	}
}

// registerJobs - registers the agent jobs
func (a *Agent) registerJobs() error {
	var err error

	var validatorReady jobFirstRunDone

	apisJob := newPollAPIsJob(a.GraviteeClient, a.agentCache, 10, a.shouldPushAPI)
	_, err = jobs.RegisterIntervalJobWithName(apisJob, a.GraviteeClient.GetConfig().GetIntervals().Api, "Poll Apis")
	if err != nil {
		return err
	}

	// register the api validator job
	validatorReady = apisJob.FirstRunComplete

	_, err = jobs.RegisterSingleRunJobWithName(newRegisterAPIValidatorJob(validatorReady, a.registerValidator), "Register API Validator")
	return err
}

// running - waits for a signal to stop the agent
func (a *Agent) running() {
	<-a.stopChan
}

// Stop - signals the agent to stop
func (a *Agent) Stop() {
	a.stopChan <- struct{}{}
}

// apiValidator - registers the agent jobs
func (a *Agent) apiValidator(apiName, envName string) bool {
	// get the api with the product name and portal name
	return true
}

func (a *Agent) registerValidator() {
	agent.RegisterAPIValidator(a.apiValidator)
}
