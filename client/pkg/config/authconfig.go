package config

// AuthConfig - represents the config for gateway
type AuthConfig struct {
	URL            string `config:"url"`
	ServerUsername string `config:"serverUsername"`
	ServerPassword string `config:"serverPassword"`
	Username       string `config:"username"`
	Password       string `config:"password"`
}

// GetServerUsername - Returns the gravitee auth server username
func (a *AuthConfig) GetServerUsername() string {
	return a.ServerUsername
}

// GetServerPassword - Returns the gravitee auth server password
func (a *AuthConfig) GetServerPassword() string {
	return a.ServerPassword
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
