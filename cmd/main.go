package main

import (
	"fmt"
	"log"

	in "github.com/barisaydogdu/PostgreSQLwithGo/internal"
	pq "github.com/barisaydogdu/PostgreSQLwithGo/pkg"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connString := in.NewEnvDbConfig().ConnString()
	conn, err := in.ConnectToDB(connString)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	repo := pq.NewUserRepository(*conn)

	users, err := repo.GetAllUsers()
	if err != nil {
		log.Fatal("Something went wrong!")
	}

	for _, user := range users {
		log.Printf("User ID: %d, Name: %s,LastName:%s Number: %d, Balance: %d, Created At: %s",
			user.Id, user.First_name, user.Last_name, user.Number, user.Balance, user.Created_at.Format("2006-01-02 15:04:05"))
	}

	fmt.Println("Successfully connected to the database")

}
