package gravitee

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	coreapi "github.com/Axway/agent-sdk/pkg/api"
	"github.com/Axway/agent-sdk/pkg/jobs"
	"github.com/Axway/agent-sdk/pkg/util/log"
)

const (
	graviteeAuthPath      = "/oauth/token"
	graviteeAuthCheckPath = "/login"
	grantTypeKey          = "grant_type"
	usernameKey           = "username"
	passwordKey           = "password"
	refreshTokenKey       = "refresh_token"
)

type AuthJobOpt func(*AuthJob)

func newAuthJob(opts ...AuthJobOpt) *AuthJob {
	a := &AuthJob{}
	for _, o := range opts {
		o(a)
	}
	return a
}

func withAPIClient(apiClient coreapi.Client) AuthJobOpt {
	return func(a *AuthJob) {
		a.apiClient = apiClient
	}
}

func withUsername(username string) AuthJobOpt {
	return func(a *AuthJob) {
		a.username = username
	}
}

func withPassword(password string) AuthJobOpt {
	return func(a *AuthJob) {
		a.password = password
	}
}

func withURL(url string) AuthJobOpt {
	return func(a *AuthJob) {
		a.url = url
	}
}

type AuthJob struct {
	jobs.Job
	apiClient    coreapi.Client
	refreshToken string
	username     string
	password     string
	url          string
}

func (j *AuthJob) Ready() bool {
	err := j.passwordAuth()
	if err != nil {
		log.Error(err)
		return false
	}
	return true
}

func (j *AuthJob) Status() error {
	return nil
}

func (j *AuthJob) checkConnection() error {
	request := coreapi.Request{
		Method: coreapi.GET,
		URL:    fmt.Sprintf("%s%s", j.url, graviteeAuthCheckPath),
	}

	// Validate we can reach the gravitee auth server
	_, err := j.apiClient.Send(request)
	if err != nil {
		log.Errorf(err.Error())
		return err
	}

	return nil
}

func (j *AuthJob) Execute() error {
	err := j.checkConnection()
	if err != nil {
		return err
	}

	if j.refreshToken != "" {
		err = j.refreshAuth()
	}
	if err != nil {
		err = j.passwordAuth()
	}
	return err
}

func (j *AuthJob) passwordAuth() error {
	log.Tracef("Getting new auth token")
	authData := url.Values{}
	authData.Set(grantTypeKey, password.String())
	authData.Set(usernameKey, j.username)
	authData.Set(passwordKey, j.password)

	err := j.postAuth(authData)
	if err != nil {
		// clear out the refreshToken attribute
		j.refreshToken = ""
	}
	return err
}

func (j *AuthJob) refreshAuth() error {
	log.Tracef("Refreshing auth token")
	authData := url.Values{}
	authData.Set(grantTypeKey, refresh.String())
	authData.Set(refreshTokenKey, j.refreshToken)

	return j.postAuth(authData)
}

func (j *AuthJob) postAuth(authData url.Values) error {
	basicAuth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", j.username, j.password)))
	request := coreapi.Request{
		Method: coreapi.POST,
		URL:    fmt.Sprintf("%s%s", j.url, graviteeAuthPath),
		Headers: map[string]string{
			"Content-Type":  "application/x-www-form-urlencoded",
			"Authorization": "Basic " + basicAuth,
		},
		Body: []byte(authData.Encode()),
	}

	// Get the initial authentication token
	response, err := j.apiClient.Send(request)
	if err != nil {
		log.Errorf(err.Error())
		return err
	}

	// if the response code is not ok log and return an err
	if response.Code != http.StatusOK {
		err := fmt.Errorf("unexpected response code %d from authentication call: %s", response.Code, response.Body)
		log.Error(err)
		return err
	}

	// save this refreshToken and send the token to the client
	authResponse := AuthResponse{}
	json.Unmarshal(response.Body, &authResponse)
	log.Trace(authResponse.AccessToken)
	j.refreshToken = authResponse.RefreshToken
	return nil
}
