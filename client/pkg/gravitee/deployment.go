package gravitee

/*"encoding/json"
"fmt"
"net/http"

"github.com/maiwennaxway/agents-gravitee/client/pkg/gravitee/models"*/

// à mettre de coté

func (a *GraviteeClient) GetDeployments(instance string) {

}

/*  /environments/{envId}/apis/{apiId}/deployments:
parameters:
- $ref: "#/components/parameters/envIdParam"
- $ref: "#/components/parameters/apiIdParam"
post:
tags:
	- APIs
summary: Request a deployment to gateway instances
description: |-
	Request a deployment for a given API. <br>
	An optional deployment label can be given to the requested deployment.

	User must have the API_DEFINITION[UPDATE] permission.
operationId: createApiDeployment
requestBody:
	content:
		application/json:
			schema:
				$ref: "#/components/schemas/ApiDeployment"
responses:
	"202":
		description: API deployment request received
	default:
		$ref: "#/components/responses/Error"*/
