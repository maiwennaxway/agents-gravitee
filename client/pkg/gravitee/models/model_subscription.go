package models

type Subscriptions struct {
	Id                    string                `json:"id"`
	Api                   Api                   `json:"api"`
	Plan                  Plan                  `json:"plan"`
	App                   App                   `json:"application"`
	ConsumerConfiguration ConsumerConfiguration `json:"consumerConfiguration,omitempty"`
	FailureCause          string                `json:"failureCause,omitempty"`
	Status                string                `json:"status,omitempty"`
	ConsumerStatus        string                `json:"consumerStatus,omitempty"`
	ProcessedBy           ProcessedBy           `json:"processedBy,omitempty"`
	SubscribedBy          SubscribedBy          `json:"subscribedBy"`
	ProcessedAt           string                `json:"processedAt"`
	StartingAt            string                `json:"startingAt"`
	EndingAt              string                `json:"endingAt,omitempty"`
	CreatedAt             string                `json:"createdAt"`
	UpdatedAt             string                `json:"updatedAt,omitempty"`
	ClosedAt              string                `json:"closedAt,omitempty"`
	PausedAt              string                `json:"pausedAt,omitempty"`
	ConsumerPausedAt      string                `json:"consumerPausedAt,omitempty"`
}
