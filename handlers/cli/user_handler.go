package cli

import (
	"log"

	userdomain "github.com/barisaydogdu/PostgreSQLwithGo/domain/user"
	userservice "github.com/barisaydogdu/PostgreSQLwithGo/service/user"
	"github.com/microcosm-cc/bluemonday"
	"github.com/spf13/cobra"
)

type userHandler struct {
	service userservice.Service
}

func NewUserHandler(userSrv userservice.Service) *cobra.Command {
	s := &userHandler{
		service: userSrv,
	}
	//Sor
	user := &userdomain.User{}

	//Bluemonday
	policy := bluemonday.UGCPolicy()

	cmd := &cobra.Command{
		Use:   "useraction",
		Short: "Perform CRUD operations on users",
		Run: func(cmd *cobra.Command, args []string) {
			if err := s.service.UserService(user); err != nil {
				user.FirstName = policy.Sanitize(user.FirstName)
				user.LastName = policy.Sanitize(user.LastName)
				user.Method = policy.Sanitize(user.Method)
				user.Typ = policy.Sanitize(user.Typ)
				log.Println("There is something wrong", err)
			}
		},
	}

	cmd.Flags().StringVar(&user.FirstName, "firstname", "", "")
	cmd.Flags().StringVar(&user.LastName, "lastname", "", "")
	cmd.Flags().StringVar(&user.Method, "method", "", "")
	cmd.Flags().StringVar(&user.Typ, "typ", "", "")
	cmd.Flags().Int64Var(&user.Id, "id", 1, "")
	cmd.Flags().IntVar(&user.Number, "number", 1, "")
	cmd.Flags().IntVar(&user.Balance, "balance", 1, "")

	return cmd
}
