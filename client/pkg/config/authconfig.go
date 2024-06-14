package config

// AuthConfig - represents the config for gateway
type AuthConfig struct {
	Token string `config:"token"`
}

func (a *AuthConfig) GetToken() string {
	return a.Token
}
