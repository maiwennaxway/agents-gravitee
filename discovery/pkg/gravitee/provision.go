package gravitee

/*import (
	"fmt"
	"time"

	v1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
	defs "github.com/Axway/agent-sdk/pkg/apic/definitions"
	prov "github.com/Axway/agent-sdk/pkg/apic/provisioning"
	"github.com/Axway/agent-sdk/pkg/util"
	"github.com/Axway/agent-sdk/pkg/util/log"
	"github.com/maiwennaxway/agents-gravitee/client/pkg/gravitee"
	"github.com/maiwennaxway/agents-gravitee/client/pkg/gravitee/models"
)


type provisioner struct {
	client                client
}

type cacheManager interface {
	GetAccessRequestsByApp(managedAppName string) []*v1.ResourceInstance
	GetAPIServiceInstanceByName(apiName string) (*v1.ResourceInstance, error)
}

type client interface {
	GetEnvId() string
}

// NewProvisioner creates a type to implement the SDK Provisioning methods for handling subscriptions
func NewProvisioner(client client, credExpDays int, cacheMan cacheManager, isProductMode bool) prov.Provisioning {
	return &provisioner{
		client:        client,
		credExpDays:   credExpDays,
		cacheManager:  cacheMan,
		isProductMode: isProductMode,
		logger:        log.NewFieldLogger().WithComponent("provision").WithPackage("gravitee"),
	}
}

// AccessRequestProvision - adds an api to an application
func (p provisioner) AccessRequestProvision(req prov.AccessRequest) (prov.RequestStatus, prov.AccessData) {
	instDetails := req.GetInstanceDetails()
	apiID := util.ToString(instDetails[defs.AttrExternalAPIID])
	stage := util.ToString(instDetails[defs.AttrExternalAPIStage])
	logger := p.logger.WithField("handler", "AccessRequestProvision").WithField("apiID", apiID).WithField("application", req.GetApplicationName())
	if stage != "" {
		logger = logger.WithField("stage", stage)
	}

	logger.Info("processing access request")
	ps := prov.NewRequestStatusBuilder()
	devID := p.client.GetEnvId()

	if apiID == "" {
		return failed(logger, ps, fmt.Errorf("%s name not found", defs.AttrExternalAPIID)), nil
	}

	// stage is required for proxy mode
	if stage == "" && !p.isProductMode {
		return failed(logger, ps, fmt.Errorf("%s name not found", defs.AttrExternalAPIStage)), nil
	}

	appName := req.GetApplicationName()
	if appName == "" {
		return failed(logger, ps, fmt.Errorf("application name not found")), nil
	}

	// get plan name from access request
	// get api product, or create new one
	apiProductName := getAPIProductName(apiID, req.GetQuota())
	quota := ""
	quotaInterval := "1"
	quotaTimeUnit := ""

	if q := req.GetQuota(); q != nil {
		quota = fmt.Sprintf("%d", q.GetLimit())

		switch q.GetInterval() {
		case prov.Daily:
			quotaTimeUnit = "day"
		case prov.Weekly:
			quotaTimeUnit = "day"
			quotaInterval = "7"
		case prov.Monthly:
			quotaTimeUnit = "month"
		case prov.Annually:
			quotaTimeUnit = "month"
			quotaInterval = "12"
		default:
			return failed(logger, ps, fmt.Errorf("invalid quota time unit: received %s", q.GetIntervalString())), nil
		}
	}

	var product *models.ApiProduct
	var err error
	if p.isProductMode {
		logger.Debug("handling for product mode")
		product, err = p.productModeCreateProduct(logger, apiProductName, apiID, quota, quotaInterval, quotaTimeUnit)
	} else {
		logger.Debug("handling for proxy mode")
		product, err = p.proxyModeCreateProduct(logger, apiProductName, apiID, stage, quota, quotaInterval, quotaTimeUnit)
	}
	if err != nil {
		return failed(logger, ps, fmt.Errorf("failed to create api product: %s", err)), nil
	}

	app, err := p.client.GetDeveloperApp(appName)
	if err != nil {
		return failed(logger, ps, fmt.Errorf("failed to retrieve app %s: %s", appName, err)), nil
	}

	if len(app.Credentials) == 0 {
		// no credentials to add access too
		return ps.AddProperty(prodNameRef, product.Name).Success(), nil
	}

	// add api to credentials that are not associated with it
	for _, cred := range app.Credentials {
		addProd := true
		enableProd := false
		for _, p := range cred.ApiProducts {
			if p.Apiproduct == apiProductName {
				addProd = false
				// already has the product, make sure its enabled
				if p.Status == "revoked" {
					enableProd = true
				}
				break
			}
		}

		// add the product to this credential
		if addProd {
			cpr := gravitee.CredentialProvisionRequest{
				ApiProducts: []string{apiProductName},
			}

			_, err = p.client.AddCredentialProduct(appName, devID, cred.ConsumerKey, cpr)
			if err != nil {
				return failed(logger, ps, fmt.Errorf("failed to add api product %s to credential: %s", apiProductName, err)), nil
			}
		}

		// enable the product for this credential
		if enableProd {
			err = p.client.UpdateCredentialProduct(appName, devID, cred.ConsumerKey, apiProductName, true)
			if err != nil {
				return failed(logger, ps, fmt.Errorf("failed to add enable api product %s on credential: %s", apiProductName, err)), nil
			}
		}
	}

	logger.Info("granted access")

	return ps.AddProperty(prodNameRef, product.Name).Success(), nil
}

func (p provisioner) productModeCreateProduct(logger log.FieldLogger, targetProductName, currentProductName, quota, quotaInterval, quotaTimeUnit string) (*models.ApiProduct, error) {
	// get the base product
	curProduct, err := p.client.GetProduct(currentProductName)
	if err != nil || targetProductName == currentProductName {
		// no new product required use the base product
		return curProduct, err
	}

	// check if the product/quota map already exists as a product
	product, err := p.client.GetProduct(targetProductName)

	// only create a product if one is not found
	if err != nil {
		attributes := []models.Attribute{}
		if p.shouldCloneAttributes {
			attributes = curProduct.Attributes
		}
		attributes = append(attributes, []models.Attribute{
			{
				Name:  agentProductTagName,
				Value: agentProductTagValue,
			},
			{
				Name:  gravitee.ClonedProdAttribute,
				Value: curProduct.Name,
			},
		}...)

		product = &models.ApiProduct{
			ApiResources: curProduct.ApiResources,
			ApprovalType: curProduct.ApprovalType,
			Attributes:   attributes,
			Description:  curProduct.Description,
			DisplayName:  targetProductName,
			Environments: curProduct.Environments,
			Name:         targetProductName,
			Proxies:      curProduct.Proxies,
			Scopes:       curProduct.Scopes,
		}
		if quota != "" {
			product.Quota = quota
			product.QuotaInterval = quotaInterval
			product.QuotaTimeUnit = quotaTimeUnit
		}
		logger.Infof("creating api product")
		product, err = p.client.CreateAPIProduct(product)
		if err != nil {
			return nil, err
		}
	}
	return product, err
}

func (p provisioner) proxyModeCreateProduct(logger log.FieldLogger, apiProductName, proxy, stage, quota, quotaInterval, quotaTimeUnit string) (*models.ApiProduct, error) {
	product, err := p.client.GetProduct(apiProductName)

	// only create a product if one is not found
	if err != nil {
		product = &models.ApiProduct{
			ApiResources: []string{},
			ApprovalType: "auto",
			DisplayName:  apiProductName,
			Environments: []string{stage},
			Name:         apiProductName,
			Proxies:      []string{proxy},
		}
		if quota != "" {
			product.Quota = quota
			product.QuotaInterval = quotaInterval
			product.QuotaTimeUnit = quotaTimeUnit
		}
		logger.Infof("creating api product")
		product, err = p.client.CreateAPIProduct(product)
		if err != nil {
			return nil, err
		}
	}
	return product, err
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

	err := p.client.RemoveDeveloperApp(appName, p.client.GetEnvId())
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
	app := models.DeveloperApp{
		Attributes: []models.Attribute{
			gravitee.GraviteeAgentAttribute,
		},
		EnvId: p.client.GetEnvId(),
		Name:  req.GetManagedApplicationName(),
	}

	newApp, err := p.client.CreateDeveloperApp(app)
	if err != nil {
		return failed(logger, ps, fmt.Errorf("failed to create app: %s", err))
	}

	// remove the credential created by default for the application, the credential request will create a new one
	p.client.RemoveAppCredential(app.Name, p.client.GetEnvId(), newApp.Credentials[0].ConsumerKey)

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

	app, err := p.client.GetDeveloperApp(appName)
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
		thisHash, _ := util.ComputeHash(cred.ConsumerKey)
		if curHash == fmt.Sprintf("%v", thisHash) {
			credKey = cred.ConsumerKey
			break
		}
	}

	if credKey == "" {
		return ps.Success()
	}

	// remove the credential created by default for the application, the credential request will create a new one
	err = p.client.RemoveAppCredential(app.Name, p.client.GetEnvId(), credKey)
	if err != nil {
		return failed(logger, ps, fmt.Errorf("unexpected error removing the credential"))
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

	curApp, err := p.client.GetDeveloperApp(appName)
	if err != nil {
		return failed(logger, ps, fmt.Errorf("error retrieving app: %s", err)), nil
	}

	// associate all products
	accReqs := p.cacheManager.GetAccessRequestsByApp(appName)
	products := []string{}
	for _, arInst := range accReqs {
		productName, err := util.GetAgentDetailsValue(arInst, prodNameRef)
		if err == nil && productName != "" {
			products = append(products, productName)
		}
	}
	if len(products) == 0 {
		return failed(logger, ps, fmt.Errorf("at least one product access is required for a credential")), nil
	}

	updateApp, err := p.client.CreateAppCredential(curApp.Name, p.client.GetEnvId(), products, p.credExpDays)
	if err != nil {
		return failed(logger, ps, fmt.Errorf("error creating app credential: %s", err)), nil
	}

	// find the new cred
	cred := models.DeveloperAppCredentials{}
	keys := map[string]struct{}{}
	for _, c := range curApp.Credentials {
		keys[c.ConsumerKey] = struct{}{}
	}

	for _, c := range updateApp.Credentials {
		if _, ok := keys[c.ConsumerKey]; !ok {
			cred = c
			break
		}
	}

	// get the cred expiry time if it is set
	credBuilder := prov.NewCredentialBuilder()
	if p.credExpDays > 0 {
		credBuilder = credBuilder.SetExpirationTime(time.UnixMilli(int64(cred.ExpiresAt)))
	}

	var cr prov.Credential
	t := req.GetCredentialType()
	if t == prov.APIKeyCRD {
		cr = credBuilder.SetAPIKey(cred.ConsumerKey)
	} else {
		cr = credBuilder.SetOAuthIDAndSecret(cred.ConsumerKey, cred.ConsumerSecret)
	}

	logger.Info("created credential")

	hash, _ := util.ComputeHash(cred.ConsumerKey)
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

	app, err := p.client.GetDeveloperApp(appName)
	if err != nil {
		return failed(logger, ps, fmt.Errorf("error retrieving app: %s", err)), nil
	}

	credKey := ""
	curHash := req.GetCredentialDetailsValue(credRefKey)
	if curHash == "" {
		return failed(logger, ps, fmt.Errorf("credential reference not found")), nil
	}

	for _, cred := range app.Credentials {
		thisHash, _ := util.ComputeHash(cred.ConsumerKey)
		if curHash == fmt.Sprintf("%v", thisHash) {
			credKey = cred.ConsumerKey
			break
		}
	}

	if credKey == "" {
		return failed(logger, ps, fmt.Errorf("error retrieving the requested credential")), nil
	}

	if req.GetCredentialAction() == prov.Suspend {
		err = p.client.UpdateAppCredential(app.Name, p.client.GetEnvId(), credKey, false)
	} else if req.GetCredentialAction() == prov.Enable {
		err = p.client.UpdateAppCredential(app.Name, p.client.GetEnvId(), credKey, true)
	} else {
		return failed(logger, ps, fmt.Errorf("could not perform the requested action: %s", req.GetCredentialAction())), nil
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
*/
