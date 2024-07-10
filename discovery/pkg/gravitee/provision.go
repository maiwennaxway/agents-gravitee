package gravitee

import (
	"fmt"
	"strings"
	"time"

	v1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
	defs "github.com/Axway/agent-sdk/pkg/apic/definitions"
	prov "github.com/Axway/agent-sdk/pkg/apic/provisioning"
	"github.com/Axway/agent-sdk/pkg/util"
	"github.com/Axway/agent-sdk/pkg/util/log"
	"github.com/maiwennaxway/agents-gravitee/client/pkg/gravitee/models"
)

const (
	credRefKey = "credentialReference"
	appRefName = "appName"
	planRef    = "plan-name"
)

type provisioner struct {
	client                client
	credExpDays           int
	cacheManager          cacheManager
	shouldCloneAttributes bool
	logger                log.FieldLogger
}

type client interface {
	GetApp(id string) (*models.App, error)
	CreateApp(appli *models.App) (*models.App, error)
	RemoveApp(id string) error
	GetApps() ([]models.App, error)
	GetApi(apiId string) (*models.Api, error)
	DeployApi(apiID string) error
	GetAPIKey(subsId, appId string) (*models.AppCredentials, error)
	SubscribetoAnAPI(appId, planId string) (*models.Subscriptions, error)
	//GetAppCredentials(appId string) (*models.App, error)
	UpdateCredential(subId, appId string) (*models.AppCredentials, error)
	RemoveAPIKey(appId, subsId, apikeyId string) error
	ListAPIsPlans(apiId string) ([]models.Plan, error)
	TransferSubs(apiId, subId, newPlanId string) (*models.Subscriptions, error)
	CreatePlan(apiId string, plan *models.Plan) (*models.Plan, error)
	PublishaPlan(apiId, planId string) error
}

type cacheManager interface {
	GetAccessRequestsByApp(managedAppName string) []*v1.ResourceInstance
	GetAPIServiceInstanceByName(apiName string) (*v1.ResourceInstance, error)
}

// NewProvisioner creates a type to implement the SDK Provisioning methods for handling subscriptions
func NewProvisioner(client client, credExpDays int, cacheMan cacheManager, cloneAttributes bool) prov.Provisioning {
	return &provisioner{
		client:                client,
		credExpDays:           credExpDays,
		cacheManager:          cacheMan,
		shouldCloneAttributes: cloneAttributes,
		logger:                log.NewFieldLogger().WithComponent("provision").WithPackage("gravitee"),
	}
}

// AccessRequestDeprovision - removes an api from an application
func (p provisioner) AccessRequestDeprovision(req prov.AccessRequest) prov.RequestStatus {
	instDetails := req.GetInstanceDetails()
	apiID := util.ToString(instDetails[defs.AttrExternalAPIID])
	logger := p.logger.WithField("handler", "AccessRequestDeprovision").WithField("apiID", apiID).WithField("application", req.GetApplicationName())

	//apiName := getAPIName(apiID, req.GetQuota())
	// remove link between api and app
	logger.Info("deprovisioning access request")
	ps := prov.NewRequestStatusBuilder()

	appName := req.GetApplicationName()
	if appName == "" {
		return failed(logger, ps, fmt.Errorf("application name not found"))
	}
	appId, _ := p.FindAppIdbyname(appName)
	app, err := p.client.GetApp(appId)
	if err != nil {
		if ok := strings.Contains(err.Error(), "404"); ok {
			return ps.Success()
		}

		return failed(logger, ps, fmt.Errorf("failed to retrieve app: %s", err))
	}

	if apiID == "" {
		return failed(logger, ps, fmt.Errorf("%s not found", defs.AttrExternalAPIID))
	}

	var cred *models.AppCredentials
	// find the credential that the api is linked to
	for _, c := range app.Credentials {
		for _, sub := range c.Subscriptions {
			if sub.Api.Id == apiID {
				cred = &c

				err = p.client.RemoveAPIKey(apiID, sub.Id, cred.Id)
				if err != nil {
					return failed(logger, ps, fmt.Errorf("failed to revoke api %s from credential: %s", sub.Api.Id, err))
				}
			}
		}
	}

	// Ensure that cred is used after the loop
	if cred != nil {
		// For example, log or perform some other actions
		fmt.Println("Credential found and updated:", cred)
	} /*else {
		return failed(logger, ps, fmt.Errorf("no credential found for api %s", apiID))
	}*/

	logger.Info("removed access")

	return ps.Success()
}

// AccessRequestProvision - adds an api to an application
func (p provisioner) AccessRequestProvision(req prov.AccessRequest) (prov.RequestStatus, prov.AccessData) {
	instDetails := req.GetInstanceDetails()
	apiID := util.ToString(instDetails[defs.AttrExternalAPIID])
	logger := p.logger.WithField("handler", "AccessRequestProvision").WithField("apiID", apiID).WithField("application", req.GetApplicationName())

	logger.Info("processing access request")
	ps := prov.NewRequestStatusBuilder()

	if apiID == "" {
		return failed(logger, ps, fmt.Errorf("%s name not found", defs.AttrExternalAPIID)), nil
	}

	appName := req.GetApplicationName()
	if appName == "" {
		return failed(logger, ps, fmt.Errorf("application name not found")), nil
	}

	// get plan name from access request

	// get api, or create new one
	quota := 1
	quotaInterval := 1
	quotaTimeUnit := ""

	if q := req.GetQuota(); q != nil {
		quota = int(q.GetLimit())

		switch q.GetInterval() {
		case prov.Daily:
			quotaTimeUnit = "day"
		case prov.Weekly:
			quotaTimeUnit = "day"
			quotaInterval = 7
		case prov.Monthly:
			quotaTimeUnit = "month"
		case prov.Annually:
			quotaTimeUnit = "month"
			quotaInterval = 12
		default:
			return failed(logger, ps, fmt.Errorf("invalid quota time unit: received %s", q.GetIntervalString())), nil
		}
	}

	planName := "Unlimited"
	if req.GetQuota() != nil {
		planName = req.GetQuota().GetPlanName()
	}

	var plan *models.Plan
	var err error
	logger.Debug("handling creation plan")
	plan, err = p.CreatePlan(logger, planName, apiID, quotaTimeUnit, quota, quotaInterval)
	if err != nil {
		return failed(logger, ps, fmt.Errorf("failed to create api : %s", err)), nil
	}
	appId, _ := p.FindAppIdbyname(appName)
	app, err := p.client.GetApp(appId)
	if err != nil {
		return failed(logger, ps, fmt.Errorf("failed to retrieve app %s: %s", appName, err)), nil
	}

	// add api to credentials that are not associated with it
	_, err = p.client.SubscribetoAnAPI(app.Id, plan.Id)
	if err != nil {
		return failed(logger, ps, fmt.Errorf("failed to subscribe app %s to plan %s due to %s", appName, plan.Name, err)), nil
	}

	logger.Info("granted access")

	return ps.AddProperty(planRef, plan.Name).Success(), nil
}

func (p provisioner) CreatePlan(logger log.FieldLogger, Planname, ApiId, quotaTimeUnit string, quota, quotaInterval int) (*models.Plan, error) {
	// only create a plan if one is not fou
	//a changer dans les plans :
	Plans, _ := p.client.ListAPIsPlans(ApiId)
	for _, p := range Plans {
		if p.Name == Planname {
			return &p, nil
		}
	}
	Planbody := &models.Plan{}
	if quota == 1 && quotaInterval == 1 && quotaTimeUnit == "" {
		logger.Debug("je passe par la")
		Planbody = &models.Plan{
			DefinitionVersion: "V2",
			Description:       Planname,
			Flows: []models.Flows{
				{
					PathOperator: models.PathOperator{
						Operator: "STARTS_WITH",
					},
				},
			},
			Name:       Planname,
			Security:   "API_KEY",
			Validation: "AUTO",
		}
	} else {
		logger.Debug("je passe par la avec quotas")
		Planbody = &models.Plan{
			DefinitionVersion: "V2",
			Description:       Planname,
			Flows: []models.Flows{
				{
					PathOperator: models.PathOperator{
						Operator: "STARTS_WITH",
					},
					Pre: []models.Pre{
						{
							Quota: models.Quota{
								Limit:          quota,
								PeriodTime:     quotaInterval,
								PeriodTimeUnit: quotaTimeUnit,
							},
						},
					},
				},
			},
			Name:       Planname,
			Security:   "API_KEY",
			Validation: "AUTO",
		}
	}
	logger.Debug("je crée le plan")
	Plan, err := p.client.CreatePlan(ApiId, Planbody)
	if err != nil {
		logger.Error("error creating the plan", err)
		return nil, err
	}
	logger.Debug("je publie le plan")
	erreur := p.client.PublishaPlan(ApiId, Plan.Id)
	if erreur != nil {
		return nil, erreur
	}
	_ = p.client.DeployApi(ApiId)
	//update le plan avec le nouveau quota mais en fait on créer un nouveau plan avec ses quota la...
	logger.Debug("creating plan", Plan.Id)

	return Plan, err
}

// FindAppIdbyname - find the id of the application by the given name
func (p provisioner) FindAppIdbyname(name string) (string, error) {
	apps, err := p.client.GetApps()
	for _, app := range apps {
		if app.Name == name {
			return app.Id, nil
		}
	}
	return "", err
}

// ApplicationRequestDeprovision - removes an app from gravitee
func (p provisioner) ApplicationRequestDeprovision(req prov.ApplicationRequest) prov.RequestStatus {
	logger := p.logger.WithField("handler", "ApplicationRequestDeprovision").WithField("application", req.GetManagedApplicationName())

	logger.Info("removing app")
	ps := prov.NewRequestStatusBuilder()

	appName := req.GetManagedApplicationName()
	if appName == "" {
		return failed(logger, ps, fmt.Errorf("managed application %s not found", appName))
	}

	id, _ := p.FindAppIdbyname(appName)

	err := p.client.RemoveApp(id)
	if err != nil {
		return failed(logger, ps, fmt.Errorf("failed to delete app: %s", err))
	}

	logger.Info("removed app")

	return ps.Success()
}

// ApplicationRequestProvision - creates an gravitee app
func (p provisioner) ApplicationRequestProvision(req prov.ApplicationRequest) prov.RequestStatus {
	logger := p.logger.WithField("handler", "ApplicationRequestProvision").WithField("application", req.GetManagedApplicationName())

	logger.Info("provisioning app")
	ps := prov.NewRequestStatusBuilder()
	apps, err := p.client.GetApps()
	for _, app := range apps {
		if app.Name == req.GetManagedApplicationName() {
			_, _ = p.client.GetApp(app.Id)
			logger.Debug("je sors de ApplicationRequestProvision car l'app est deja existante")
			return ps.Success()
		}

	}
	logger.Debug("je sors de la boucle car l'app n'est pas deja existante donc je dois la créer")
	app := models.App{
		Name:       req.GetManagedApplicationName(),
		Descripion: req.GetManagedApplicationName(),
	}
	_, _ = p.client.CreateApp(&app)

	if err != nil {
		return failed(logger, ps, fmt.Errorf("failed to create app: %s", err))
	}
	logger.Info("provisioned app")
	return ps.Success()
}

// CredentialDeprovision - Return success because there are no credentials to remove until the app is deleted
func (p provisioner) CredentialDeprovision(req prov.CredentialRequest) prov.RequestStatus {
	logger := p.logger.WithField("handler", "CredentialDeprovision").WithField("application", req.GetApplicationName())

	logger.Info("removing credential")
	ps := prov.NewRequestStatusBuilder()

	appName := req.GetCredentialDetailsValue(appRefName)
	if appName == "" {
		return failed(logger, ps, fmt.Errorf("application name not found"))
	}

	appId, _ := p.FindAppIdbyname(appName)

	app, err := p.client.GetApp(appId)
	if err != nil {
		logger.Trace("application had previously been removed")
		return ps.Success()
	}

	credKey := ""
	curHash := req.GetCredentialDetailsValue(credRefKey)
	if curHash == "" {
		return failed(logger, ps, fmt.Errorf("credential reference not found"))
	}
	for _, cred := range app.Credentials {
		thisHash, _ := util.ComputeHash(cred.ApiKey)
		if curHash == fmt.Sprintf("%v", thisHash) {
			credKey = cred.ApiKey
			break
		}
	}

	if credKey == "" {
		return ps.Success()
	}

	// remove the credential created by default for the application, the credential request will create a new one
	for _, cred := range app.Credentials {
		for _, sub := range cred.Subscriptions {
			err = p.client.RemoveAPIKey(app.Id, sub.Id, credKey)
			if err != nil {
				return failed(logger, ps, fmt.Errorf("unexpected error removing the credential"))
			}
		}
	}

	logger.Info("removed credential")
	return ps.Success()
}

// CredentialProvision - retrieves the app credentials for oauth or api key authentication
func (p provisioner) CredentialProvision(req prov.CredentialRequest) (prov.RequestStatus, prov.Credential) {
	logger := p.logger.WithField("handler", "CredentialProvision").WithField("application", req.GetApplicationName())

	logger.Info("provisioning credential")
	ps := prov.NewRequestStatusBuilder()

	appName := req.GetApplicationName()
	if appName == "" {
		return failed(logger, ps, fmt.Errorf("application name not found")), nil
	}

	appId, _ := p.FindAppIdbyname(appName)

	curApp, err := p.client.GetApp(appId)
	if err != nil {
		return failed(logger, ps, fmt.Errorf("error retrieving app: %s", err)), nil
	}

	// associate all apis
	accReqs := p.cacheManager.GetAccessRequestsByApp(appName)
	apis := []string{}
	for _, arInst := range accReqs {
		apiName, err := util.GetAgentDetailsValue(arInst, planRef)
		logger.Debug("apiName : ", apiName)
		if err == nil && apiName != "" {
			apis = append(apis, apiName)
		}
	}
	if len(apis) == 0 {
		return failed(logger, ps, fmt.Errorf("at least one api access is required for a credential")), nil
	}

	updateApp, err := p.client.CreateApp(curApp)
	if err != nil {
		return failed(logger, ps, fmt.Errorf("error creating app credential: %s", err)), nil
	}

	// find the new cred
	cred := models.AppCredentials{}
	keys := map[string]struct{}{}
	for _, c := range curApp.Credentials {
		keys[c.ApiKey] = struct{}{}
	}

	for _, c := range updateApp.Credentials {
		if _, ok := keys[c.ApiKey]; !ok {
			cred = c
			break
		}
	}

	// get the cred expiry time if it is set
	credBuilder := prov.NewCredentialBuilder()
	if p.credExpDays > 0 {
		credBuilder = credBuilder.SetExpirationTime(time.UnixMilli(int64(cred.ExpiresAt)))
	}

	//var cr prov.Credential
	cr := credBuilder.SetAPIKey(cred.ApiKey)

	logger.Info("created credential")

	hash, _ := util.ComputeHash(cred.ApiKey)
	return ps.AddProperty(credRefKey, fmt.Sprintf("%v", hash)).AddProperty(appRefName, appName).Success(), cr
}

// CredentialUpdate -
func (p provisioner) CredentialUpdate(req prov.CredentialRequest) (prov.RequestStatus, prov.Credential) {
	logger := p.logger.WithField("handler", "CredentialDeprovision").WithField("application", req.GetApplicationName())

	logger.Info("updating credential")
	ps := prov.NewRequestStatusBuilder()

	appName := req.GetCredentialDetailsValue(appRefName)
	if appName == "" {
		return failed(logger, ps, fmt.Errorf("application name not found")), nil
	}

	app, err := p.client.GetApp(appName)
	if err != nil {
		return failed(logger, ps, fmt.Errorf("error retrieving app: %s", err)), nil
	}

	credKey := ""
	curHash := req.GetCredentialDetailsValue(credRefKey)
	if curHash == "" {
		return failed(logger, ps, fmt.Errorf("credential reference not found")), nil
	}

	for _, cred := range app.Credentials {
		thisHash, _ := util.ComputeHash(cred.ApiKey)
		if curHash == fmt.Sprintf("%v", thisHash) {
			credKey = cred.ApiKey
			break
		}
	}

	if credKey == "" {
		return failed(logger, ps, fmt.Errorf("error retrieving the requested credential")), nil
	}

	for _, cred := range app.Credentials {
		for _, sub := range cred.Subscriptions {
			if req.GetCredentialAction() == prov.Suspend {
				_, err = p.client.UpdateCredential(sub.Id, app.Id)
			} else if req.GetCredentialAction() == prov.Enable {
				_, err = p.client.UpdateCredential(sub.Id, app.Id)
			} else {
				return failed(logger, ps, fmt.Errorf("could not perform the requested action: %s", req.GetCredentialAction())), nil
			}
		}
	}

	if err != nil {
		return failed(logger, ps, fmt.Errorf("error updating the app credential: %s", err)), nil
	}

	logger.Info("updated credential")
	return ps.Success(), nil
}

func failed(logger log.FieldLogger, ps prov.RequestStatusBuilder, err error) prov.RequestStatus {
	ps.SetMessage(err.Error())
	logger.WithError(err).Error("provisioning event failed", err)
	return ps.Failed()
}
