package models

type Analytics struct {
	Enabled  bool     `json:"enabled"`
	Sampling Sampling `json:"sampling"`
	Logging  Logging  `json:"logging"`
}
