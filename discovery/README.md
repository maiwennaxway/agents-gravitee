# gravitee Discovery Agent

The Discovery agent finds deployed API Proxies in gravitee then sends them to API Central

## Build and run

The following make targets are available

| Target          | Description                                                    | Output(s)                     |
| --------------- | -------------------------------------------------------------- | ----------------------------- |
| dep             | downloads all dependencies needed to build the discovery agent | /vendor                       |
| test            | runs go test against all test files int he repo                | test results                  |
| update-sdk      | pulls the latest changes to main on the SDK repo               |                               |
| build           | builds the binary discovery agent                              | bin/gravitee_discovery_agent    |
| gravitee-generate | generates the models for the gravitee APIs                       | pkg/gravitee/models             |
| docker-build    | builds the discovery agent in a docker container               | gravitee-discovery-agent:latest |

### Build (Docker)

```shell
make docker-build
```

### Run (Docker)

```shell
docker run --env-file env_vars -v `pwd`/keys:/keys gravitee-discovery-agent:latest
```

### Build (Windows)

* Build the agent using the following command

```shell
go build -tags static_all \
    -ldflags="-X 'github.com/Axway/agent-sdk/pkg/cmd.BuildTime=$${time}' \
        -X 'github.com/Axway/agent-sdk/pkg/cmd.BuildVersion=$${version}' \
        -X 'github.com/Axway/agent-sdk/pkg/cmd.BuildCommitSha=$${commit_id}' \
        -X 'github.com/Axway/agent-sdk/pkg/cmd.BuildAgentName=graviteeDiscoveryAgent'" \
    -a -o ./bin/gravitee_discovery_agent.exe ./main.go
```

### Run (Windows)

* After a successful build, you should see the executable under the bin folder.   And you can execute it using the following command

```shell
./gravitee_discovery_agent.exe --envFile env_vars
```

## Discovery Mode - Proxy

This is the default operating mode that discoveries API Proxies and attempts to match them to Specs

### Proxy discovery

* Find all specs
  * Parse all specs to determine endpoints with in
  * Save info to cache
* Find all Deployed API Proxies
  * Find the Spec
    * Proxy Revision has spec set, use it
    * Proxy Revision has association.json resource file, get path
      * Using path check to see if it is in the specs that were found by agent, use it
    * Using deployed URL path check for specs for match, use it
  * Check proxy for Key or Oauth policy for authentication
  * Create API Service
    * If spec was found use it in revision
    * If spec was not found create as unstructured
    * Attach appropriate Credential Request Definition based on policy in proxy

### Proxy provisioning

* Managed Application
  * Creates a new App on gravitee under the configured developer
* Access Request
  * Creates a new Product, or uses existing, based off the gravitee-Proxy and Central Plan combination
  * Associates the new Product to any existing Credentials on the Application
* Credential
  * Creates a new Credential on the App and associates all Access Requests products to it

## Discovery Mode - Product

This mode can be setting the `gravitee_DISCOVERYMODE` environment variable to `product`

### Product discovery

* Find all specs
  * Parsing is not necessary in this mode
  * Save info to cache
* Find all Products defined
  * Using the product's name or display name, match it to a spec (case insensitive)
  * If a spec is found create an API Service
    * Use product definition, add attributes to Service
    * Donwload and attach spec file

### Product provisioning

* Managed Application
  * Creates a new App on gravitee under the configured developer
* Access Request
  * Creates a new Product, or uses existing, using the product associated with the API Service as a template
  * Associates the new Product to any existing Credentials on the Application
* Credential
  * Creates a new Credential on the App and associates all Access Requests products to it

## Quota enforcement

In both modes the provisioning process will set quota values on the created Product when handling Access Requests. In order for gravitee to enforce quota based on the values set in teh Product a Quota Enforcement Policy needs to be set on the deployed Proxy.

Here is a sample Quota policy that may be added to the desired Proxies.

```xml
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<Quota async="false" continueOnError="false" enabled="true" name="impose-quota">
    <DisplayName>Impose Quota</DisplayName>
    <Synchronous>true</Synchronous>
    <Distributed>true</Distributed>
    <Identifier ref="developer.app.name"/>
    <Allow countRef="verifyapikey.Verify-API-Key-1.apiproduct.developer.quota.limit"/>
    <Interval ref="verifyapikey.Verify-API-Key-1.apiproduct.developer.quota.interval "/>
    <TimeUnit ref="verifyapikey.Verify-API-Key-1.apiproduct.developer.quota.timeunit"/>
</Quota>
```

* Display Name - the name of the policy
* Distributed and Synchronous - set to true to ensure the counter is updated and enforced in a distributed and synchronous manner. Setting to false may allow extra calls.
* Identifier - how the proxy will count usage across multiple apps, so each get their own quota
* Allow - in this case using the API Key policy gets the quota limit from the product definition
* Interval - in this case using the API Key policy gets the quota interval from the product definition
* TimeUnit - in this case using the API Key policy gets the quota time unit from the product definition
Í

| Environment Variable       | Description                                                                          | Default (if applicable)           |
| -------------------------- | ------------------------------------------------------------------------------------ | --------------------------------- |
| gravitee_URL                 | The base gravitee URL for this agent to connect to                                     | https://api.enterprise.gravitee.com |
| gravitee_APIVERSION          | The version of the API for the agent to use                                          | v1                                |
| gravitee_DATAURL             | The base gravitee Data API URL for this agent to connect to                            | https://gravitee.com/dapi/api       |
| gravitee_ORGANIZATION        | The gravitee organization name                                                         |                                   |
| gravitee_EnvId         | The gravitee developer, email, that will own all apps                                  |                                   |
| gravitee_DISCOVERYMODE       | The mode in which the agent operates, discover proxies (proxy) or products (product) | proxy                             |
| gravitee_CLONEATTRIBUTES     | Set this to true if the tags on a product should also be cloned on provisioning      | false                             |
| gravitee_INTERVAL_PROXY      | The polling interval checking for API Proxy changes, only in proxy mode              | 30s (30 seconds), >=30s, <=5m     |
| gravitee_INTERVAL_PRODUCT    | The polling interval checking for Product changes, only in product mode              | 30s (30 seconds), >=30s, <=5m     |
| gravitee_INTERVAL_SPEC       | The polling interval for checking for new Specs                                      | 30m (30 minute), >=1m             |
| gravitee_WORKERS_PROXY       | The number of workers processing API Proxies, only in proxy mode                     | 10                                |
| gravitee_WORKERS_PRODUCT     | The number of workers processing Products, only in product mode                      | 10                                |
| gravitee_WORKERS_SPEC        | The number of workers processing API Specs                                           | 20                                |
| gravitee_AUTH_USERNAME       | The gravitee account username/email address                                            |                                   |
| gravitee_AUTH_PASSWORD       | The gravitee account password                                                          |                                   |
| gravitee_AUTH_URL            | The IDP URL                                                                          | https://login.gravitee.com          |
| gravitee_AUTH_SERVERUSERNAME | The IDP username for requesting tokens                                               | edgecli                           |
| gravitee_AUTH_SERVERPASSWORD | The IDP password for requesting tokens                                               | edgeclisecret                     |
