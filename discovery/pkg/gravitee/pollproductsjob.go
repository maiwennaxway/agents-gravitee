package gravitee

/*import (
	"context"
	"sync"
	"time"

	"github.com/Axway/agent-sdk/pkg/agent"
	"github.com/Axway/agent-sdk/pkg/apic"
	"github.com/Axway/agent-sdk/pkg/jobs"
	"github.com/Axway/agent-sdk/pkg/util/log"
)

const (
	productNameField        ctxKeys = "product"
	productDisplayNameField ctxKeys = "productDisplay"
	productDetailsField     ctxKeys = "productDetails"
	productModDateField     ctxKeys = "productModDate"
	specDetailsField        ctxKeys = "specDetails"
	specModDateField        ctxKeys = "specModDate"
)

type productCacheItem struct {
	Name        string
	ModDate     time.Time
	SpecModDate time.Time
}

type productCache interface {
	GetSpecWithName(name string) (*specCacheItem, error)
	AddProductToCache(name string, modDate time.Time, specModDate time.Time)
	HasProductChanged(name string, modDate time.Time, specModDate time.Time) bool
	GetProductWithName(name string) (*productCacheItem, error)
}

type isPublishedFunc func(string) bool
type getAttributeFunc func(string, string) string

// job that will poll for any new portals on gravitee Edge
type pollProductsJob struct {
	jobs.Job
	//client 		 api
	cache            productCache
	firstRun         bool
	specsReady       jobFirstRunDone
	pubLock          sync.Mutex
	isPublishedFunc  isPublishedFunc
	getAttributeFunc getAttributeFunc
	publishFunc      agent.PublishAPIFunc
	logger           log.FieldLogger
	workers          int
	running          bool
	runningLock      sync.Mutex
	shouldPushAPI    func(map[string]string) bool
}

func newPollProductsJob(client GraviteeClient, cache productCache, specsReady jobFirstRunDone, workers int) *pollProductsJob {
	job := &pollProductsJob{
		cache:            cache,
		firstRun:         true,
		specsReady:       specsReady,
		logger:           log.NewFieldLogger().WithComponent("pollProducts").WithPackage("gravitee"),
		isPublishedFunc:  agent.IsAPIPublishedByID,
		getAttributeFunc: agent.GetAttributeOnPublishedAPIByID,
		publishFunc:      agent.PublishAPI,
		workers:          workers,
		runningLock:      sync.Mutex{},
	}
	return job
}

func (j *pollProductsJob) Ready() bool {
	j.logger.Trace("checking if the gravitee client is ready for calls")

	j.logger.Trace("checking if specs have been cached")
	return j.specsReady()
}

func (j *pollProductsJob) Status() error {
	return nil
}

func (j *pollProductsJob) updateRunning(running bool) {
	j.runningLock.Lock()
	defer j.runningLock.Unlock()
	j.running = running
}

func (j *pollProductsJob) isRunning() bool {
	j.runningLock.Lock()
	defer j.runningLock.Unlock()
	return j.running
}

func (j *pollProductsJob) Execute() error {
	j.logger.Trace("executing")

	if j.isRunning() {
		j.logger.Warn("previous spec poll job run has not completed, will run again on next interval")
		return nil
	}
	j.updateRunning(true)
	defer j.updateRunning(false)

	limiter := make(chan string, j.workers)

	wg := sync.WaitGroup{}
	/*wg.Add(len(products))
	for _, p := range apis {
		go func() {
			defer wg.Done()
			name := <-limiter
			j.handleProduct(name)
		}()
		limiter <- p
	}*/

/*wg.Wait()
	close(limiter)

	j.firstRun = false
	return nil
}

func (j *pollProductsJob) FirstRunComplete() bool {
	return !j.firstRun
}

/*func (j *pollProductsJob) handleProduct(productName string) {
	logger := j.logger.WithField(productNameField.String(), productName)
	logger.Trace("handling product")

	// get product full details
	ctx := addLoggerToContext(context.Background(), logger)
	ctx = context.WithValue(ctx, productNameField, productName)

	// try to get spec by using the name of the product
	specDetails, err := j.getSpecDetails(ctx)
	ctx = context.WithValue(ctx, specDetailsField, specDetails)
	if err != nil {
		logger.Trace("could not find spec for product by name")
		return
	}
	ctx = context.WithValue(ctx, specPathField, specDetails.ContentPath)

	// Check DiscoveryCache for API
	j.pubLock.Lock() // only publish one at a time
	defer j.pubLock.Unlock()
	value := agent.GetAttributeOnPublishedAPIByID(productName, "hash")

	err = nil
	if !j.isPublishedFunc(productName) {
		// call new API
		err = j.publishAPI(*serviceBody, hashString, cacheKey)
	} else if value != hashString {
		// handle update
		log.Tracef("%s has been updated, push new revision", productName)
		serviceBody.APIUpdateSeverity = "Major"
		log.Tracef("%+v", serviceBody)
		err = j.publishAPI(*serviceBody, hashString, cacheKey)
	}

	if err == nil {
		j.cache.AddProductToCache(productName, ctx.Value(productModDateField).(time.Time), specDetails.ModDate)
	}
}

func (j *pollProductsJob) getSpecDetails(ctx context.Context) (*specCacheItem, error) {
	productName := getStringFromContext(ctx, productNameField)
	displayName := getStringFromContext(ctx, productDisplayNameField)

	specDetails, err := j.cache.GetSpecWithName(productName)
	if err != nil {
		// try to find the spec details with the display name before giving up
		specDetails, err = j.cache.GetSpecWithName(displayName)
	}
	return specDetails, err
}*/
