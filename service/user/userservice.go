package user

import (
	"log"

	repository "github.com/barisaydogdu/PostgreSQLwithGo/repository/user"
)

type UserService interface {
	UserActionService(typ string, method string, id int, firstName string, lastName string, number int, balance int) error
	UserActions(method string, id int, firstName string, lastName string, number int, balance int) error
	getAllUserService() error
	getUserService(id int) error
	createUserService(firstName string, lastName string, number int, balance int) error
	updateUserService(id int, firstName string, lastName string, number int, balance int) error
	deleteUserService(id int) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (us *userService) UserActionService(typ string, method string, id int, firstName string, lastName string, number int, balance int) error {
	switch typ {
	case "user":
		return us.UserActions(method, id, firstName, lastName, number, balance)
	default:
		log.Printf("Unkown type %s", typ)
		return nil
	}

}

func (us *userService) UserActions(method string, id int, firstName string, lastName string, number int, balance int) error {
	switch method {
	case "all":
		return us.getAllUserService()
	case "get":
		return us.getUserService(id)
	case "create":
		return us.createUserService(firstName, lastName, number, balance)
	case "update":
		return us.updateUserService(id, firstName, lastName, number, balance)
	case "delete":
		return us.deleteUserService(id)
	default:
		return nil
	}
}

func (us *userService) getAllUserService() error {
	users, err := us.repo.GetAllUsers()
	if err != nil {
		log.Fatalf("Something went wrong with get all users %v", err)
		return err
	}
	us.repo.PrintUsers(users)
	return nil
}

func (us *userService) getUserService(id int) error {
	user, err := us.repo.GetUserByID(id)
	if err != nil {
		log.Fatalf("Something went wrong with get user by ID %v", err)
		return err
	}
	us.repo.PrintUser(user)
	return nil
}

func (us *userService) createUserService(firstName string, lastName string, number int, balance int) error {
	newUser := repository.User{
		First_name: firstName,
		Last_name:  lastName,
		Number:     number,
		Balance:    balance,
	}
	id, err := us.repo.InsertUser(&newUser)
	if err != nil {
		log.Fatalf("Something went wrong with insert user %v", err)
		return err
	}
	log.Printf("Succesffuly created user with ID: %d", id)
	return nil
}

func (us *userService) updateUserService(id int, firstName string, lastName string, number int, balance int) error {
	updatedUser := repository.User{
		First_name: firstName,
		Last_name:  lastName,
		Number:     number,
		Balance:    balance,
	}
	_, err := us.repo.UpdateUser(&updatedUser, id)
	if err != nil {
		log.Fatalf("Something went wrong with update user %d", err)
		return err
	}
	log.Printf("Successfully updated user with ID: %d", id)
	return nil
}

func (us *userService) deleteUserService(id int) error {
	rowsAffedted, err := us.repo.DeleteUser(id)
	if err != nil {
		log.Fatalf("Something went wrong with delete user %d", err)
		return err
	}
	log.Printf("Rows affected %d", rowsAffedted)
	return nil
}

// func UserActionHandler(repo repository.UserRepository, typ string, method string) error {
// 	switch typ {
// 	case "user":
// 		switch method {
// 		case "all":
// 			users, err := repo.GetAllUsers()
// 			if err != nil {
// 				log.Fatalf("Something went wrong with get all users %v", err)
// 			}
// 			repository.PrintUsers(users)

// 			break
// 		case "get":
// 			user, err := repo.GetUserByID(id)
// 			if err != nil {
// 				log.Fatalf("Something went wrong with get user by id %v", err)
// 			}
// 			repository.PrintUser(user)

// 			break
// 		case "create":
// 			newUser := repository.User{
// 				First_name: firstName,
// 				Last_name:  lastName,
// 				Number:     number,
// 				Balance:    balance,
// 			}
// 			id, err := repo.InsertUser(&newUser)
// 			if err != nil {
// 				log.Fatalf("Error adding user: %v", err)
// 			}
// 			log.Printf("Successfully added user with ID: %d", id)

// 			break
// 		case "update":
// 			updatedUser := repository.User{
// 				First_name: firstName,
// 				Last_name:  lastName,
// 				Number:     number,
// 				Balance:    balance,
// 			}
// 			id, err := repo.UpdateUser(&updatedUser, int64(id))
// 			if err != nil {
// 				log.Fatalf("Error updated user: %v", err)
// 			}
// 			log.Printf("Succesfully updated user with ID: %d", id)

// 			break
// 		case "delete":
// 			rowsAffected, err := repo.DeleteUser(id)
// 			if err != nil {
// 				log.Fatalf("Error deleting user: %v", err)
// 			}
// 			log.Printf("Rows Affected: %d", rowsAffected)

// 			break
// 		}
// 	}
// 	return nil
// }
