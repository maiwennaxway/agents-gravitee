package gravitee

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/maiwennaxway/agents-gravitee/client/pkg/gravitee/models"
	"github.com/sirupsen/logrus"
)

// SubscribetoAnAPI - Request for your application to subscribe to an api
func (a *GraviteeClient) SubscribetoAnAPI(appId, planId string) (*models.Subscriptions, error) {
	req, err := a.newRequest(http.MethodPost, fmt.Sprintf("%s/organizations/%s/environments/%s/applications/%s/subscriptions/?plan=%s", a.GetConfig().GetURL(), a.OrgId, a.EnvId, appId, planId),
		WithHeader("Content-Type", "application/json"),
		WithToken(a.GetConfig().Auth.GetToken()),
	).Execute()
	if err != nil {
		return nil, err
	}

	subs := models.Subscriptions{}
	err = json.Unmarshal(req.Body, &subs)
	if err != nil {
		return nil, err
	}
	return &subs, err
}

// GetAPIKey - Request to get the api key assigned to your subscription
func (a *GraviteeClient) GetAPIKey(subsId, appId string) ([]models.AppCredentials, error) {
	req, err := a.newRequest(http.MethodGet, fmt.Sprintf("%s/organizations/%s/environments/%s/applications/%s/subscriptions/%s/apikeys", a.GetConfig().GetURL(), a.OrgId, a.EnvId, appId, subsId),
		WithHeader("Content-Type", "application/json"),
		WithToken(a.GetConfig().Auth.GetToken()),
	).Execute()

	if err != nil {
		return nil, err
	}

	apikey := []models.AppCredentials{}
	_ = json.Unmarshal(req.Body, &apikey)

	return apikey, err
}

// GetAPIKey - Request to get the api key assigned to your subscription
func (a *GraviteeClient) RemoveAPIKey(appId, subsId, apikeyId string) error {
	req, err := a.newRequest(http.MethodDelete, fmt.Sprintf("%s/organizations/%s/environments/%s/applications/%s/subscriptions/%s/apikeys/%s", a.GetConfig().GetURL(), a.OrgId, a.EnvId, appId, subsId, apikeyId),
		WithHeader("Content-Type", "application/json"),
		WithToken(a.GetConfig().Auth.GetToken()),
	).Execute()

	if err != nil {
		return err
	}
	if req.Code != http.StatusOK {
		return fmt.Errorf("received an unexpected response code %d from Gravitee when deleting the apikey", req.Code)
	}
	return err
}

// GetSpecificSubscription - Request to get the API's subscription by its identifier
func (a *GraviteeClient) GetSpecificSubscription(subsId, appId string) (*models.Subscriptions, error) {
	req, err := a.newRequest(http.MethodGet, fmt.Sprintf("%s/organizations/%s/environments/%s/applications/%s/subscriptions/%s", a.GetConfig().GetURL(), a.OrgId, a.EnvId, appId, subsId),
		WithHeader("Content-Type", "application/json"),
		WithToken(a.GetConfig().Auth.GetToken()),
	).Execute()
	if err != nil {
		return nil, err
	}

	subs := models.Subscriptions{}
	err = json.Unmarshal(req.Body, &subs)
	if err != nil {
		return nil, err
	}
	return &subs, err
}

// GetSubscriptions - Request to list subscriptions for a given API
func (a *GraviteeClient) GetSubscriptions(appid string) ([]models.Subscriptions, error) {
	req, err := a.newRequest(http.MethodGet, fmt.Sprintf("%s/organizations/%s/environments/%s/applications/%s/subscriptions", a.GetConfig().GetURL(), a.OrgId, a.EnvId, appid),
		WithHeader("Content-Type", "application/json"),
		WithToken(a.GetConfig().Auth.GetToken()),
	).Execute()
	if err != nil {
		return nil, err
	}

	logrus.Debug(string(req.Body))

	var subs AllSubs
	_ = json.Unmarshal(req.Body, &subs)
	logrus.Debug("le subs", subs)
	logrus.Debug("le subs.Subs ", subs.Subs)
	for _, s := range subs.Subs {
		logrus.Debug("un sub", s)
	}
	return subs.Subs, err
}

// CreateAppCredentials - Request to create an application
func (a *GraviteeClient) CreateApp(appli *models.App) (*models.App, error) {
	body, _ := json.Marshal(appli)

	req, err := a.newRequest(http.MethodPost, fmt.Sprintf("%s/organizations/%s/environments/%s/applications", a.GetConfig().GetURL(), a.OrgId, a.EnvId),
		WithHeader("Content-Type", "application/json"),
		WithToken(a.GetConfig().Auth.GetToken()),
		WithBody(body),
	).Execute()

	if err != nil {
		return nil, err
	}

	if req.Code != http.StatusCreated {
		return nil, fmt.Errorf("received an unexpected response code %d from Gravitee when creating the app", req.Code)
	}

	App := models.App{}
	err = json.Unmarshal(req.Body, &App)
	if err != nil {
		return nil, err
	}
	return &App, err
}

func (a *GraviteeClient) RemoveApp(appId string) error {
	response, err := a.newRequest(http.MethodDelete, fmt.Sprintf("%s/organizations/%s/environments/%s/applications/%s", a.GetConfig().GetURL(), a.OrgId, a.EnvId, appId),
		WithHeader("Content-Type", "application/json"),
		WithToken(a.GetConfig().Auth.GetToken()),
	).Execute()
	if err != nil {
		return err
	}
	if response.Code != http.StatusOK {
		return fmt.Errorf("received an unexpected response code %d from Gravitee when deleting the app", response.Code)
	}
	return err
}

func (a *GraviteeClient) UpdateCredential(appId, subId string) (*models.AppCredentials, error) {
	req, err := a.newRequest(http.MethodPost, fmt.Sprintf("%s/organizations/%s/environments/%s/applications/%s/subscriptions/%s/apikeys/_renew", a.GetConfig().GetURL(), a.OrgId, a.EnvId, appId, subId),
		WithHeader("Content-Type", "application/json"),
		WithToken(a.GetConfig().Auth.GetToken()),
	).Execute()

	if err != nil {
		return nil, err
	}

	newkey := models.AppCredentials{}
	err = json.Unmarshal(req.Body, &newkey)
	if err != nil {
		return nil, err
	}
	return &newkey, err
}

func (a *GraviteeClient) ListAPIsPlans(apiId string) ([]models.Plan, error) {
	response, err := a.newRequest(http.MethodGet, fmt.Sprintf("%s/v2/organizations/%s/environments/%s/apis/%s/plans", a.GetConfig().GetURL(), a.OrgId, a.EnvId, apiId),
		WithHeader("Content-Type", "application/json"),
		WithToken(a.GetConfig().Auth.GetToken()),
	).Execute()

	if err != nil {
		return nil, err
	}

	var plans AllPlans
	err = json.Unmarshal(response.Body, &plans)
	if err != nil {
		return nil, err
	}
	return plans.Plans, err
}

func (a *GraviteeClient) PublishaPlan(apiId, planId string) error {
	response, err := a.newRequest(http.MethodPost, fmt.Sprintf("%s/v2/organizations/%s/environments/%s/apis/%s/plans/%s/_publish", a.GetConfig().GetURL(), a.OrgId, a.EnvId, apiId, planId),
		WithHeader("Content-Type", "application/json"),
		WithToken(a.GetConfig().Auth.GetToken()),
	).Execute()

	if err != nil {
		return err
	}

	if response.Code != http.StatusOK {
		return fmt.Errorf("received an unexpected response code %d from Gravitee when deleting the app", response.Code)
	}
	return err

}

func (a *GraviteeClient) TransferSubs(apiId, subId, newPlanId string) (*models.Subscriptions, error) {
	body, _ := json.Marshal(newPlanId)
	response, err := a.newRequest(http.MethodPost, fmt.Sprintf("%s/v2/organizations/%s/environments/%s/apis/%s/subscriptions/%s/_transfer", a.GetConfig().URL, a.OrgId, a.EnvId, apiId, subId),
		WithHeader("Content-Type", "application/json"),
		WithToken(a.GetConfig().Auth.GetToken()),
		WithBody(body),
	).Execute()

	if err != nil {
		return nil, err
	}

	newsub := models.Subscriptions{}
	err = json.Unmarshal(response.Body, &newsub)
	if err != nil {
		return nil, err
	}
	return &newsub, err
}

func (a *GraviteeClient) CreatePlan(apiId string, plan *models.Plan) (*models.Plan, error) {
	body, _ := json.Marshal(plan)
	logrus.Debug("le plan : ", plan)
	logrus.Debug("le plan en json", string(body))
	post, err := a.newRequest(http.MethodPost, fmt.Sprintf("%s/v2/organizations/%s/environments/%s/apis/%s/plans", a.GetConfig().GetURL(), a.OrgId, a.EnvId, apiId),
		WithHeader("Content-Type", "application/json"),
		WithToken(a.GetConfig().Auth.GetToken()),
		WithBody(body),
	).Execute()

	if err != nil {
		return nil, err
	}

	newplan := models.Plan{}
	err = json.Unmarshal(post.Body, &newplan)
	if err != nil {
		return nil, err
	}
	return &newplan, err

}
