package models

type Subscriptions struct {
	Id                    string                `json:"id"`
	Api                   Api                   `json:"api"`
	Plan                  Plan                  `json:"plan"`
	App                   App                   `json:"application"`
	ConsumerConfiguration ConsumerConfiguration `json:"consumerConfiguration"`
	FailureCause          string                `json:"failureCause"`
	Status                string                `json:"status"`
	ConsumerStatus        string                `json:"consumerStatus"`
	ProcessedBy           ProcessedBy           `json:"processedBy"`
	SubscribedBy          SubscribedBy          `json:"subscribedBy"`
	ProcessedAt           string                `json:"processedAt"`
	StartingAt            string                `json:"startingAt"`
	EndingAt              string                `json:"endingAt"`
	CreatedAt             string                `json:"createdAt"`
	UpdatedAt             string                `json:"updatedAt"`
	ClosedAt              string                `json:"closedAt"`
	PausedAt              string                `json:"pausedAt"`
	ConsumerPausedAt      string                `json:"consumerPausedAt"`
}
