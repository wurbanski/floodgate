package config

// DefaultLocation is where the config is usually at
// TODO(wurbanski): move it
var DefaultLocation = "config.yaml"

// Config is the default configuration for the app
type Config struct {
	Endpoint string `yaml:"endpoint"`
	Insecure bool   `yaml:"insecure"`
	// TODO(wurbanski): use other auths than basic
	Auth struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"auth"`
	Libraries []string `yaml:"libraries"`
	Resources []string `yaml:"resources"`
}
