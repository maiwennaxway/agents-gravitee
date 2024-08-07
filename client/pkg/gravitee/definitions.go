package gravitee

import "github.com/maiwennaxway/agents-gravitee/client/pkg/gravitee/models"

const (
	ClonedProdAttribute = "ClonedProduct"
)

type AllApis struct {
	Apis       []models.Api      `json:"data"`
	Pagination models.Pagination `json:"pagination"`
}

type AllSpecs struct {
	Specs      []models.Spec     `json:"pages"`
	Pagination models.Pagination `json:"pagination"`
}

type AllSubs struct {
	Subs []models.Subscriptions `json:"data"`
}

type AllPlans struct {
	Plans []models.Plan `json:"data"`
}

// PortalResponse
type PortalResponse struct {
	Status    string     `json:"status"`
	Message   string     `json:"message"`
	Code      string     `json:"code"`
	ErrorCode string     `json:"error_code"`
	RequestID string     `json:"request_id"`
	Data      PortalData `json:"data"`
}

// PortalsResponse
type PortalsResponse struct {
	Status    string       `json:"status"`
	Message   string       `json:"message"`
	Code      string       `json:"code"`
	ErrorCode string       `json:"error_code"`
	RequestID string       `json:"request_id"`
	Data      []PortalData `json:"data"`
}

type PortalData struct {
	ID                   string `json:"id"`
	Name                 string `json:"name"`
	Description          string `json:"description"`
	CustomDomain         string `json:"customDomain"`
	OrgName              string `json:"orgName"`
	Status               string `json:"status"`
	VisibleToCustomers   bool   `json:"visibleToCustomers"`
	HTTPS                bool   `json:"https"`
	DefaultDomain        string `json:"defaultDomain"`
	CustomeDomainEnabled bool   `json:"customDomainEnabled"`
	DefaultURL           string `json:"defaultURL"`
	CurrentURL           string `json:"currentURL"`
	CurrentDomain        string `json:"currentDomain"`
}

// VirtualHosts
type VirtualHosts []string

type PolicyDetail struct {
	PolicyType string `json:"policyType"`
}
