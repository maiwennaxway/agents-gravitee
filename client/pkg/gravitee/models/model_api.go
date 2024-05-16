package models

type Api struct {
	EnvironmentId                  string            `json:"environmentId,omitempty"`
	ExecutionMode                  string            `json:"executionMode,omitempty"`
	ContextPath                    string            `json:"contextPath,omitempty"`
	Proxy                          Proxy             `json:"proxy,omitempty"`
	Services                       []Services        `json:"services"`
	PathMappings                   []string          `json:"pathMappings"`
	Entrypoints                    []Entrypoints     `json:"entrypoints,omitempty"`
	Id                             string            `json:"id"`
	Name                           string            `json:"name"`
	Description                    string            `json:"description"`
	CrossId                        string            `json:"crossId,omitempty"`
	ApiVersion                     string            `json:"apiVersion"`
	DefinitionVersion              string            `json:"definitionVersion"`
	DeployedAt                     string            `json:"deployedAt,omitempty"`
	CreatedAt                      string            `json:"createdAt"`
	UpdatedAt                      string            `json:"updatedAt"`
	LastModifiedAt                 int               `json:"lastModifiedAt,omitempty"`
	DisableMembershipNotifications bool              `json:"disableMembershipNotifications"`
	Groups                         []string          `json:"groups,omitempty"`
	State                          string            `json:"state,omitempty"`
	DeploymentState                string            `json:"deploymentState,omitempty"`
	Visibility                     string            `json:"visibility"`
	Labels                         []string          `json:"labels,omitempty"`
	LifecycleState                 string            `json:"lifecycleState,omitempty"`
	Tags                           []string          `json:"tags,omitempty"`
	PrimaryOwner                   PrimaryOwner      `json:"primaryOwner,omitempty"`
	Categories                     []string          `json:"categories,omitempty"`
	DefinitionContext              DefinitionContext `json:"definitionContext,omitempty"`
	WorkflowState                  string            `json:"workflowState,omitempty"`
	ResponseTemplates              interface{}       `json:"responseTemplates,omitempty"`
	Resources                      Resources         `json:"resources,omitempty"`
	Properties                     Properties        `json:"properties,omitempty"`
	Links                          Links             `json:"_links,omitempty"`
}
