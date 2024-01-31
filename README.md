# Prerequisites

* You need an Axway Platform user account that is assigned the AMPLIFY Central admin role
* Your gravitee Gateway/Manager should be up and running and have APIs to be discovered and exposed in AMPLIFY Central

Let’s get started!

## Prepare AMPLIFY Central Environments

In this section we'll:

* [Create an environment in Central](#create-an-environment-in-central)
* [Create a service account](#create-a-service-account)

### Create an environment in Central

* Log into [Amplify Central](https://apicentral.axway.com)
* Navigate to "Topology" then "Environments"
* Click "+ Environment"
  * Select a name
  * Click "Save"
* To enable the viewing of the agent status in Amplify see [Visualize the agent status](https://docs.axway.com/bundle/amplify-central/page/docs/connect_manage_environ/environment_agent_resources/index.html#add-your-agent-resources-to-the-environment)

### Create a service account

* Create a public and private key pair locally using the openssl command

```sh
openssl genpkey -algorithm RSA -out private_key.pem -pkeyopt rsa_keygen_bits: 2048
openssl rsa -in private_key.pem -pubout -out public_key.pem
```

* Log into the [Amplify Platform](https://platform.axway.com)
* Navigate to "Organization" then "Service Accounts"
* Click "+ Service Account"
  * Select a name
  * Optionally add a description
  * Select "Client Certificate"
  * Select "Provide public key"
  * Select or paste the contents of the public_key.pem file
  * Select "Central admin"
  * Click "Save"
* Note the Client ID value, this and the key files will be needed for the agents

## Prepare gravitee

* Create an gravitee account
* Note the username and password used as the agents will need this to run
* Add a developer that will be the owner of all applications created by the agent

## Setup agent Environment Variables

The following environment variables file should be created for executing both of the agents

```ini
CENTRAL_ORGANIZATIONID=<Amplify Central Organization ID>
CENTRAL_TEAM=<Amplify Central Team Name>
CENTRAL_ENVIRONMENT=<Amplify Central Environment Name>   # created in Prepare AMPLIFY Central Environments step

CENTRAL_AUTH_CLIENTID=<Amplify Central Service Account>  # created in Prepare AMPLIFY Central Environments step
CENTRAL_AUTH_PRIVATEKEY=/keys/private_key.pem            # path to the key file created with openssl
CENTRAL_AUTH_PUBLICKEY=/keys/public_key.pem              # path to the key file created with openssl

gravitee_ORGANIZATION=<gravitee Organization>                # created in Prepare gravitee step
gravitee_EnvId=dev@email.address                     # created in Prepare gravitee step
gravitee_AUTH_USERNAME=<gravitee Username>                   # created in Prepare gravitee step
gravitee_AUTH_PASSWORD=<gravitee Password>                   # created in Prepare gravitee step
gravitee_AUTH_URL=<IDP URL>                                # The IDP the agent should request an auth token from for gravitee API Access (default: https://login.gravitee.com)
gravitee_AUTH_SERVERUSERNAME=<Auth Server Username>        # The username for requesting a token from the IDP server (default: edgecli)
gravitee_AUTH_SERVERPASSWORD=<Auth Server Password>        # The password for requesting a token from the IDP server (default: edgeclisecret)

LOG_LEVEL=info
LOG_OUTPUT=stdout
```

## Discovery Agent

Reference: [Discovery Agent](/discovery/README.md)

## Traceability Agent

Reference: [Traceability Agent](/traceability/README.md)
