// repository sdfsfdsdfdsfd
package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	userdomain "github.com/barisaydogdu/PostgreSQLwithGo/domain/user"
)

// UserRepository sasdadsasdasdad
type UserRepository interface {
	CreateUser(user *userdomain.User) error
	Count() (int64, error)
	CountRows(mp map[string]interface{}, id int64) (int64, error)
	GetAllUsers() ([]*userdomain.User, error)
	GetUserByID(id int64) (*userdomain.User, error)
	UpdateUser(usr *userdomain.User, id int64) error
	DeleteUser(id int64) error
	PrintUser(*userdomain.User)
	PrintUsers([]*userdomain.User)
}

type userRepository struct {
	ctx context.Context
	Db  *sql.DB
}

func NewUserRepository(ctx context.Context, db *sql.DB) UserRepository {
	return &userRepository{
		Db:  db,
		ctx: ctx,
	}
}

// GetAllUsers sdfsfdsdfsdfsdfsdffds
// sdadsasd
func (u *userRepository) CreateUser(usr *userdomain.User) error {
	inctx, cancel := context.WithTimeout(u.ctx, 5*time.Second)
	defer cancel()
	err := u.Db.QueryRowContext(inctx, "INSERT INTO account (first_name,last_name,number,balance,created_at) VALUES ($1,$2,$3,$4,$5) RETURNING id",
		usr.FirstName, usr.LastName, usr.Number, usr.Balance, time.Now()).Scan(&usr.Id)
	if err != nil {
		log.Printf("Add User: %v", err)
		return err
	}

	log.Println("Succesfully created a user")

	return nil
}

func (u *userRepository) Count() (int64, error) {
	count, err := u.count(nil, 0)
	if err != nil {
		//
		return 0, err
	}

	return count, nil
}

func (u *userRepository) CountRows(mp map[string]interface{}, id int64) (int64, error) {
	count, err := u.count(mp, id)
	if err != nil {
		//
		return 0, err
	}

	return count, nil
}

func (u *userRepository) count(mp map[string]interface{}, id int64) (int64, error) {
	var wc string
	var args []interface{}
	if mp != nil {
		val, ok := mp["columnName"]
		if !ok {
			return 0, errors.New("invaild name first if")
		}

		columnName := val.(string)

		// if columnName != "first_name" || columnName != "last_name" || columnName != "number" || columnName != "balance" {
		// 	return 0, errors.New("invaild name")
		// }

		if !(columnName == "id" || columnName == "first_name" || columnName == "last_name" || columnName == "number" || columnName == "balance") {
			return 0, errors.New("invalid name second if")
		}

		value, ok := mp["value"]
		if !ok {
			return 0, errors.New("invaild name with value")
		}
		wc = "WHERE " + columnName + " = $" + strconv.Itoa(len(args)+1)
		args = append(args, value)
	} else if id > 0 {
		wc = "WHERE id = $" + strconv.Itoa(len(args)+1)
	}
	query := "SELECT COUNT (*) FROM account " + wc

	var count int64
	if err := u.Db.QueryRow(query, args...).Scan(&count); err != nil {
		return 0, fmt.Errorf("query database error %v", err)
	}

	return count, nil
}

func (u *userRepository) GetAllUsers() ([]*userdomain.User, error) {
	rows, err := u.Db.Query("SELECT * FROM account")

	if err != nil {
		fmt.Errorf("There is something error with get all users %v", err)
	}
	defer rows.Close()

	var users []*userdomain.User

	for rows.Next() {
		var user userdomain.User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Number, &user.Balance, &user.CreatedAt); err != nil {
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

func (u *userRepository) GetUserByID(id int64) (*userdomain.User, error) {
	row := u.Db.QueryRow("SELECT * FROM account WHERE id=$1", id)

	var usr userdomain.User

	if err := row.Scan(&usr.Id, &usr.FirstName, &usr.LastName, &usr.Number, &usr.Balance, &usr.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no user found with ID %d", id)
		}

		return nil, fmt.Errorf("failed to scan user %v", err)
	}
	log.Printf("Succesfully retrieved user with ID %d", id)

	return &usr, nil
}

func (u *userRepository) UpdateUser(usr *userdomain.User, id int64) error {
	result, err := u.Db.Exec("UPDATE account SET first_name=$1, last_name=$2, number=$3, balance=$4 WHERE id=$5",
		usr.FirstName, usr.LastName, usr.Number, usr.Balance, id)
	if err != nil {
		log.Printf("Error when updating a user: %v\n", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("update User: %v", err)
	}

	if rowsAffected == 0 {
		return errors.New("ftygdfgdsfs")
	}

	return err
}

func (u *userRepository) DeleteUser(id int64) error {
	result, err := u.Db.Exec("DELETE FROM account WHERE id= $1", id)
	if err != nil {
		fmt.Errorf("Error when delete a user: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("delete User: %v", err)
	}

	if rowsAffected == 0 {
		return errors.New("ftygdfgdsfs")
	}

	return nil
}

func (u *userRepository) PrintUsers(user []*userdomain.User) {
	for _, usr := range user {
		log.Printf("User ID: %d, Name: %s,LastName:%s Number: %d, Balance: %d, Created At: %s",
			usr.Id, usr.FirstName, usr.LastName, usr.Number, usr.Balance, usr.CreatedAt.Format("2006-01-02 15:04:05"))
	}
}
func (u *userRepository) PrintUser(user *userdomain.User) {
	if user == nil {
		log.Println("User not found")
		return
	}
	log.Printf("User ID: %d, Name: %s %s, Number: %d, Balance: %d, Created At: %s",
		user.Id, user.FirstName, user.LastName, user.Number, user.Balance, user.CreatedAt.Format("2006-01-02 15:04:05"))
}
