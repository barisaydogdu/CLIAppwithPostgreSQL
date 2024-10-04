package pkg

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type User struct {
	Id         int
	First_name string
	Last_name  string
	Number     int
	Balance    int
	Created_at time.Time
}

type userRepository struct {
	DB *sql.DB
}

func NewUserRepository(db sql.DB) *userRepository {
	return &userRepository{
		DB: &db,
	}
}

func (n *userRepository) GetAllUsers() ([]User, error) {
	rows, err := n.DB.Query("SELECT * FROM public.account")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.First_name, &user.Last_name, &user.Number, &user.Balance, &user.Created_at); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occured during rows iteration: %w", err)
	}
	log.Printf("Successfully retrieved %d users", len(users))
	return users, nil
}
