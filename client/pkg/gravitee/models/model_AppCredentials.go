package models

type AppCredentials struct {
	ApiKey                             string          `json:"key"`
	Id                                 string          `json:"id"`
	Application                        []App           `json:"application"`
	Subscriptions                      []Subscriptions `json:"subscriptions"`
	Revoked                            bool            `json:"revoked"`
	Paused                             bool            `json:"paused"`
	Expired                            bool            `json:"expired"`
	DaysToExpirationOnLastNotification int             `json:"daysToExpirationOnLastNotification"`
	ExpiresAt                          int             `json:"expiresAt"`
	CreatedAt                          string          `json:"createdAt"`
	UpdatedAt                          string          `json:"updatedAt"`
	RevokedAt                          string          `json:"revokedAt"`
}
