package pkg

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectToDB(connStr string) (*sql.DB, error) {
	conn, err := sql.Open("postgres", connStr)
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

func (c *EnvDbConfig) TestConnString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?",
		c.User(), c.Password(), c.Host(), c.Port(), c.TestDatabaseName())

}
