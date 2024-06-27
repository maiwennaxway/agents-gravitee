package models

type Pre struct {
	Name      string `json:"name"`
	Desc      string `json:"description"`
	Enabled   bool   `json:"enabled"`
	Policy    string `json:"policy"`
	Quota     Quota  `json:"configuration"`
	Condition string `json:"condition"`
}
