package models

type Post struct {
	Name          string        `json:"name"`
	Desc          string        `json:"description"`
	Enabled       bool          `json:"enabled"`
	Policy        string        `json:"policy"`
	Configuration Configuration `json:"configuration"`
	Condition     string        `json:"condition"`
}
