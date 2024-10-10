package cli

import (
	userservice "github.com/barisaydogdu/PostgreSQLwithGo/service/user"
	"github.com/spf13/cobra"
)

type userHandler struct {
	service *userservice.UserService
}

func NewUserHandler(userSrv userservice.UserService) *cobra.Command {
	s := &userHandler{
		service: &userSrv,
	}
	var (
		id        int
		typ       string
		firstName string
		lastName  string
		method    string
		number    int
		balance   int
	)

	cmd := &cobra.Command{
		Use:   "useraction",
		Short: "Perform CRUD operations on users",
		Run: func(cmd *cobra.Command, args []string) {
			s.service.UserActionService(typ, method, id, firstName, lastName, number, balance)
		},
	}

	cmd.Flags().StringVar(&firstName, "firstname", "", "")
	cmd.Flags().StringVar(&lastName, "lastname", "", "")
	cmd.Flags().StringVar(&method, "method", "", "")
	cmd.Flags().StringVar(&typ, "typ", "", "")
	cmd.Flags().IntVar(&id, "id", 1, "")
	cmd.Flags().IntVar(&number, "number", 1, "")
	cmd.Flags().IntVar(&balance, "balance", 1, "")

	return cmd
}
