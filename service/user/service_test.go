package user

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	userdomain "github.com/barisaydogdu/PostgreSQLwithGo/domain/user"
	userRepository "github.com/barisaydogdu/PostgreSQLwithGo/repository/user"
)

var (
	userservice Service
	repo        userRepository.UserRepository
	db          *sql.DB
)

func init() {
	var err error
	LoadEnv()
	db, err = ConnectToDB()
	if err != nil {
		fmt.Errorf("Something went wrong with database %v", err)
	}
	ctx := context.Background()
	repo = userRepository.NewUserRepository(ctx, db)

	userservice = NewUserService(repo)
}

func TestCreateUser(t *testing.T) {
	firstName := "testFirstName"
	lastName := "testLastName"
	number := 343
	balance := 343

	mp := make(map[string]interface{})
	mp["columnName"] = "first_name"
	mp["columnName"] = "last_name"
	mp["value"] = firstName
	mp["value"] = lastName

	user, err := userservice.CreateUser(firstName, lastName, number, balance)
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	if user == nil {
		t.Error("user is nil")
	}

	if user.Id == 0 {
		t.Error("user not inserted")
	}

	count, err := repo.CountRows(mp, user.Id)
	if err != nil {
		fmt.Errorf("cannot count rows %v", err)
	}
	if err != nil {
		t.Errorf("error failed to scan row %v", err)
	}

	if count == 0 {
		t.Errorf("error user doesnt created %v", count)
	}

}

func TestGetAllUsers(t *testing.T) {
	for i := 0; i < 3; i++ {
		TestCreateUser(t)
	}
	err := userservice.GetAllUser()
	if err != nil {
		t.Errorf("error with fetching all users %v", err)
	}
	count, err := repo.Count()
	if err != nil {
		t.Errorf("error with counting rows %v", err)
	}

	if count < 1 {
		t.Errorf("error: expected row count is %v", count)
	}

}

func TestGetUser(t *testing.T) {
	firstName := "testFirstName"
	lastName := "testLastName"
	number := 343
	balance := 343
	//err := repo.CreateUser(&user)
	user, err := userservice.CreateUser(firstName, lastName, number, balance)
	if err != nil {
		t.Errorf("error with scaning user id %v", err)
	}
	var firstNameResult, lastNameResult string
	var numberResult, balanceResult int

	fetchedUser, err := repo.GetUserByID(user.Id)
	if err != nil {
		t.Errorf("error when get user by id %v", err)
	}

	firstNameResult = user.FirstName
	lastNameResult = user.LastName
	numberResult = user.Number
	balanceResult = user.Balance

	if fetchedUser.FirstName != firstNameResult || fetchedUser.LastName != lastNameResult || fetchedUser.Number != numberResult || fetchedUser.Balance != balanceResult {
		t.Error("users doesnt match")
	}

	err = userservice.DeleteUser(user.Id)
	if err != nil {
		t.Errorf("error while deleting user %v", err)
	}
}

func TestUpdateUser(t *testing.T) {
	firstName := "testFirstName"
	lastName := "testLastName"
	number := 343
	balance := 343

	_, err := userservice.CreateUser(firstName, lastName, number, balance)
	if err != nil {
		fmt.Errorf("error when created user %v", err)
	}
	var userID int64 = 44
	updatedFirstName := "UpdatedFirstName"
	updatedLastName := "UpdatedLastName"
	updatedNumber := 999
	updatedBalance := 8888

	err = userservice.UpdateUser(userID, updatedFirstName, updatedLastName, updatedNumber, updatedBalance)
	if err != nil {
		t.Errorf("There is some error with update user: %v", err)
	}
	var firstNameResult, lastNameResult string
	var numberResult, balanceResult int

	err = userservice.GetUser(userID)
	if err != nil {
		t.Errorf("failed to fetching data from account %v", err)
	}

	if firstNameResult != updatedFirstName || lastNameResult != updatedLastName || numberResult != updatedNumber || balanceResult != updatedBalance {
		t.Errorf("user details does not match %v", err)
	}

	_, err = db.Exec("DELETE from account")
	if err != nil {
		t.Errorf("error while deleting table %v", err)
	}
}

func TestDeleteUser(t *testing.T) {
	var err error
	firstName := "TestFirstName"
	lastName := "TestLastName"
	number := 999
	balance := 989

	var newUser *userdomain.User
	newUser, err = userservice.CreateUser(firstName, lastName, number, balance)
	if err != nil {
		t.Errorf("something went wrong with creating user %v", err)
	}

	err = userservice.DeleteUser(newUser.Id)
	if err != nil {
		t.Errorf("something went wrong with deleting user %v", err)
	}

	count, err := repo.Count()
	if err != nil {
		t.Errorf("error with scaning count %v", err)
	}

	if count > 0 {
		t.Errorf("count should be 0 but got: %v", count)
	}

}
