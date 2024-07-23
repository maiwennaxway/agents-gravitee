package models

type Plan struct {
	Flows             []Flows  `json:"flows,omitempty"`
	Id                string   `json:"id,omitempty"`
	Name              string   `json:"name,omitempty"`
	Description       string   `json:"description,omitempty"`
	ApiId             string   `json:"apiId,omitempty"`
	Security          Security `json:"security,omitempty"`
	Characteristics   []string `json:"characteristics,omitempty"`
	ClosedAt          string   `json:"closedAt,omitempty"`
	CommentMessage    string   `json:"commentMessage,omitempty"`
	CommentRequired   bool     `json:"commentRequired,omitempty"`
	CreatedAt         string   `json:"createdAt,omitempty"`
	CrossId           string   `json:"crossId,omitempty"`
	DefinitionVersion string   `json:"definitionVersion"`
	ExcludedGroups    []string `json:"excludedGroups,omitempty"`
	GeneralConditions string   `json:"generalConditions,omitempty"`
	Order             int      `json:"order,omitempty"`
	PublishedAt       string   `json:"publishedAt,omitempty"`
	SelectionRule     string   `json:"selectionRule,omitempty"`
	Status            string   `json:"status,omitempty"`
	Tags              []string `json:"tags,omitempty"`
	Type              string   `json:"type,omitempty"`
	UpdatedAt         string   `json:"updatedAt,omitempty"`
	Validation        string   `json:"validation,omitempty"`
}
