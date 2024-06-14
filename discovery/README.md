# gravitee Discovery Agent

The Discovery agent finds deployed API in Gravitee then sends them to API Central

## Build and run

The following make targets are available

| Target            | Description                                                    | Output(s)                       |
| ---------------   | -------------------------------------------------------------- | ------------------------------- |
| test              | runs go test against all test files int he repo                | test results                    |
| update-sdk        | pulls the latest changes to main on the SDK repo               |                                 |
| build             | builds the binary discovery agent                              | bin/gravitee_discovery_agent    |
| docker-build      | builds the discovery agent in a docker container               | gravitee-discovery-agent:latest |

### Build (Docker)

```shell
docker build -t agent:version .
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
    -a -o ./bin/GRAVITEE_discovery_agent.exe ./main.go
```

### Run (Windows)

* After a successful build, you should see the executable under the bin folder.   And you can execute it using the following command

```shell
./gravitee_discovery_agent.exe --envFile env_vars
```
## Discovery -

* Find all specs
  * Save info to cache
* Find all Apis defined
  * Using the api's id, match it to a spec
  * If a spec is found create an API Service
    * Use api definition, add attributes to Service
    * Download and attach spec file

| Environment Variable         | Description                                                                          | Default (if applicable)           |
| --------------------------   | ------------------------------------------------------------------------------------ | --------------------------------- |
| GRAVITEE_URL                 | The base gravitee URL for this agent to connect to                                   | https://api.company.com/          |
| GRAVITEE_APIVERSION          | The version of the API for the agent to use                                          | v1                                |
| GRAVITEE_ENVIRONNEMENT       | The gravitee environment on which API ill be foundable                               | DEFAULT                           |
| GRAVITEE_INTERVAL_API        | The polling interval checking for Api changes,                                       | 30s (30 seconds), >=30s, <=5m     |
| GRAVITEE_INTERVAL_SPEC       | The polling interval for checking for new Specs                                      | 30m (30 minute), >=1m             |
| GRAVITEE_WORKERS_API         | The number of workers processing Apis                                                | 10                                |
| GRAVITEE_WORKERS_SPEC        | The number of workers processing API Specs                                           | 20                                |
| GRAVITEE_AUTH_TOKEN          | The Gravitee account bearer token                                                    |                                   |
| GRAVITEE_API_URL             | The machine URL                                                                      |                                   |
