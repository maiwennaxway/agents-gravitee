package gravitee

import (
	"context"
	"sync"

	"github.com/Axway/agent-sdk/pkg/jobs"
	"github.com/Axway/agent-sdk/pkg/util/log"
	"github.com/maiwennaxway/agents-gravitee/client/pkg/config"
	"github.com/maiwennaxway/agents-gravitee/client/pkg/gravitee"
)

const (
	apiIdField          ctxKeys = "api"
	apiDisplayIdField   ctxKeys = "apiDisplay"
	apiDetailsField     ctxKeys = "apiDetails"
	apiModDateField     ctxKeys = "apiModDate"
	specDetailsFieldApi ctxKeys = "specDetails"
	specModDateFieldApi ctxKeys = "specModDate"
)

type ctxKeys string

type APIClient interface {
	GetConfig() *config.GraviteeConfig
	GetApis() (gravitee.Apis, error)
	GetApi(ApiID string) (*ApiProduct, error)
	GetSpecFile(specPath string) ([]byte, error)
	IsReady() bool
}

type getApiAttributeFunc func(string, string) string

type pollAPIsJob struct {
	jobs.Job
	logger      log.FieldLogger
	Client      gravitee.GraviteeClient
	apiClient   APIClient
	firstRun    bool
	specsReady  jobFirstRunDone
	workers     int
	running     bool
	runningLock sync.Mutex
}

func newPollAPIsJob(client APIClient, specsReady jobFirstRunDone, workers int) *pollAPIsJob {
	job := &pollAPIsJob{
		logger:      log.NewFieldLogger().WithComponent("pollAPIs").WithPackage("gravitee"),
		apiClient:   client,
		firstRun:    true,
		specsReady:  specsReady,
		workers:     workers,
		runningLock: sync.Mutex{},
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
	j.logger.Trace("executing")

	if j.isRunning() {
		j.logger.Warn("previous spec poll job run has not completed, will run again on next interval")
		return nil
	}
	j.updateRunning(true)
	defer j.updateRunning(false)

	limiter := make(chan string, j.workers)

	wg := sync.WaitGroup{}

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
	j.logger.Trace("checking if the gravitee client is ready for calls")
	return j.specsReady()
}

func (j *pollAPIsJob) FirstRunComplete() bool {
	return !j.firstRun
}

// addLoggerToContext ajoute un logger au contexte fourni
func addLoggerToContext(ctx context.Context, logger log.FieldLogger) context.Context {
	return context.WithValue(ctx, "logger", logger)
}

// getStringFromContext extrait une valeur de type chaîne de caractères à partir du contexte pour une clé donnée
func getStringFromContext(ctx context.Context, key ctxKeys) string {
	if value, ok := ctx.Value(key).(string); ok {
		return value
	}
	return "" // Valeur par défaut si la clé n'est pas présente ou si la valeur n'est pas de type string
}

func (j *pollAPIsJob) HandleAPI(Api string) []string {
	logger := j.logger
	logger.Trace("handling Api")
	ctx := addLoggerToContext(context.Background(), logger)
	ctx = context.WithValue(ctx, Api, Api)

	apis := []string{}
	return apis
}

// getAPIDetails récupère les détails d'une API à partir de son ID
func (j *pollAPIsJob) getAPIDetails(ctx context.Context) (*APIDetails, error) {
	// Récupération de l'ID de l'API à partir du contexte
	apiID := getStringFromContext(ctx, apiIdField)

	// Utilisation de l'API client ou d'autres méthodes pour récupérer les détails de l'API
	// Par exemple, une requête à une base de données ou à un service externe
	apiDetails, err := j.Client.GetApibyApiId(apiID)
	if err != nil {
		return nil, err
	}

	return apiDetails, nil
}
