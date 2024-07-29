package models

type AppCredentials struct {
	ApiKey                             string          `json:"key"`
	Id                                 string          `json:"id"`
	Application                        []App           `json:"application"`
	Subscriptions                      []Subscriptions `json:"subscriptions,omitempty"`
	Revoked                            bool            `json:"revoked,omitempty"`
	Paused                             bool            `json:"paused,omitempty"`
	Expired                            bool            `json:"expired,omitempty"`
	DaysToExpirationOnLastNotification int             `json:"daysToExpirationOnLastNotification,omitempty"`
	ExpiresAt                          int             `json:"expiresAt,omitempty"`
	CreatedAt                          string          `json:"createdAt,omitempty"`
	UpdatedAt                          string          `json:"updatedAt,omitempty"`
	RevokedAt                          string          `json:"revokedAt,omitempty"`
}
