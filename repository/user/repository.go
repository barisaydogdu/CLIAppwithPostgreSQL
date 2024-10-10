package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type UserRepository interface {
	GetAllUsers() ([]*User, error)
	GetUserByID(id int) (*User, error)
	InsertUser(user *User) (int64, error)
	UpdateUser(usr *User, id int) (int, error)
	DeleteUser(id int) (int64, error)
	PrintUser(*User)
	PrintUsers([]*User)
}

type userRepository struct {
	Db  *sql.DB
	ctx context.Context
}

func NewUserRepository(ctx context.Context, db *sql.DB) UserRepository {
	return &userRepository{
		Db:  db,
		ctx: ctx,
	}
}

func (u *userRepository) GetAllUsers() ([]*User, error) {
	rows, err := u.Db.Query("SELECT * FROM account")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []*User

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.First_name, &user.Last_name, &user.Number, &user.Balance, &user.Created_at); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		users = append(users, &user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occured during rows iteration: %w", err)
	}
	log.Printf("Successfully retrieved %d users", len(users))
	return users, nil
}

func (u *userRepository) GetUserByID(id int) (*User, error) {
	row := u.Db.QueryRow("SELECT * FROM account WHERE id=$1", id)

	var usr User

	if err := row.Scan(&usr.Id, &usr.First_name, &usr.Last_name, &usr.Number, &usr.Balance, &usr.Created_at); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no user found with ID %d", id)
		}

		return nil, fmt.Errorf("failed to scan user %v", err)
	}
	log.Printf("Succesfully retrieved user with ID %d", id)

	return &usr, nil
}

func (u *userRepository) InsertUser(usr *User) (int64, error) {
	inctx, cancel := context.WithTimeout(u.ctx, 5*time.Second)
	defer cancel()
	var id int64
	err := u.Db.QueryRowContext(inctx, "INSERT INTO account (first_name,last_name,number,balance,created_at) VALUES ($1,$2,$3,$4,$5) RETURNING id",
		usr.First_name, usr.Last_name, usr.Number, usr.Balance, time.Now()).Scan(&id)
	if err != nil {
		log.Fatal("Add User: %v", err)
	}
	return id, nil
}

func (u *userRepository) DeleteUser(id int) (int64, error) {
	result, err := u.Db.Exec("DELETE FROM account WHERE id= $1", id)
	if err != nil {
		log.Fatalf("Error when delete a user: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("Delete User: %v", err)
	}
	return rowsAffected, nil
}

func (u *userRepository) UpdateUser(usr *User, id int) (int, error) {
	err := u.Db.QueryRow("UPDATE account SET first_name=$1, last_name=$2, number=$3, balance=$4 WHERE id=$5 RETURNING id",
		usr.First_name, usr.Last_name, usr.Number, usr.Balance, id).Scan(&id)
	if err != nil {
		log.Fatalf("Error when updating a user: %v", err)
	}
	return id, err
}

func (u *userRepository) PrintUsers(user []*User) {
	for _, usr := range user {
		log.Printf("User ID: %d, Name: %s,LastName:%s Number: %d, Balance: %d, Created At: %s",
			usr.Id, usr.First_name, usr.Last_name, usr.Number, usr.Balance, usr.Created_at.Format("2006-01-02 15:04:05"))
	}
}

func (u *userRepository) PrintUser(user *User) {
	if user == nil {
		log.Println("User not found")
		return
	}
	log.Printf("User ID: %d, Name: %s %s, Number: %d, Balance: %d, Created At: %s",
		user.Id, user.First_name, user.Last_name, user.Number, user.Balance, user.Created_at.Format("2006-01-02 15:04:05"))
}
