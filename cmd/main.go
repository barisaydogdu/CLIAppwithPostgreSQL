package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/barisaydogdu/PostgreSQLwithGo/handlers"
	pstgr "github.com/barisaydogdu/PostgreSQLwithGo/infrastructure/postgre"
	repository "github.com/barisaydogdu/PostgreSQLwithGo/repository/user"
	service "github.com/barisaydogdu/PostgreSQLwithGo/service/user"
	"github.com/joho/godotenv"
)

var (
	method    string
	typ       string
	id        int
	firstName string
	lastName  string
	number    int
	balance   int
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	flag.StringVar(&typ, "typ", "", "")
	flag.StringVar(&method, "method", "", "")
	flag.StringVar(&firstName, "firstname", "", "")     //firstName parameter
	flag.StringVar(&lastName, "lastname", "", "")       //lastName parameter
	flag.IntVar(&id, "id", 1, "")                       // id parameter
	flag.IntVar(&number, "number", 0, "number flag")    //number parameter
	flag.IntVar(&balance, "balance", 0, "balance flag") //balance parameter
	flag.Parse()

	conn, err := pstgr.ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	hndlrs := handlers.NewHandlers()

	hndlrs.Start()

	//call repository
	repo := repository.NewUserRepository(conn, ctx)
	//call service
	userService := service.NewUserService(repo)

	err = userService.UserActionService(typ, method, id, firstName, lastName, number, balance)
	if err != nil {
		log.Fatalf("Error handling user action: %v", err)
	}

	fmt.Println("Successfully connected to the database")

	hndlrs.Stop()
	conn.Close()
	cancel()
}
