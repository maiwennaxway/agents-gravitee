package models

type Flows struct {
	Id           string       `json:"id,omitempty"`
	Name         string       `json:"name,omitempty"`
	PathOperator PathOperator `json:"pathOperator,omitempty"`
	Pre          []Pre        `json:"pre,omitempty"`
	Post         []Post       `json:"post,omitempty"`
	Enabled      bool         `json:"enabled,omitempty"`
	Methods      []string     `json:"methods,omitempty"`
	Condition    string       `json:"condition,omitempty"`
	Consumers    []Consumer   `json:"consumers,omitempty"`
	Stage        string       `json:"stage,omitempty"`
}
