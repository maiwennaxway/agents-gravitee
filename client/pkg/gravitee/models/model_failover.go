package models

type Failover struct {
	MaxAttempts  int      `json:"maxAttempts,omitempty"`
	RetryTimeout int      `json:"retryTimeout,omitempty"`
	Cases        []string `json:"cases,omitempty"`
}
