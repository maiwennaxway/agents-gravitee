package config

// AuthConfig - represents the config for gateway
type AuthConfig struct {
	URL      string `config:"url"`
	Username string `config:"username"`
	Password string `config:"password"`
	Token    string `config:"token"`
}

// GetURL - Returns the gravitee username
func (a *AuthConfig) GetURL() string {
	return a.URL
}

// GetUsername - Returns the gravitee username
func (a *AuthConfig) GetUsername() string {
	return a.Username
}

// GetPassword - Returns the gravitee password
func (a *AuthConfig) GetPassword() string {
	return a.Password
}

func (a *AuthConfig) GetToken() string {
	return a.Token
}
