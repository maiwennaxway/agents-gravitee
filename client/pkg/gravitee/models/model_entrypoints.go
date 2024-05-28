package models

type Entrypoints struct {
	Entrypoints_target string   `json:"target"`
	Entrypoints_host   string   `json:"host"`
	Entrypoints_tags   []string `json:"tags"`
}
