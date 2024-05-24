package models

type Spec struct {
	Id                     string       `json:"id"`
	CrossId                string       `json:"crossId"`
	Name                   string       `json:"name"`
	Type                   string       `json:"type"`
	Content                string       `json:"content"`
	Order                  int          `json:"order"`
	LastContributor        string       `json:"lastContributor"`
	Published              bool         `json:"published"`
	Visibility             string       `json:"visibility"`
	UpdatedAt              string       `json:"updatedAt"`
	ContentType            string       `json:"contentType"`
	Homepage               bool         `json:"homepage,omitempty"`
	ParentId               string       `json:"parentId,omitempty"`
	ParentPath             string       `json:"parentPath,omitempty"`
	ExcludedGroups         []string     `json:"excludedGroups"`
	ExcludedAccessControls bool         `json:"excludedAccessControls"`
	Hidden                 bool         `json:"hidden,omitempty"`
	GeneralConditions      bool         `json:"generalConditions"`
	Breadcrumb             []Breadcrumb `json:"breadcrumb,omitempty"`
}
