package gravitee

import (
	"github.com/Axway/agent-sdk/pkg/agent"
	corecfg "github.com/Axway/agent-sdk/pkg/config"
	"github.com/Axway/agent-sdk/pkg/filter"
	"github.com/Axway/agent-sdk/pkg/jobs"

	"github.com/maiwennaxway/agents-gravitee/client/pkg/gravitee"
)

// AgentConfig - represents the config for agent
type AgentConfig struct {
	CentralCfg  corecfg.CentralConfig  `config:"central"`
	graviteeCfg *config.graviteeConfig `config:"gravitee"`
}

// Agent - Represents the Gateway client
type Agent struct {
	cfg             *AgentConfig
	graviteeClient  *gravitee.graviteeClient
	discoveryFilter filter.Filter
	stopChan        chan struct{}
	agentCache      *agentCache
}

// NewAgent - Creates a new Agent
func NewAgent(agentCfg *AgentConfig) (*Agent, error) {
	graviteeClient, err := gravitee.NewClient(agentCfg.graviteeCfg)
	if err != nil {
		return nil, err
	}

	discoveryFilter, err := filter.NewFilter(agentCfg.graviteeCfg.Filter)
	if err != nil {
		return nil, err
	}

	newAgent := &Agent{
		graviteeClient:  graviteeClient,
		cfg:             agentCfg,
		discoveryFilter: discoveryFilter,
		stopChan:        make(chan struct{}),
		agentCache:      newAgentCache(),
	}

	// newAgent.handleSubscriptions()
	provisioner := NewProvisioner(
		newAgent.graviteeClient,
		agentCfg.CentralCfg.GetCredentialConfig().GetExpirationDays(),
		agent.GetCacheManager(),
		agentCfg.graviteeCfg.IsProductMode(),
		agentCfg.graviteeCfg.ShouldCloneAttributes(),
	)
	agent.RegisterProvisioner(provisioner)

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

// registerJobs - registers the agent jobs
func (a *Agent) registerJobs() error {
	var err error

	specsJob := newPollSpecsJob(a.graviteeClient, a.agentCache, a.cfg.graviteeCfg.GetWorkers().Spec, a.cfg.graviteeCfg.IsProxyMode())
	_, err = jobs.RegisterIntervalJobWithName(specsJob, a.graviteeClient.GetConfig().GetIntervals().Spec, "Poll Specs")
	if err != nil {
		return err
	}

	var validatorReady jobFirstRunDone

	if a.cfg.graviteeCfg.IsProxyMode() {
		proxiesJob := newPollProxiesJob(a.graviteeClient, a.agentCache, specsJob.FirstRunComplete, a.cfg.graviteeCfg.GetWorkers().Proxy)
		_, err = jobs.RegisterIntervalJobWithName(proxiesJob, a.graviteeClient.GetConfig().GetIntervals().Proxy, "Poll Proxies")
		if err != nil {
			return err
		}

		// register the api validator job
		validatorReady = proxiesJob.FirstRunComplete
	} else {
		productsJob := newPollProductsJob(a.graviteeClient, a.agentCache, specsJob.FirstRunComplete, a.cfg.graviteeCfg.GetWorkers().Product, a.shouldPushAPI)
		_, err = jobs.RegisterIntervalJobWithName(productsJob, a.graviteeClient.GetConfig().GetIntervals().Product, "Poll Products")
		if err != nil {
			return err
		}

		// register the api validator job
		validatorReady = productsJob.FirstRunComplete
	}
	_, err = jobs.RegisterSingleRunJobWithName(newRegisterAPIValidatorJob(validatorReady, a.registerValidator), "Register API Validator")

	agent.NewAPIKeyCredentialRequestBuilder(agent.WithCRDIsSuspendable()).Register()
	agent.NewAPIKeyAccessRequestBuilder().Register()
	agent.NewOAuthCredentialRequestBuilder(agent.WithCRDOAuthSecret(), agent.WithCRDIsSuspendable()).Register()
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

// shouldPushAPI - callback used determine if the Product should be pushed to Central or not
func (a *Agent) shouldPushAPI(attributes map[string]string) bool {
	// Evaluate the filter condition
	return a.discoveryFilter.Evaluate(attributes)
}

// apiValidator - registers the agent jobs
func (a *Agent) apiValidator(proxyName, envName string) bool {
	// get the api with the product name and portal name
	return true
}

func (a *Agent) registerValidator() {
	agent.RegisterAPIValidator(a.apiValidator)
}
