package handlers

import (
	"context"

	"database/sql"

	"github.com/barisaydogdu/PostgreSQLwithGo/handlers/cli"
	userRepository "github.com/barisaydogdu/PostgreSQLwithGo/repository/user"
	userService "github.com/barisaydogdu/PostgreSQLwithGo/service/user"
)

type Handlers interface {
	Start() error
	Stop() error
}

type handlers struct {
	ctx            context.Context
	userRepository userRepository.UserRepository
	userService    userService.UserServiceInterface
	cli            cli.CLI
}

func NewHandlers(ctx context.Context, db *sql.DB) Handlers {
	userRepo := userRepository.NewUserRepository(ctx, db)
	userSrv := userService.NewUserService(userRepo)

	return &handlers{
		ctx:            ctx,
		userRepository: userRepo,
		userService:    userSrv,
		cli:            cli.NewCLI(ctx, userSrv),
	}
}

func (h *handlers) Start() error {
	if h.cli != nil {
		if err := h.cli.Start(); err != nil {
			return err
		}
	}

	return nil
}

func (h *handlers) Stop() error {
	return nil
}
