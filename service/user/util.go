package user

import (
	"database/sql"
	"fmt"

	postgre "github.com/barisaydogdu/PostgreSQLwithGo/infrastructure/postgre"
	"github.com/joho/godotenv"
)

func ConnectToDB() (*sql.DB, error) {
	conn, err := postgre.ConnectToDB(postgre.NewEnvDbConfig().TestConnString())
	if err != nil {
		return nil, fmt.Errorf("Could not connect database %v", err)
	}
	return conn, nil
}

func LoadEnv() {
	if err := godotenv.Load("../../.env"); err != nil {
		fmt.Errorf("Error loading .env file %v", err)
	}
}
