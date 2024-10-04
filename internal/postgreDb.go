package pkg

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectToDB(conStr string) (*sql.DB, error) {
	conn, err := sql.Open("postgres", conStr)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}

func (c EnvDbConfig) GoDotEnvVariable(connString string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(connString)
}

func (c EnvDbConfig) ConnString() string {
	conStr := fmt.Sprintf("host=%s port=%s user=%s password=%s database=%s",
		c.Host(), c.Port(), c.User(), c.Password(), c.DatabaseName())

	return conStr
}
