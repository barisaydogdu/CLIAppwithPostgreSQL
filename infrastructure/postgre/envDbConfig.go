package pkg

import "os"

type EnvDbConfig struct {
	host         string
	port         string
	user         string
	password     string
	database     string
	testdatabase string
	dbtesthost   string
	testdbuser   string
}

func NewEnvDbConfig() *EnvDbConfig {
	return &EnvDbConfig{
		host:         os.Getenv("DB_HOST"),
		port:         os.Getenv("DB_PORT"),
		user:         os.Getenv("DB_USER"),
		password:     os.Getenv("DB_PASSWORD"),
		database:     os.Getenv("DB_NAME"),
		testdatabase: os.Getenv("TEST_DB_NAME"),
		dbtesthost:   os.Getenv("DB_TEST_HOST"),
		testdbuser:   os.Getenv("TEST_DB_USER"),
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

func (c *EnvDbConfig) TestDatabaseName() string {
	return c.testdatabase
}

func (c *EnvDbConfig) TestDBHost() string {
	return c.dbtesthost
}

func (c *EnvDbConfig) TestDBUser() string {
	return c.testdbuser
}
