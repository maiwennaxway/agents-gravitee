package apigee

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Axway/agent-sdk/pkg/cache"
	corecfg "github.com/Axway/agent-sdk/pkg/config"
	"github.com/Axway/agent-sdk/pkg/jobs"
	"github.com/Axway/agent-sdk/pkg/traceability"

	"github.com/maiwennaxway/agents-gravitee/client/pkg/gravitee"
	"github.com/maiwennaxway/agents-gravitee/traceability/pkg/gravitee/definitions"
	"github.com/maiwennaxway/agents-gravitee/traceability/pkg/gravitee/statsmock"
)

const (
	apiStatCacheFile = "stat-cache-data.json"
)

// AgentConfig - represents the config for agent
type AgentConfig struct {
	CentralCfg  corecfg.CentralConfig  `config:"central"`
	graviteeCfg *config.graviteeConfig `config:"gravitee"`
}

var thisAgent *Agent

// GetAgent - returns the agent
func GetAgent() *Agent {
	return thisAgent
}

// Agent - Represents the Gateway client
type Agent struct {
	cfg            *AgentConfig
	graviteeClient *gravitee.graviteeClient
	statCache      cache.Cache
	cacheFilePath  string
	ready          bool
}

// NewAgent - Creates a new Agent
func NewAgent(agentCfg *AgentConfig) (*Agent, error) {
	graviteeClient, err := gravitee.NewClient(agentCfg.graviteeCfg)
	if err != nil {
		return nil, err
	}

	thisAgent = &Agent{
		graviteeClient: graviteeClient,
		cfg:            agentCfg,
		statCache:      cache.New(),
	}

	return thisAgent, nil
}

// BeatsReady - signal that the beats are ready
func (a *Agent) BeatsReady() {
	a.setupCache()
	a.registerPollStatsJob()
	a.ready = true
}

func (a *Agent) IsReady() bool {
	return a.ready && a.graviteeClient.IsReady()
}

// setupCache -
func (a *Agent) setupCache() {
	a.cacheFilePath = filepath.Join(traceability.GetDataDirPath(), "cache", apiStatCacheFile)
	a.statCache.Load(a.cacheFilePath)
}

func (a *Agent) registerPollStatsJob() (string, error) {
	var client definitions.StatsClient = a.graviteeClient

	val := os.Getenv("QA_SIMULATE_gravitee_STATS")
	if strings.ToLower(val) == "true" {
		products, _ := a.graviteeClient.GetProducts()
		client = statsmock.NewStatsMock(a.graviteeClient, products)
	}

	// create the job that runs every minute
	baseOpts := []func(*pollgraviteeStats){
		withStatsClient(client),
		withIsReady(a.IsReady),
		withStatsCache(a.statCache),
		withCachePath(a.cacheFilePath),
		withAllTraffic(a.cfg.graviteeCfg.ShouldReportAllTraffic()),
		withNotSetTraffic(a.cfg.graviteeCfg.ShouldReportNotSetTraffic()),
	}
	if a.cfg.graviteeCfg.IsProductMode() {
		baseOpts = append(baseOpts, withProductMode())
	}

	lastStatTimeIface, err := a.statCache.Get(lastStartTimeKey)
	var lastStartTime time.Time
	if err == nil {
		// there was a last time in the cache
		lastStartTime, _ = time.Parse(time.RFC3339Nano, lastStatTimeIface.(string))
	}

	if !lastStartTime.IsZero() {
		// last start time not zero

		// create the job that executes once to get all the stats that were missed
		catchUpJob := newPollStatsJob(
			append(baseOpts,
				withStartTime(lastStartTime),
			)...,
		)
		catchUpJob.Execute()

		// register the regular running job after one interval has passed
		go func() {
			time.Sleep(a.cfg.graviteeCfg.Intervals.Stats)
			job := newPollStatsJob(append(baseOpts, withCacheClean(), withStartTime(catchUpJob.startTime))...)
			jobs.RegisterIntervalJobWithName(job, a.cfg.graviteeCfg.Intervals.Stats, "gravitee Stats")
		}()
	} else {
		// register a regular running job, only grabbing hte last hour of stats
		job := newPollStatsJob(append(baseOpts, withCacheClean(), withStartTime(time.Now().Add(time.Hour*-1).Truncate(time.Minute)))...)
		jobs.RegisterIntervalJobWithName(job, a.cfg.graviteeCfg.Intervals.Stats, "gravitee Stats")
	}

	return "", nil
}
