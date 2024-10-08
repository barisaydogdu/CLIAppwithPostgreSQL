package pkg

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectToDB() (*sql.DB, error) {
	conn, err := sql.Open("postgres", NewEnvDbConfig().ConnString())
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}
	return conn, nil
}

func (c *EnvDbConfig) ConnString() string {
	conStr := fmt.Sprintf("host=%s port=%s user=%s password=%s database=%s",
		c.Host(), c.Port(), c.User(), c.Password(), c.DatabaseName())

	return conStr
}
