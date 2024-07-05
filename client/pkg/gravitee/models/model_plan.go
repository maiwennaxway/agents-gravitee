package models

type Plan struct {
	Flows             []Flows    `json:"flows"`
	Id                string     `json:"id"`
	Name              string     `json:"name"`
	Description       string     `json:"description"`
	ApiId             string     `json:"apiId"`
	Security          Security   `json:"security"`
	Characteristics   []string   `json:"characteristics"`
	ClosedAt          string     `json:"closedAt"`
	CommentMessage    string     `json:"commentMessage"`
	CommentRequired   bool       `json:"commentRequired"`
	CreatedAt         string     `json:"createdAt"`
	CrossId           string     `json:"crossId"`
	DefinitionVersion string     `json:"definitionVersion"`
	ExcludedGroups    []string   `json:"excludedGroups"`
	GeneralConditions string     `json:"generalConditions"`
	Order             int        `json:"order"`
	PublishedAt       string     `json:"publishedAt"`
	SelectionRule     string     `json:"selectionRule"`
	Status            string     `json:"status"`
	Tags              []string   `json:"tags"`
	Type              string     `json:"type"`
	UpdatedAt         string     `json:"updatedAt"`
	Validation        string     `json:"validation"`
	Pagination        Pagination `json:"pagination,omitempty"`
}
