package models

type App struct {
	Id          string           `json:"id"`
	Name        string           `json:"name"`
	Descripion  string           `json:"description"`
	Status      string           `json:"status"`
	Domain      string           `json:"domain"`
	Type        string           `json:"type"`
	CreatedAt   string           `json:"created_at"`
	UpdatedAt   string           `json:"updated_at"`
	Owner       PrimaryOwner     `json:"primaryOwner"`
	ApiKeyMode  string           `json:"apiKeyMode"`
	Credentials []AppCredentials `json:"credentials"`
}
