package cli

import (
	"context"

	userService "github.com/barisaydogdu/PostgreSQLwithGo/service/user"
	"github.com/spf13/cobra"
)

type CLI interface {
	Start() error
	Stop() error
}

type cli struct {
	ctx         context.Context
	userService userService.Service
	root        *cobra.Command
}

func NewCLI(ctx context.Context, userSrv userService.Service) CLI {
	c := &cli{
		ctx:         ctx,
		userService: userSrv,
	}

	c.root = &cobra.Command{
		Use:   "PostgreSqlwithGo",
		Short: "CLI application to manage users",
		Long:  "A CLI application to perform CRUD operations on users PostgreSQL",
	}
	c.root.AddCommand(NewUserHandler(userSrv))

	return c
}

func (c *cli) Start() error {
	if err := c.root.Execute(); err != nil {
		return err
	}

	return nil
}

func (c *cli) Stop() error {
	return nil
}
