package gravitee

import (
	coreapi "github.com/Axway/agent-sdk/pkg/api"
	"github.com/Axway/agent-sdk/pkg/jobs"
	"github.com/Axway/agent-sdk/pkg/util/log"
)

type authJobOpt func(*authJob)

func newAuthJob(opts ...authJobOpt) *authJob {
	a := &authJob{}
	for _, o := range opts {
		o(a)
	}
	return a
}

func withAPIClient(apiClient coreapi.Client) authJobOpt {
	return func(a *authJob) {
		a.apiClient = apiClient
	}
}

func withURL(url string) authJobOpt {
	return func(a *authJob) {
		a.url = url
	}
}

func withToken(token string) authJobOpt {
	return func(a *authJob) {
		a.token = token
	}
}

type authJob struct {
	jobs.Job
	apiClient coreapi.Client
	url       string
	token     string
}

func (j *authJob) Ready() bool {
	err := j.checkConnection()
	if err != nil {
		log.Error(err)
		return false
	}
	return true
}

func (j *authJob) Status() error {
	return nil
}

func (j *authJob) checkConnection() error {
	request := coreapi.Request{
		Method: coreapi.GET,
		URL:    j.url,
		Headers: map[string]string{
			"Content-Type":  "application/x-www-form-urlencoded",
			"Authorization": "Bearer " + j.token,
		},
	}

	// Validate we can reach the gravitee auth server
	_, err := j.apiClient.Send(request)
	if err != nil {
		log.Errorf(err.Error())
		return err
	}

	return nil
}

func (j *authJob) Execute() error {
	err := j.checkConnection()
	if err != nil {
		return err
	}
	return nil
}
