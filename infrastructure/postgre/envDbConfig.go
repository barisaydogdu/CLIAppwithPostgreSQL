package pkg

import "os"

type EnvDbConfig struct {
	host     string
	port     string
	user     string
	password string
	database string
}

func NewEnvDbConfig() *EnvDbConfig {
	return &EnvDbConfig{
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
		database: os.Getenv("DB_NAME"),
	}
}

func (c *EnvDbConfig) Host() string {
	return c.host
}

func (c *EnvDbConfig) Port() string {
	return c.port
}

func (c *EnvDbConfig) User() string {
	return c.user
}

func (c *EnvDbConfig) Password() string {
	return c.password
}

func (c *EnvDbConfig) DatabaseName() string {
	return c.database
}
