package user

import (
	"context"
	"database/sql"
	"testing"
	"time"

	_ "github.com/lib/pq"

	userDomain "github.com/barisaydogdu/PostgreSQLwithGo/domain/user"
)

var (
	db  *sql.DB
	ctx context.Context
)

func init() {
	TestLoadEnv()
	var err error
	db, err = ConnectToDB()
	if err != nil {
		panic(err)
	}
	ctx = context.Background()

}

func TestCreateUser(t *testing.T) {
	repo := NewUserRepository(ctx, db)

	newUser := &userDomain.User{
		FirstName: "Barış",
		LastName:  "Aydoğdu",
		Number:    232,
		Balance:   3473,
		CreatedAt: time.Now(),
	}
	err := repo.CreateUser(newUser)
	if err != nil {
		t.Errorf("eailed to Create User %v %v ", err, newUser.Id)
	}
	count, err := repo.Count()
	if err != nil {
		t.Errorf("error with scaning count")
	}

	if count != 1 {
		t.Errorf("expected count to be 1, got %d", count)
	}
}

func TestGetAllUsers(t *testing.T) {
	defer db.Exec("truncate account restart identity cascade")

	repo := NewUserRepository(ctx, db)

	users, err := repo.GetAllUsers()
	if err != nil {
		t.Errorf("Something went wrong with get all user %v", err)
	}

	expectedUserCount := 1
	if len(users) != expectedUserCount {
		t.Errorf("Expected %d users but got %d", expectedUserCount, len(users))
	}
}

func TestGetUserByID(t *testing.T) {
	repo := NewUserRepository(ctx, db)

	newUser := userDomain.User{
		FirstName: "Simge",
		LastName:  "Okumuş",
		Number:    243,
		Balance:   3948,
		CreatedAt: time.Now(),
	}
	err := repo.CreateUser(&newUser)

	if err != nil {
		t.Errorf("There is something error with creating user")
	}
	user, err := repo.GetUserByID(newUser.Id)
	if err != nil {
		t.Errorf("Error when get user by id %v", err)
	}
	if user.FirstName != newUser.FirstName || user.LastName != newUser.LastName || user.Number != newUser.Number || user.Balance != newUser.Balance {
		t.Errorf("fetched users doesnt match expected %v got:%v", user, newUser)
	}
}

func TestUpdateUser(t *testing.T) {
	repo := NewUserRepository(ctx, db)

	newUser := &userDomain.User{
		FirstName: "Zeynep",
		LastName:  "Karakan",
		Number:    38364,
		Balance:   2823,
		CreatedAt: time.Now(),
	}
	err := repo.CreateUser(newUser)
	if err != nil {
		t.Errorf("there is some error with creating user%v", err)
	}
	var userID int64
	newUser.Id = userID
	newUser.FirstName = "Updated Name"
	newUser.LastName = "Updated LastName"
	newUser.Number = 9292
	newUser.Balance = 8379

	repo.UpdateUser(newUser, newUser.Id)

	if err != nil {
		t.Errorf("unexpected error while updating user %v", err)
	}

	var updatedUser userDomain.User

	user, err := repo.GetUserByID(newUser.Id)

	if err != nil {
		t.Errorf("error fetching uğdated user %v", user)
	}

	if updatedUser.FirstName != newUser.FirstName || updatedUser.LastName != newUser.LastName || updatedUser.Number != newUser.Number || updatedUser.Balance != newUser.Balance {
		t.Errorf("user details does not match after updated got:%v expected %v", updatedUser, newUser)
	}
}

func TestDeleteUser(t *testing.T) {
	repo := NewUserRepository(ctx, db)

	user := userDomain.User{
		FirstName: "Barış",
		LastName:  "Aydoğdu",
		Number:    212,
		Balance:   3338,
		CreatedAt: time.Now(),
	}
	mp := map[string]interface{}{
		"columnName": "id",
		"value":      user.Id,
	}
	err := repo.CreateUser(&user)
	if err != nil {
		t.Errorf("There is some error with creating user %v", err)
	}
	err = repo.DeleteUser(user.Id)
	if err != nil {
		t.Errorf("unexpected error while deleting user: %s", err)
	}

	count, err := repo.CountRows(mp, user.Id)
	if err != nil {
		t.Errorf("error checking user existence: %s", err)
	}
	if count != 0 {
		t.Errorf("Exprected user to be delete, but found %d users with ID %d", count, user.Id)
	}
}
