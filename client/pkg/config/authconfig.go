package config

// AuthConfig - represents the config for gateway
type AuthConfig struct {
	URL            string `config:"url"`
	ServerUsername string `config:"serverUsername"`
	ServerPassword string `config:"serverPassword"`
	Username       string `config:"username"`
	Password       string `config:"password"`
	tlsProtocols   string `config:"tlsProtocole"`
	tlsCiphers     string `config:"tlsCiphers"`
	keystore_type  string `config:"keystoreType"`
	keystore_path  string `config:"keystorePath"`
	keystore_passd string `config:"keystorePassword"`
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

func (a *AuthConfig) GettlsProtocols() string {
	return a.tlsProtocols
}

func (a *AuthConfig) GettlsCiphers() string {
	return a.tlsCiphers
}

func (a *AuthConfig) GetKeystoreType() string {
	return a.keystore_type
}

func (a *AuthConfig) GetKeystorePath() string {
	return a.keystore_path
}

func (a *AuthConfig) GetKeystorePassword() string {
	return a.keystore_passd
}
