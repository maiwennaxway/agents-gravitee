package gravitee

import (
	"github.com/Axway/agent-sdk/pkg/agent"
	corecfg "github.com/Axway/agent-sdk/pkg/config"
	"github.com/Axway/agent-sdk/pkg/filter"
	"github.com/Axway/agent-sdk/pkg/jobs"
	"github.com/maiwennaxway/agents-gravitee/client/pkg/config"
	"github.com/maiwennaxway/agents-gravitee/client/pkg/gravitee"
)

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
	apiClient       APIClient
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
	/*provisioner := NewProvisioner(
		newAgent.GraviteeClient,
		agentCfg.CentralCfg.GetCredentialConfig().GetExpirationDays(),
		agent.GetCacheManager(),
		agentCfg.GraviteeCfg.IsProductMode(),
	)
	agent.RegisterProvisioner(provisioner)*/

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

// registerJobs - registers the agent jobs
func (a *Agent) registerJobs() error {
	var err error

	specsJob := newPollSpecsJob().
		SetSpecClient(a.GraviteeClient).
		SetWorkers(a.cfg.GraviteeCfg.GetWorkers().Spec)

	_, err = jobs.RegisterIntervalJobWithName(specsJob, a.GraviteeClient.GetConfig().GetIntervals().Spec, "Poll Specs")
	if err != nil {
		return err
	}

	var validatorReady jobFirstRunDone

	apisJob := newPollAPIsJob(a.apiClient, a.agentCache, specsJob.FirstRunComplete, 10, a.shouldPushAPI)
	_, err = jobs.RegisterIntervalJobWithName(apisJob, a.GraviteeClient.GetConfig().GetIntervals().Api, "Poll Apis")
	if err != nil {
		return err
	}

	// register the api validator job
	validatorReady = apisJob.FirstRunComplete

	_, err = jobs.RegisterSingleRunJobWithName(newRegisterAPIValidatorJob(validatorReady, a.registerValidator), "Register API Validator")

	agent.NewAPIKeyCredentialRequestBuilder(agent.WithCRDIsSuspendable()).Register()
	agent.NewAPIKeyAccessRequestBuilder().Register()
	agent.NewOAuthCredentialRequestBuilder(agent.WithCRDOAuthSecret(), agent.WithCRDIsSuspendable()).Register()
	return err
}

// registerJobs - registers the agent jobs
/*func (a *Agent) registerJobs() error {
	var err error

	// create a function to let the proxy or product poll job to start if spec polling is disabled
	startPollingJob := func() bool {
		return true
	}

	parseSpec := a.cfg.GraviteeCfg.IsProxyMode() && a.cfg.GraviteeCfg.Specs.MatchOnURL // parse specs if proxy mode and match on url set
	if !a.cfg.GraviteeCfg.Specs.DisablePollForSpecs {
		specsJob := newPollSpecsJob().
			SetSpecClient(a.graviteeClient).
			SetSpecCache(a.agentCache).
			SetWorkers(a.cfg.GraviteeCfg.GetWorkers().Spec).
			SetParseSpec(parseSpec)

		_, err = jobs.RegisterIntervalJobWithName(specsJob, a.graviteeClient.GetConfig().GetIntervals().Spec, "Poll Specs")
		if err != nil {
			return err
		}
		startPollingJob = specsJob.FirstRunComplete
	}

	var validatorReady jobFirstRunDone

	if a.cfg.GraviteeCfg.IsProxyMode() {
		proxiesJob := newPollProxiesJob().
			SetSpecClient(a.graviteeClient).
			SetSpecCache(a.agentCache).
			SetSpecsReady(startPollingJob).
			SetWorkers(a.cfg.GraviteeCfg.GetWorkers().Proxy).
			SetMatchOnURL(a.cfg.GraviteeCfg.Specs.MatchOnURL)

		_, err = jobs.RegisterIntervalJobWithName(proxiesJob, a.graviteeClient.GetConfig().GetIntervals().Proxy, "Poll Proxies")
		if err != nil {
			return err
		}

		// register the api validator job
		validatorReady = proxiesJob.FirstRunComplete
	} else {
		productsJob := newPollProductsJob(a.graviteeClient, a.agentCache, startPollingJob, a.cfg.GraviteeCfg.GetWorkers().Product, a.shouldPushAPI)
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
}*/

// running - waits for a signal to stop the agent
func (a *Agent) running() {
	<-a.stopChan
}

// Stop - signals the agent to stop
func (a *Agent) Stop() {
	a.stopChan <- struct{}{}
}

// apiValidator - registers the agent jobs
func (a *Agent) apiValidator(proxyName, envName string) bool {
	// get the api with the product name and portal name
	return true
}

func (a *Agent) registerValidator() {
	agent.RegisterAPIValidator(a.apiValidator)
}
