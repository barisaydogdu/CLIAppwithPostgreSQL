package pkg

import (
	"fmt"
	"testing"

	"github.com/joho/godotenv"
)

func TestConnectToDB(t *testing.T) {
	connString := NewEnvDbConfig().TestConnString()

	conn, err := ConnectToDB(connString)

	if err != nil {
		t.Errorf("Error with connecting database %v", err)
	}

	if conn == nil {
		t.Errorf("Error there is not a valid connection. Connection is nil :%v", err)
	}
	conn.Close()
}

func TestLoadFile(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Errorf("Error loading .env file")
	}

}
