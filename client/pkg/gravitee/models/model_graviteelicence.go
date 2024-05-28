package models

type GraviteeLicence struct {
	Tier     string   `json:"tier"`
	Packs    []string `json:"packs"`
	Features []string `json:"features"`
}
