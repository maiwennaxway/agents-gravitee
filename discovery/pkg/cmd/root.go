package cmd

import (
	corecmd "github.com/Axway/agent-sdk/pkg/cmd"
	corecfg "github.com/Axway/agent-sdk/pkg/config"
	"github.com/Axway/agent-sdk/pkg/migrate"
	"github.com/Axway/agent-sdk/pkg/notify"

	config "github.com/maiwennaxway/agents-gravitee/client/pkg/config"

	"github.com/maiwennaxway/agents-gravitee/discovery/pkg/gravitee"
)

// RootCmd - Agent root command
var RootCmd corecmd.AgentRootCmd
var graviteeClient *gravitee.Agent

func init() {
	// Create new root command with callbacks to initialize the agent config and command execution.
	// The first parameter identifies the name of the yaml file that agent will look for to load the config
	RootCmd = corecmd.NewRootCmd(
		"gravitee_discovery_agent", // Name of the yaml file
		"Gravitee Discovery Agent", // Agent description
		initConfig,                 // Callback for initializing the agent config
		run,                        // Callback for executing the agent
		corecfg.DiscoveryAgent,     // Agent Type (Discovery or Traceability)
	)

	// Get the root command properties and bind the config property in YAML definition
	rootProps := RootCmd.GetProperties()
	config.AddProperties(rootProps)

	migrate.MatchAttrPattern("-hash")
}

// Callback that agent will call to process the execution
func run() error {
	return graviteeClient.Run()
}

// Callback that agent will call to initialize the config. CentralConfig is parsed by Agent SDK
// and passed to the callback allowing the agent code to access the central config
func initConfig(centralConfig corecfg.CentralConfig) (interface{}, error) {
	rootProps := RootCmd.GetProperties()
	// Parse the config from bound properties and setup gateway config
	agentConfig := &gravitee.AgentConfig{
		CentralCfg:  centralConfig,
		graviteeCfg: config.ParseConfig(rootProps),
	}
	notify.SetSubscriptionConfig(centralConfig.GetSubscriptionConfig())

	var err error
	graviteeClient, err = gravitee.NewAgent(agentConfig)

	return agentConfig, err
}
