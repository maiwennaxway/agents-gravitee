# gravitee Traceability Agent

The Traceability agent finds logs from consumed gravitee proxies and sends the traffic data to Amplify Central

## Build and run

The following make targets are available

| Target          | Description                                                    | Output(s)                        |
| --------------- | -------------------------------------------------------------- | -------------------------------- |
| dep             | downloads all dependencies needed to build the discovery agent | /vendor                          |
| test            | runs go test against all test files int he repo                | test results                     |
| update-sdk      | pulls the latest changes to main on the SDK repo               |                                  |
| build           | builds the binary traceability agent                           | bin/gravitee_traceability_agent    |
| gravitee-generate | generates the models for the gravitee APIs                       | pkg/gravitee/models                |
| docker-build    | builds the traceability agent in a docker container            | gravitee-traceability-agent:latest |

### Build (Docker)

```shell
make docker-build
```

### Run (Docker)

```shell
docker run --env-file env_vars  -v `pwd`/data:/data -v `pwd`/keys:/keys gravitee-traceability-agent:latest
```

### Build (Windows)

* Build the agent using the following command

```shell
go build -tags static_all \
    -ldflags="-X 'github.com/Axway/agent-sdk/pkg/cmd.BuildTime=$${time}' \
        -X 'github.com/Axway/agent-sdk/pkg/cmd.BuildVersion=$${version}' \
        -X 'github.com/Axway/agent-sdk/pkg/cmd.BuildCommitSha=$${commit_id}' \
        -X 'github.com/Axway/agent-sdk/pkg/cmd.BuildAgentName=graviteeTraceabilityAgent'" \
    -a -o ./bin/gravitee_traceability_agent.exe ./main.go
```

### Run (Windows)

* After a successful build, you should see the executable under the bin folder.   And you can execute it using the following command

```shell
./gravitee_traceability_agent.exe --envFile env_vars
```

## Traceability agent variables

| Environment Variable       | Description                                                                                                              | Default (if applicable)           |
| -------------------------- | ------------------------------------------------------------------------------------------------------------------------ | --------------------------------- |
| gravitee_URL                 | The base gravitee URL for this agent to connect to                                                                         | https://api.enterprise.gravitee.com |
| gravitee_APIVERSION          | The version of the API for the agent to use                                                                              | v1                                |
| gravitee_DATAURL             | The base gravitee Data API URL for this agent to connect to                                                                | https://gravitee.com/dapi/api       |
| gravitee_ORGANIZATION        | The gravitee organization name                                                                                             |                                   |
| gravitee_DEVELOPERID         | The gravitee developer, email, that will own all apps                                                                      |                                   |
| gravitee_DISCOVERYMODE       | The mode in which the discovery agent operates, determines how stats are gathered, proxies (proxy) or products (product) | proxy                             |
| gravitee_INTERVAL_STATS      | The polling interval checking for API Proxy changes, only in proxy mode                                                  | 5m (5 minutes), >=1m, <=15m       |
| gravitee_AUTH_USERNAME       | The gravitee account username/email address                                                                                |                                   |
| gravitee_AUTH_PASSWORD       | The gravitee account password                                                                                              |                                   |
| gravitee_AUTH_URL            | The IDP URL                                                                                                              | https://login.gravitee.com          |
| gravitee_AUTH_SERVERUSERNAME | The IDP username for requesting tokens                                                                                   | edgecli                           |
| gravitee_AUTH_SERVERPASSWORD | The IDP password for requesting tokens                                                                                   | edgeclisecret                     |

