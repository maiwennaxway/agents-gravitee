package models

type Phase struct {
	Logging_phaserequest  bool `json:"request"`
	Logging_phaseresponse bool `json:"response"`
}
