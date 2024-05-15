package gravitee

import "github.com/maiwennaxway/agents-gravitee/client/pkg/gravitee/models"

const (
	ClonedProdAttribute = "ClonedProduct"
)

// Apis
type Apis []models.Api

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

type SpecDetails struct {
	ID          string        `json:"id"`
	Kind        string        `json:"kind"`
	Name        string        `json:"name"`
	Created     string        `json:"created"`
	Creator     string        `json:"creator"`
	Modified    string        `json:"modified"`
	IsTrashed   bool          `json:"isTrashed"`
	Permissions *string       `json:"permissions"`
	SelfLink    string        `json:"self"`
	ContentLink string        `json:"content"`
	Contents    []SpecDetails `json:"contents"`
	FolderLink  string        `json:"folder"`
	FolderID    string        `json:"folderId"`
	Body        *string       `json:"body"`
}

// VirtualHosts
type VirtualHosts []string

type PolicyDetail struct {
	PolicyType string `json:"policyType"`
}
