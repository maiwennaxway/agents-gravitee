package models

type Flows struct {
	Id           string       `json:"id"`
	Name         string       `json:"name"`
	PathOperator PathOperator `json:"pathOperator"`
	Pre          []Pre        `json:"pre"`
	Post         []Post       `json:"post"`
	Enabled      bool         `json:"enabled"`
	Methods      []string     `json:"methods"`
	Condition    string       `json:"condition"`
	Consumers    []Consumer   `json:"consumers"`
	Stage        string       `json:"stage"`
}
