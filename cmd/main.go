package main

import (
	"context"
	"fmt"
	"log"

	"github.com/barisaydogdu/PostgreSQLwithGo/handlers"
	pstgr "github.com/barisaydogdu/PostgreSQLwithGo/infrastructure/postgre"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conn, err := pstgr.ConnectToDB()
	if err != nil {
		log.Fatal(err)
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
