package gravitee

import (
	"github.com/Axway/agent-sdk/pkg/agent"
	corecfg "github.com/Axway/agent-sdk/pkg/config"
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
	cfg            *AgentConfig
	GraviteeClient *gravitee.GraviteeClient
	stopChan       chan struct{}
	agentCache     *agentCache
}

// NewAgent - Creates a new Agent
func NewAgent(agentCfg *AgentConfig) (*Agent, error) {
	GraviteeClient, err := gravitee.NewClient(agentCfg.GraviteeCfg)
	if err != nil {
		return nil, err
	}

	newAgent := &Agent{
		GraviteeClient: GraviteeClient,
		cfg:            agentCfg,
		stopChan:       make(chan struct{}),
		agentCache:     newAgentCache(),
	}

	// newAgent.handleSubscriptions()
	provisioner := NewProvisioner(
		newAgent.GraviteeClient,
		agentCfg.CentralCfg.GetCredentialConfig().GetExpirationDays(),
		agent.GetCacheManager(),
		agentCfg.GraviteeCfg.IsProductMode(),
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

	specsJob := newPollSpecsJob(a.GraviteeClient, a.agentCache, a.cfg.GraviteeCfg.GetWorkers().Spec, a.cfg.GraviteeCfg.IsProxyMode())
	_, err = jobs.RegisterIntervalJobWithName(specsJob, a.GraviteeClient.GetConfig().GetIntervals().Spec, "Poll Specs")
	if err != nil {
		return err
	}

	var validatorReady jobFirstRunDone

	if a.cfg.GraviteeCfg.IsProxyMode() {
		proxiesJob := newPollProxiesJob(a.GraviteeClient, a.agentCache, specsJob.FirstRunComplete, a.cfg.GraviteeCfg.GetWorkers().Proxy)
		_, err = jobs.RegisterIntervalJobWithName(proxiesJob, a.GraviteeClient.GetConfig().GetIntervals().Proxy, "Poll Proxies")
		if err != nil {
			return err
		}

		// register the api validator job
		validatorReady = proxiesJob.FirstRunComplete
	} else {
		productsJob := newPollProductsJob(a.GraviteeClient, a.agentCache, specsJob.FirstRunComplete, a.cfg.GraviteeCfg.GetWorkers().Product)
		_, err = jobs.RegisterIntervalJobWithName(productsJob, a.GraviteeClient.GetConfig().GetIntervals().Product, "Poll Products")
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

// apiValidator - registers the agent jobs
func (a *Agent) apiValidator(proxyName, envName string) bool {
	// get the api with the product name and portal name
	return true
}

func (a *Agent) registerValidator() {
	agent.RegisterAPIValidator(a.apiValidator)
}
