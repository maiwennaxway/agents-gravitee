package gravitee

import (
	"context"
	"fmt"
	"path"
	"strings"
	"sync"
	"time"

	"github.com/Axway/agent-sdk/pkg/agent"
	"github.com/Axway/agent-sdk/pkg/apic"
	v1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
	"github.com/Axway/agent-sdk/pkg/jobs"

	coreutil "github.com/Axway/agent-sdk/pkg/util"
	"github.com/Axway/agent-sdk/pkg/util/log"
	"github.com/maiwennaxway/agents-gravitee/client/pkg/config"
	"github.com/maiwennaxway/agents-gravitee/client/pkg/gravitee"

	"github.com/maiwennaxway/agents-gravitee/discovery/pkg/util"

	//"github.com/maiwennaxway/agents-gravitee/client/pkg/gravitee"
	"github.com/maiwennaxway/agents-gravitee/client/pkg/gravitee/models"
)

const specLocalTag = "spec_local"

const (
	apiIdField          ctxKeys = "api"
	apiDisplayIdField   ctxKeys = "apiDisplay"
	apiDetailsField     ctxKeys = "apiDetails"
	apiModDateField     ctxKeys = "apiModDate"
	specDetailsFieldApi ctxKeys = "specDetails"
	specModDateFieldApi ctxKeys = "specModDate"
	specPathField       ctxKeys = "specPath"
	gatewayType                 = "Gravitee"
)

type APIClient interface {
	GetConfig() *config.GraviteeConfig
	GetApis() (gravitee.Apis, error)
	GetApi(ApiID, EnvId string) (*models.Api, error)
	GetSpecFile(specPath string) ([]byte, error)
	IsReady() bool
}

type ApiCacheItem struct {
	Name     string
	SpecHash string
	ModDate  time.Time
}

type APISpec interface {
	GetSpecWithName(name string) (*specItem, error)
	AddApiToCache(name string, modDate time.Time, specHash string)
}

type isPublishedFunc func(string) bool
type getAttributeFunc func(string, string) string

type pollAPIsJob struct {
	jobs.Job
	logger log.FieldLogger
	//Client           Gravitee.GraviteeClient
	apiClient        APIClient
	specClient       APISpec
	firstRun         bool
	specsReady       jobFirstRunDone
	pubLock          sync.Mutex
	isPublishedFunc  isPublishedFunc
	publishFunc      agent.PublishAPIFunc
	getAttributeFunc getAttributeFunc
	workers          int
	running          bool
	runningLock      sync.Mutex
	shouldPushAPI    func(map[string]string) bool
}

func newPollAPIsJob(client APIClient, cache APISpec, specsReady jobFirstRunDone, workers int, shouldPushAPI func(map[string]string) bool) *pollAPIsJob {
	job := &pollAPIsJob{
		logger:           log.NewFieldLogger().WithComponent("pollAPIs").WithPackage("gravitee"),
		apiClient:        client,
		specClient:       cache,
		firstRun:         true,
		specsReady:       specsReady,
		isPublishedFunc:  agent.IsAPIPublishedByID,
		getAttributeFunc: agent.GetAttributeOnPublishedAPIByID,
		publishFunc:      agent.PublishAPI,
		workers:          workers,
		runningLock:      sync.Mutex{},
		shouldPushAPI:    shouldPushAPI,
	}

	return job
}

func (j *pollAPIsJob) updateRunning(running bool) {
	j.runningLock.Lock()
	defer j.runningLock.Unlock()
	j.running = running
}

func (j *pollAPIsJob) isRunning() bool {
	j.runningLock.Lock()
	defer j.runningLock.Unlock()
	return j.running
}

// Execute executes the job
func (j *pollAPIsJob) Execute() error {
	//débuter l'éxecution du poll des APIS
	j.logger.Trace("Executing")

	if j.isRunning() {
		j.logger.Warn("previous spec poll job run has not completed, will run again on next interval")
		return nil
	}
	j.updateRunning(true)
	defer j.updateRunning(false)

	apis, err := j.apiClient.GetApis()
	j.logger.Trace(fmt.Sprint("Apis :", apis))
	if err != nil {
		j.logger.WithError(err).Error("getting apis")
		return err
	}

	limiter := make(chan string, j.workers)

	wg := sync.WaitGroup{}
	wg.Add(len(apis))
	for _, p := range apis {
		j.logger.Trace("apis : %s", p)
		go func() {
			defer wg.Done()
			name := <-limiter
			j.logger.Trace("name : %s", name)
			j.HandleAPI(name)
		}()
		limiter <- p
	}

	wg.Wait()
	close(limiter)

	j.firstRun = false
	return nil
}

// Status retourne le statut actuel du job
func (j *pollAPIsJob) Status() error {
	return nil
}

// Ready retourne true si le job est prêt à s'exécuter
func (j *pollAPIsJob) Ready() bool {
	j.logger.Trace("checking if specs have been cached")
	return true //j.specsReady()
}

func (j *pollAPIsJob) FirstRunComplete() bool {
	return j.firstRun
}

func (j *pollAPIsJob) getSpecDetails(ctx context.Context, apiDetails *models.Api) (context.Context, error) {
	// Recherche de la spécification associée à l'API
	for _, att := range apiDetails.Attributes {
		// Recherche de la balise spécifique dans les attributs de l'API
		if strings.ToLower(att.Name) == specLocalTag {
			// Si la balise est trouvée, ajout du chemin de la spécification au contexte
			ctx = context.WithValue(ctx, specPathField, strings.Join([]string{specLocalTag, att.Value}, "_"))
			break
		}
	}

	specDetails, err := j.specClient.GetSpecWithName(apiDetails.Name)
	if err != nil {
		// try to find the spec details with the display name before giving up
		specDetails, err = j.specClient.GetSpecWithName(apiDetails.Id)
		if err != nil {
			return ctx, err
		}
	}
	ctx = context.WithValue(ctx, specPathField, specDetails.ContentPath)
	// Retourner le contexte mis à jour avec les détails de l'API et la spécification, ainsi que les détails de l'API
	return ctx, nil
}

func (j *pollAPIsJob) buildServiceBody(ctx context.Context, api *models.Api) (*apic.ServiceBody, uint64, error) {
	logger := getLoggerFromContext(ctx)
	specPath := getStringFromContext(ctx, specPathField)

	var spec []byte
	var err error
	if strings.HasPrefix(specPath, specLocalTag) {
		logger = logger.WithField("specLocalDir", "true")
		fileName := strings.TrimPrefix(specPath, specLocalTag+"_")
		config := j.apiClient.GetConfig()
		if config != nil && config.Specs != nil {
			filePath := path.Join(config.Specs.LocalPath, fileName)
			spec, err = loadSpecFile(logger, filePath)
		} else {
			return nil, 0, err
		}

	} else {
		logger = logger.WithField("specLocalDir", "false")
		// get the spec to build the service body
		spec, err = j.apiClient.GetSpecFile(specPath)
	}

	if err != nil {
		logger.WithError(err).Error("could not download spec")
		return nil, 0, err
	}

	if len(spec) == 0 && !j.apiClient.GetConfig().Specs.Unstructured {
		return nil, 0, fmt.Errorf("spec had no content")
	}

	specHash, _ := coreutil.ComputeHash(spec)

	// create the agent details with the modification dates
	serviceDetails := map[string]interface{}{
		"apiModDate":      time.UnixMilli(int64(api.LastModifiedAt)).Format(v1.APIServerTimeFormat),
		"specContentHash": specHash,
	}

	// create attributes to be added to service
	serviceAttributes := make(map[string]string)
	for _, att := range api.Attributes {
		name := strings.ToLower(att.Name)
		name = strings.ReplaceAll(name, " ", "_")
		serviceAttributes[name] = att.Value
	}

	logger.Debug("creating service body")
	sb, err := apic.NewServiceBodyBuilder().
		SetID(api.Id).
		SetAPIName(api.Name).
		SetDescription(api.Description).
		SetAPISpec(spec).
		SetTitle(api.Name).
		SetServiceAttribute(serviceAttributes).
		SetServiceAgentDetails(serviceDetails).
		Build()
	return &sb, specHash, err
}

type APIContextKey string

// Définir une clé pour l'API
const APIKey APIContextKey = "api"

func (j *pollAPIsJob) HandleAPI(ApiID string) {
	logger := j.logger
	logger.Trace("handling Api")
	ctx := addLoggerToContext(context.Background(), logger)
	//ctx = context.WithValue(ctx, APIKey, Api)

	// get the full api details
	apidetails, err := j.apiClient.GetApi(ApiID, "DEFAULT")
	if err != nil {
		logger.WithError(err).Trace("could not retrieve api details")
		return
	}
	logger = logger.WithField("ApiDisplay", apidetails.Name)

	if !j.shouldPublishAPI(logger, apidetails) {
		logger.Trace("Api has been filtered out")
		return
	}

	// try to get spec by using the name of the api
	ctx, err = j.getSpecDetails(ctx, apidetails)
	if err != nil {
		logger.Trace("could not find spec for api by name")
		return
	}

	// create service
	serviceBody, specHash, err := j.buildServiceBody(ctx, apidetails)
	if err != nil {
		logger.WithError(err).Error("building service body")
		return
	}

	serviceBodyHash, _ := coreutil.ComputeHash(*serviceBody)
	hashString := coreutil.ConvertUnitToString(serviceBodyHash)
	spechashString := util.ConvertUnitToString(specHash)
	cacheKey := createApiCacheKey(ApiID)

	j.pubLock.Lock() // only publish one at a time
	defer j.pubLock.Unlock()
	value := agent.GetAttributeOnPublishedAPIByID(apidetails.Id, "hash")

	err = nil
	if !j.isPublishedFunc(apidetails.Id) {
		// call new API
		_ = j.PublishAPI(*serviceBody, hashString, cacheKey)
	} else if value != hashString {
		// handle update
		log.Tracef("%s has been updated, push new revision", ApiID)
		serviceBody.APIUpdateSeverity = "Major"
		log.Tracef("%+v", serviceBody)
		_ = j.PublishAPI(*serviceBody, hashString, cacheKey)
	}

	j.specClient.AddApiToCache(apidetails.Id, time.UnixMilli(int64(apidetails.LastModifiedAt)), spechashString)

}

func (j *pollAPIsJob) shouldPublishAPI(logger log.FieldLogger, api *models.Api) bool {
	// get the api attributes in a map
	attributes := make(map[string]string)
	for _, att := range api.Attributes {
		// ignore access attribute
		if strings.ToLower(att.Name) == "access" {
			continue
		}
		attributes[att.Name] = att.Value
	}
	logger = logger.WithField("attributes", attributes)

	if val, ok := attributes[agentApiTagName]; ok && val == agentApiTagValue {
		logger.Trace("Api was created by agent, skipping")
		return false
	}

	logger.WithField("attributes", attributes).Trace("checking against discovery filter")
	return j.shouldPushAPI(attributes)
}

func (j *pollAPIsJob) PublishAPI(serviceBody apic.ServiceBody, hashString, cacheKey string) error {
	serviceBody.ServiceAttributes["GatewayType"] = gatewayType
	serviceBody.ServiceAgentDetails["hash"] = hashString
	serviceBody.InstanceAgentDetails[cacheKeyAttribute] = cacheKey

	err := j.publishFunc(serviceBody)
	if err == nil {
		log.Infof("Published API %s to AMPLIFY Central", serviceBody.NameToPush)
		return err
	}
	return nil
}
