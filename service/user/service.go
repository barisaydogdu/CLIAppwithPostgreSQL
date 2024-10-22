package user

import (
	"fmt"
	"log"

	userdomain "github.com/barisaydogdu/PostgreSQLwithGo/domain/user"
	repository "github.com/barisaydogdu/PostgreSQLwithGo/repository/user"
)

type Service interface {
	UserService(user *userdomain.User) error
	UserActions(user *userdomain.User) error
	GetAllUser() error
	GetUser(id int64) error
	CreateUser(firstName string, lastName string, number int, balance int) (*userdomain.User, error)
	UpdateUser(id int64, firstName string, lastName string, number int, balance int) error
	DeleteUser(id int64) error
}

type service struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) Service {
	return &service{repo: repo}
}

func (us *service) UserService(user *userdomain.User) error {
	switch user.Typ {
	case "user":
		return us.UserActions(user)
	default:
		log.Printf("Unkown type %s", user.Typ)
		return nil
	}
}

// func (us *userService) UserActions(method string, id int, firstName string, lastName string, number int, balance int) error {
func (us *service) UserActions(user *userdomain.User) error {
	switch user.Method {
	case "all":
		return us.GetAllUser()
	case "get":
		return us.GetUser(user.Id)
	case "create":
		_, err := us.CreateUser(user.FirstName, user.LastName, user.Number, user.Balance)
		return err
	case "update":
		return us.UpdateUser(user.Id, user.FirstName, user.LastName, user.Number, user.Balance)
	case "delete":
		return us.DeleteUser(user.Id)
	default:
		return nil
	}
}

func (us *service) GetAllUser() error {
	users, err := us.repo.GetAllUsers()
	if err != nil {
		fmt.Errorf("Something went wrong with get all users %v", err)
		return err
	}
	us.repo.PrintUsers(users)
	return nil
}

func (us *service) GetUser(id int64) error {
	user, err := us.repo.GetUserByID(id)
	if err != nil {
		fmt.Errorf("Something went wrong with get user by ID %v", err)
		return err
	}
	us.repo.PrintUser(user)
	return nil
}

func (us *service) CreateUser(firstName string, lastName string, number int, balance int) (*userdomain.User, error) {
	//burası da patladı o yüzden domaini de Repositorye sıktım
	var err error
	newUser := &userdomain.User{
		FirstName: firstName,
		LastName:  lastName,
		Number:    number,
		Balance:   balance,
	}

	err = us.repo.CreateUser(newUser)
	if err != nil {
		fmt.Errorf("Something went wrong with insert user %v", err)
		return nil, err
	}

	log.Printf("Succesffuly created user with ID: %d", newUser.Id)

	return newUser, nil
}

func (us *service) UpdateUser(id int64, firstName string, lastName string, number int, balance int) error {
	updatedUser := userdomain.User{
		FirstName: firstName,
		LastName:  lastName,
		Number:    number,
		Balance:   balance,
	}
	err := us.repo.UpdateUser(&updatedUser, id)
	if err != nil {
		fmt.Errorf("Something went wrong with update user %d", err)
		return err
	}
	log.Printf("Successfully updated user with ID: %d", id)
	return nil
}

func (us *service) DeleteUser(id int64) error {
	err := us.repo.DeleteUser(id)
	if err != nil {
		fmt.Errorf("Something went wrong with delete user %d", err)
		return err
	}
	return nil
}
