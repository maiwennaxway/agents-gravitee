package gravitee

import (
	coreapi "github.com/Axway/agent-sdk/pkg/api"
	"github.com/Axway/agent-sdk/pkg/jobs"
	"github.com/Axway/agent-sdk/pkg/util/log"
)

/*const (
	graviteeAuthPath      = "/oauth/token"
	graviteeAuthCheckPath = "/login"
	grantTypeKey          = "grant_type"
	usernameKey           = "username"
	passwordKey           = "password"
	refreshTokenKey       = "refresh_token"
)*/

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

type authJob struct {
	jobs.Job
	apiClient coreapi.Client
	url       string
}

/*func (j *authJob) Ready() bool {
	if err != nil {
		log.Error(err)
		return false
	}
	return true
}*/

func (j *authJob) Status() error {
	return nil
}

func (j *authJob) checkConnection() error {
	request := coreapi.Request{
		Method: coreapi.GET,
		URL:    j.url,
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
	return err
}
