package main

import (
	"context"
	"fmt"

	"github.com/barisaydogdu/PostgreSQLwithGo/handlers"
	pstgr "github.com/barisaydogdu/PostgreSQLwithGo/infrastructure/postgre"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Errorf("Error loading .env file %v", err)
	}

	conn, err := pstgr.ConnectToDB(pstgr.NewEnvDbConfig().ConnString())
	if err != nil {
		fmt.Errorf("There is something error with connecting database %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	hndlrs := handlers.NewHandlers(ctx, conn)

	hndlrs.Start()

	fmt.Println("Successfully connected to the database")

	hndlrs.Stop()
	conn.Close()
	cancel()
}
