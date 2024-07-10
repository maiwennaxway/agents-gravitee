package gravitee

import (
	"fmt"
	"testing"

	v1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
	management "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/management/v1alpha1"
	defs "github.com/Axway/agent-sdk/pkg/apic/definitions"
	"github.com/Axway/agent-sdk/pkg/apic/provisioning"
	"github.com/Axway/agent-sdk/pkg/apic/provisioning/mock"
	"github.com/Axway/agent-sdk/pkg/util"
	"github.com/maiwennaxway/agents-gravitee/client/pkg/gravitee/models"
	"github.com/stretchr/testify/assert"
)

func TestAccessRequestDeprovision(t *testing.T) {
	tests := []struct {
		name        string
		status      provisioning.Status
		appId       string
		apiID       string
		getAppErr   error
		upCredErr   error
		missingCred bool
	}{
		{
			name:   "should deprovision an access request",
			appId:  "app-one",
			apiID:  "abc-123",
			status: provisioning.Success,
		},
		{
			name:      "should return success when the developer app is already removed",
			appId:     "app-one",
			apiID:     "abc-123",
			status:    provisioning.Success,
			getAppErr: fmt.Errorf("404"),
		},
		{
			name:      "should fail to deprovision an access request when retrieving the app, and the error is not a 404",
			appId:     "app-one",
			apiID:     "abc-123",
			getAppErr: fmt.Errorf("error"),
			status:    provisioning.Error,
		},
		{
			name:      "should fail to deprovision an access request when revoking the credential",
			appId:     "app-one",
			apiID:     "abc-123",
			status:    provisioning.Error,
			upCredErr: fmt.Errorf("error"),
		},
		{
			name:   "should return an error when the appId is not found",
			appId:  "",
			apiID:  "abc-123",
			status: provisioning.Error,
		},
		{
			name:   "should return an error when the apiID is not found",
			appId:  "app-one",
			apiID:  "",
			status: provisioning.Error,
		},
		{
			name:        "should succeed if no credentials are found, nothing to do",
			appId:       "app-one",
			apiID:       "api-123",
			status:      provisioning.Success,
			missingCred: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			app := newApp(fmt.Sprintf("%s-no-quota", tc.apiID), tc.appId)

			p := NewProvisioner(&mockClient{
				t:         t,
				getAppErr: tc.getAppErr,
				upCredErr: tc.upCredErr,
				app:       app,
				appId:     tc.appId,
				key:       app.Credentials[0].ApiKey,
			}, 30, &mockCache{t: t}, false)

			if tc.missingCred {
				app.Credentials = nil
			}

			mar := mock.MockAccessRequest{
				InstanceDetails: map[string]interface{}{
					defs.AttrExternalAPIID: tc.apiID,
				},
				AppDetails: nil,
				AppName:    tc.appId,
			}

			status := p.AccessRequestDeprovision(&mar)
			assert.Equal(t, tc.status.String(), status.GetStatus().String())
			assert.Equal(t, 0, len(status.GetProperties()))
		})
	}
}

func TestAccessRequestProvision(t *testing.T) {
	tests := []struct {
		name         string
		status       provisioning.Status
		appId        string
		apiID        string
		newAPIID     string
		apiStage     string
		getAppErr    error
		addCredErr   error
		addProdErr   error
		upCredErr    error
		existingProd bool
		noCreds      bool
		isApiLinked  bool
	}{
		{
			name:     "should provision an access request",
			appId:    "app-one",
			apiID:    "abc-123",
			newAPIID: "abc-123-no-quota",
			apiStage: "prod",
			status:   provisioning.Success,
		},
		{
			name:     "should provision an access request when there are no credentials on the app",
			appId:    "app-one",
			apiID:    "abc-123",
			newAPIID: "abc-123-no-quota",
			apiStage: "prod",
			status:   provisioning.Success,
			noCreds:  true,
		},
		{
			name:        "should provision an access request when the api is already linked to a credential",
			appId:       "app-one",
			apiID:       "abc-123",
			newAPIID:    "abc-123-no-quota",
			apiStage:    "prod",
			status:      provisioning.Success,
			isApiLinked: true,
		},
		{
			name:        "should fail to deprovision an access request when the api is already linked but could not be enabled",
			appId:       "app-one",
			apiID:       "abc-123",
			newAPIID:    "abc-123-no-quota",
			apiStage:    "prod",
			status:      provisioning.Success,
			upCredErr:   fmt.Errorf("error"),
			isApiLinked: true,
		},
		{
			name:       "should fail to deprovision an access request",
			appId:      "app-one",
			apiID:      "abc-123",
			newAPIID:   "abc-123-no-quota",
			apiStage:   "prod",
			status:     provisioning.Error,
			addCredErr: fmt.Errorf("error"),
		},
		{
			name:      "should fail to deprovision when unable to retrieve the app",
			appId:     "app-one",
			apiID:     "abc-123",
			newAPIID:  "abc-123-no-quota",
			apiStage:  "prod",
			status:    provisioning.Error,
			getAppErr: fmt.Errorf("error"),
		},
		{
			name:     "should return an error when the apiID is not found",
			appId:    "app-one",
			apiID:    "",
			apiStage: "prod",
			status:   provisioning.Error,
		},
		{
			name:     "should return an error when the appId is not found",
			appId:    "",
			apiID:    "abc-123",
			apiStage: "prod",
			status:   provisioning.Error,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			apiID := ""
			if tc.isApiLinked {
				apiID = tc.apiID
			}
			app := newApp(apiID, tc.appId)

			key := app.Credentials[0].ApiKey

			if tc.noCreds {
				app.Credentials = nil
			}

			p := NewProvisioner(&mockClient{
				addCredErr: tc.addCredErr,
				app:        app,
				appId:      tc.appId,
				key:        key,
				getAppErr:  tc.getAppErr,
				upCredErr:  tc.upCredErr,
				t:          t,
			}, 30, &mockCache{t: t}, false)

			mar := mock.MockAccessRequest{
				InstanceDetails: map[string]interface{}{
					defs.AttrExternalAPIID:    tc.apiID,
					defs.AttrExternalAPIStage: tc.apiStage,
				},
				AppDetails: nil,
				AppName:    tc.appId,
			}

			status, _ := p.AccessRequestProvision(&mar)
			assert.Equal(t, tc.status.String(), status.GetStatus().String())
			if tc.status == provisioning.Success {
				assert.Equal(t, 1, len(status.GetProperties()))
			} else {
				assert.Equal(t, 0, len(status.GetProperties()))
			}
		})
	}
}

func TestApplicationRequestDeprovision(t *testing.T) {
	tests := []struct {
		name     string
		status   provisioning.Status
		appId    string
		apiID    string
		rmAppErr error
	}{
		{
			name:   "should deprovision an application",
			status: provisioning.Success,
			appId:  "app-one",
			apiID:  "api-123",
		},
		{
			name:   "should return an error when the app name is not found",
			status: provisioning.Error,
			appId:  "",
			apiID:  "api-123",
		},
		{
			name:   "should return an error when the app name is not found",
			status: provisioning.Error,
			appId:  "",
			apiID:  "api-123",
		},
		{
			name:     "should return an error failing to remove the app",
			status:   provisioning.Error,
			appId:    "app-one",
			apiID:    "api-123",
			rmAppErr: fmt.Errorf("err"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			app := newApp(tc.apiID, tc.appId)

			p := NewProvisioner(&mockClient{
				app:      app,
				appId:    tc.appId,
				key:      "key",
				t:        t,
				rmAppErr: tc.rmAppErr,
			}, 30, &mockCache{t: t}, false)

			mar := mock.MockApplicationRequest{
				AppName:  tc.appId,
				TeamName: "team-one",
			}

			status := p.ApplicationRequestDeprovision(&mar)
			assert.Equal(t, tc.status.String(), status.GetStatus().String())
			assert.Equal(t, 0, len(status.GetProperties()))
		})
	}
}

func TestApplicationRequestProvision(t *testing.T) {
	tests := []struct {
		name         string
		status       provisioning.Status
		appId        string
		apiID        string
		createAppErr error
	}{
		{
			name:   "should provision an application",
			status: provisioning.Success,
			appId:  "app-one",
			apiID:  "api-123",
		},
		{
			name:         "should return an error when creating the app",
			status:       provisioning.Error,
			appId:        "app-one",
			apiID:        "api-123",
			createAppErr: fmt.Errorf("err"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			app := newApp(tc.apiID, tc.appId)

			p := NewProvisioner(&mockClient{
				app:          app,
				appId:        tc.appId,
				t:            t,
				createAppErr: tc.createAppErr,
			}, 30, &mockCache{t: t}, false)

			mar := mock.MockApplicationRequest{
				AppName:  tc.appId,
				TeamName: "team-one",
			}

			status := p.ApplicationRequestProvision(&mar)
			assert.Equal(t, tc.status.String(), status.GetStatus().String())
			assert.Equal(t, 0, len(status.GetProperties()))
		})
	}
}

func TestCredentialDeprovision(t *testing.T) {
	tests := []struct {
		name      string
		status    provisioning.Status
		appId     string
		apiID     string
		getAppErr error
		credType  string
	}{
		{
			name:     "should deprovision an api-key credential",
			status:   provisioning.Success,
			appId:    "app-one",
			apiID:    "api-123",
			credType: "api-key",
		},
		{
			name:     "should deprovision an oauth credential",
			status:   provisioning.Success,
			appId:    "app-one",
			apiID:    "api-123",
			credType: "oauth",
		},
		{
			name:      "should return success when unable to retrieve the app",
			status:    provisioning.Success,
			appId:     "app-one",
			apiID:     "api-123",
			getAppErr: fmt.Errorf("err"),
			credType:  "oauth",
		},
		{
			name:   "should return success when credential not on app",
			status: provisioning.Success,
			appId:  "app-one",
			apiID:  "api-123",
		},
		{
			name:     "should return an error when the app name is not found",
			status:   provisioning.Error,
			appId:    "",
			apiID:    "api-123",
			credType: "oauth",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			app := newApp(tc.apiID, tc.appId)

			key := "consumer-key"
			if tc.credType == "" {
				key = ""
			}
			p := NewProvisioner(&mockClient{
				app:       app,
				appId:     tc.appId,
				key:       key,
				t:         t,
				getAppErr: tc.getAppErr,
			}, 30, &mockCache{t: t, appId: tc.appId}, false)

			thisHash, _ := util.ComputeHash(key)
			mcr := mock.MockCredentialRequest{
				AppName:     tc.appId,
				CredDefName: tc.credType,
				Details: map[string]string{
					appRefName: tc.appId,
					credRefKey: fmt.Sprintf("%v", thisHash),
				},
			}

			status := p.CredentialDeprovision(&mcr)
			assert.Equal(t, tc.status.String(), status.GetStatus().String())
		})
	}
}

func TestCredentialProvision(t *testing.T) {
	tests := []struct {
		name      string
		status    provisioning.Status
		appId     string
		apiID     string
		getAppErr error
		credType  string
	}{
		{
			name:     "should provision an api-key credential",
			status:   provisioning.Success,
			appId:    "app-one",
			apiID:    "api-123",
			credType: "api-key",
		},
		{
			name:     "should provision an oauth credential",
			status:   provisioning.Success,
			appId:    "app-one",
			apiID:    "api-123",
			credType: "oauth",
		},
		{
			name:     "should return an error when the app name is not found",
			status:   provisioning.Error,
			appId:    "",
			apiID:    "api-123",
			credType: "oauth",
		},
		{
			name:      "should return an error when unable to retrieve the app",
			status:    provisioning.Error,
			appId:     "app-one",
			apiID:     "api-123",
			getAppErr: fmt.Errorf("err"),
			credType:  "oauth",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			app := newApp(tc.apiID, tc.appId)

			p := NewProvisioner(&mockClient{
				app:       app,
				appId:     tc.appId,
				t:         t,
				getAppErr: tc.getAppErr,
			}, 30, &mockCache{t: t, appId: tc.appId}, false)

			mcr := mock.MockCredentialRequest{
				AppName:     tc.appId,
				CredDefName: tc.credType,
			}

			status, cred := p.CredentialProvision(&mcr)
			if tc.status == provisioning.Error {
				assert.Nil(t, cred)
				assert.Equal(t, 0, len(status.GetProperties()))
			} else {
				assert.NotNil(t, cred)
				assert.Equal(t, 2, len(status.GetProperties()))
				if tc.credType == "oauth" {
					assert.Contains(t, cred.GetData(), provisioning.OauthClientID)
					assert.Contains(t, cred.GetData(), provisioning.OauthClientSecret)
					assert.NotContains(t, cred.GetData(), provisioning.APIKey)
				} else {
					assert.NotContains(t, cred.GetData(), provisioning.OauthClientID)
					assert.NotContains(t, cred.GetData(), provisioning.OauthClientSecret)
					assert.Contains(t, cred.GetData(), provisioning.APIKey)
				}
			}
			assert.Equal(t, tc.status.String(), status.GetStatus().String())
		})
	}
}

func TestCredentialUpdate(t *testing.T) {
	tests := []struct {
		name           string
		status         provisioning.Status
		appId          string
		apiID          string
		getAppErr      error
		credType       string
		skipAppDetail  bool
		skipCredDetail bool
		action         provisioning.CredentialAction
	}{
		{
			name:     "should revoke credential",
			status:   provisioning.Success,
			appId:    "app-one",
			apiID:    "api-123",
			credType: "api-key",
			action:   provisioning.Suspend,
		},
		{
			name:     "should enable credential",
			status:   provisioning.Success,
			appId:    "app-one",
			apiID:    "api-123",
			credType: "api-key",
			action:   provisioning.Enable,
		},
		{
			name:          "should return an error when unable to get appId detail",
			status:        provisioning.Error,
			appId:         "app-one",
			skipAppDetail: true,
			apiID:         "api-123",
		},
		{
			name:           "should return an error when unable to get cred ref detail",
			status:         provisioning.Error,
			appId:          "app-one",
			skipCredDetail: true,
			apiID:          "api-123",
		},
		{
			name:      "should return an error when unable to retrieve the app",
			status:    provisioning.Error,
			appId:     "app-one",
			apiID:     "api-123",
			getAppErr: fmt.Errorf("err"),
		},
		{
			name:     "should return an error when credential action unknown",
			status:   provisioning.Error,
			appId:    "app-one",
			apiID:    "api-123",
			credType: "api-key",
			action:   provisioning.Rotate,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			app := newApp(tc.apiID, tc.appId)

			key := "consumer-key"
			if tc.credType == "" {
				key = ""
			}
			p := NewProvisioner(&mockClient{
				app:       app,
				appId:     tc.appId,
				key:       key,
				t:         t,
				getAppErr: tc.getAppErr,
			}, 30, &mockCache{t: t, appId: tc.appId}, false)

			thisHash, _ := util.ComputeHash(key)
			details := map[string]string{
				appRefName: tc.appId,
				credRefKey: fmt.Sprintf("%v", thisHash),
			}
			if tc.skipAppDetail {
				delete(details, appRefName)
			}
			if tc.skipCredDetail {
				delete(details, credRefKey)
			}

			mcr := mock.MockCredentialRequest{
				AppName:     tc.appId,
				CredDefName: tc.credType,
				Details:     details,
				Action:      tc.action,
			}

			status, _ := p.CredentialUpdate(&mcr)
			assert.Equal(t, tc.status.String(), status.GetStatus().String())
		})
	}
}

type mockClient struct {
	addCredErr   error
	app          *models.App
	appId        string
	createAppErr error
	getAppErr    error
	key          string
	subId        string
	apikey       string
	apiId        string
	planId       string
	rmAppErr     error
	rmCredErr    error
	upCredErr    error
	getPlanErr   error
	TransferErr  error
	createSubErr error

	t *testing.T
}

func (m mockClient) CreateApp(newApp *models.App) (*models.App, error) {
	assert.Equal(m.t, m.appId, newApp.Id)
	return &models.App{
		Credentials: []models.AppCredentials{
			{
				Id:     m.key,
				ApiKey: "secret",
			},
		},
	}, m.createAppErr
}

func (m mockClient) RemoveApp(id string) error {
	assert.Equal(m.t, m.appId, id)
	return m.rmAppErr
}

func (m mockClient) GetApp(id string) (*models.App, error) {
	assert.Equal(m.t, m.appId, id)
	return nil, m.getAppErr
}

func (m mockClient) GetApps() ([]models.App, error) {
	return nil, m.getAppErr
}

func (m mockClient) CreatePlan(apiId string, plan *models.Plan) (*models.Plan, error) {
	assert.Equal(m.t, m.apiId, apiId)
	return &models.Plan{
		ApiId: apiId,
		Name:  "test",
	}, m.createAppErr
}

func (m mockClient) ListAPIsPlans(apiId string) ([]models.Plan, error) {
	assert.Equal(m.t, m.apiId, apiId)
	return nil, m.getPlanErr
}

func (m mockClient) PublishaPlan(apiId, planId string) error {
	assert.Equal(m.t, m.apiId, apiId)
	assert.Equal(m.t, m.planId, planId)
	return m.getPlanErr
}

func (m mockClient) TransferSubs(apiId, subId, newPlanId string) (*models.Subscriptions, error) {
	assert.Equal(m.t, m.apiId, apiId)
	assert.Equal(m.t, m.subId, subId)
	assert.Equal(m.t, m.planId, newPlanId)
	return nil, m.TransferErr
}

func (m mockClient) SubscribetoAnAPI(appId, planId string) (*models.Subscriptions, error) {
	assert.Equal(m.t, m.appId, appId)
	assert.Equal(m.t, m.planId, planId)
	return nil, m.createSubErr
}

func (m mockClient) GetAPIKey(subsId, appId string) (*models.AppCredentials, error) {
	assert.Equal(m.t, m.appId, appId)
	assert.Equal(m.t, m.subId, subsId)
	return nil, nil
}

func (m mockClient) RemoveAPIKey(appId, subsId, key string) error {
	assert.Equal(m.t, m.appId, appId)
	assert.Equal(m.t, m.subId, subsId)
	assert.Equal(m.t, m.apikey, key)
	return m.rmCredErr
}

func (m mockClient) UpdateCredential(subId, appId string) (*models.AppCredentials, error) {
	assert.Equal(m.t, m.appId, appId)
	assert.Equal(m.t, m.subId, subId)
	return nil, nil
}

func (m mockClient) GetApi(apiId string) (*models.Api, error) {
	return &models.Api{
		Id: m.apiId,
	}, nil
}

func (m mockClient) DeployApi(apiId string) error {
	assert.Equal(m.t, m.apiId, apiId)
	return nil
}
func newApp(apiId string, appId string) *models.App {
	cred := &models.App{
		Credentials: []models.AppCredentials{
			{
				Subscriptions: nil,
				Id:            "consumer-key",
				ApiKey:        "consumer-secret",
			},
		},
		Name: appId,
		Id:   apiId,
	}

	return cred
}

type mockCache struct {
	t       *testing.T
	appId   string
	apiName string
}

func (m *mockCache) GetAccessRequestsByApp(managedappId string) []*v1.ResourceInstance {
	assert.Equal(m.t, m.appId, managedappId)
	ar1 := management.NewAccessRequest("ar1", "env")
	ar1.Spec.ManagedApplication = m.appId
	util.SetAgentDetailsKey(ar1, "apiKey", "api")
	ri, _ := ar1.AsInstance()
	return []*v1.ResourceInstance{ri}
}

func (m *mockCache) GetAPIServiceInstanceByName(apiName string) (*v1.ResourceInstance, error) {
	assert.Equal(m.t, m.apiName, apiName)
	apisi := management.NewAPIServiceInstance(apiName, "env")
	util.SetAgentDetailsKey(apisi, defs.AttrExternalAPIID, "apiName")
	ri, _ := apisi.AsInstance()
	return ri, nil
}
