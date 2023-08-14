package config

type JiraConfig struct {
	BaseURL  string `mapstructure:"base_url"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}
