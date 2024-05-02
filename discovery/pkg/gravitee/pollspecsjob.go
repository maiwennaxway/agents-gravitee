package gravitee

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/Axway/agent-sdk/pkg/apic"
	"github.com/Axway/agent-sdk/pkg/jobs"
	"github.com/Axway/agent-sdk/pkg/util/log"

	"github.com/maiwennaxway/agents-gravitee/client/pkg/gravitee"
)

type specClient interface {
	GetSpecFile(specPath string) ([]byte, error)
	GetAllSpecs() ([]gravitee.SpecDetails, error)
	IsReady() bool
}

type specCache interface {
	AddSpecToCache(id, path, name string, modDate time.Time, endpoints ...string)
	HasSpecChanged(is string, modDate time.Time) bool
}

// job that will poll for any new portals on gravitee Edge
type pollSpecsJob struct {
	jobs.Job
	firstRun    bool
	running     bool
	workers     int
	cache       specCache
	client      specClient
	logger      log.FieldLogger
	runningLock sync.Mutex
}

func newPollSpecsJob() *pollSpecsJob {
	job := &pollSpecsJob{
		firstRun:    true,
		logger:      log.NewFieldLogger().WithComponent("pollSpecs").WithPackage("gravitee"),
		runningLock: sync.Mutex{},
	}
	return job
}

func (j *pollSpecsJob) SetSpecClient(client specClient) *pollSpecsJob {
	j.client = client
	return j
}

func (j *pollSpecsJob) SetSpecCache(cache specCache) *pollSpecsJob {
	j.cache = cache
	return j
}

func (j *pollSpecsJob) SetWorkers(workers int) *pollSpecsJob {
	j.workers = workers
	return j
}

func (j *pollSpecsJob) Ready() bool {
	j.logger.Trace("specs checking if the gravitee client is ready for calls")
	//return j.client.IsReady()
	return true
}

func (j *pollSpecsJob) Status() error {
	return nil
}

func (j *pollSpecsJob) updateRunning(running bool) {
	j.runningLock.Lock()
	defer j.runningLock.Unlock()
	j.running = running
}

func (j *pollSpecsJob) isRunning() bool {
	j.runningLock.Lock()
	defer j.runningLock.Unlock()
	return j.running
}

func (j *pollSpecsJob) Execute() error {
	j.logger.Trace("executing")

	if j.isRunning() {
		j.logger.Warn("previous spec poll job run has not completed, will run again on next interval")
		return nil
	}
	j.updateRunning(true)
	defer j.updateRunning(false)

	allSpecs, err := j.client.GetAllSpecs()
	if err != nil {
		j.logger.WithError(err).Error("getting specs")
		return err
	}

	limiter := make(chan gravitee.SpecDetails, j.workers)

	wg := sync.WaitGroup{}
	wg.Add(len(allSpecs))
	for _, spec := range allSpecs {
		go func() {
			defer wg.Done()
			specDetails := <-limiter
			j.handleSpec(specDetails)
		}()
		limiter <- spec
	}

	wg.Wait()
	close(limiter)

	j.firstRun = false
	return nil
}

func (j *pollSpecsJob) FirstRunComplete() bool {
	return j.firstRun
}

func (j *pollSpecsJob) handleSpec(spec gravitee.SpecDetails) {
	logger := j.logger.WithField("specName", spec.Name).WithField("specID", spec.ID)
	logger.Trace("handling spec")
	modDate, _ := time.Parse("2006-01-02T15:04:05.000000Z", spec.Modified)
	modDate = modDate.Truncate(time.Millisecond) // truncate the nanoseconds

	endpoints := []string{}
	// get the spec content
	content, err := j.client.GetSpecFile(spec.ContentLink)
	if err != nil {
		j.logger.WithError(err).Error("getting spec content")
		return
	}

	// parse the spec
	parser := apic.NewSpecResourceParser(content, "")
	err = parser.Parse()
	if err != nil {
		j.logger.WithError(err).Error("could not parse spec")
		return
	}

	// gather spec info
	endpointDefs, err := parser.GetSpecProcessor().GetEndpoints()
	if err != nil {
		j.logger.WithError(err).Error("could not get spec endpoints")
		return
	}
	for _, ep := range endpointDefs {
		endpoints = append(endpoints, endpointToString(ep))
	}

	// add spec details to cache
	j.cache.AddSpecToCache(spec.ID, spec.ContentLink, spec.Name, modDate, endpoints...)
}

func endpointToString(endpoint apic.EndpointDefinition) string {
	if endpoint.Port > 0 &&
		((strings.ToLower(endpoint.Protocol) == "http" && endpoint.Port != 80) ||
			(strings.ToLower(endpoint.Protocol) == "https" && endpoint.Port != 443)) {
		return fmt.Sprintf("%v://%v:%v%v", endpoint.Protocol, endpoint.Host, endpoint.Port, endpoint.BasePath)
	}
	return fmt.Sprintf("%v://%v%v", endpoint.Protocol, endpoint.Host, endpoint.BasePath)
}
