package handlers

import "context"

type Handlers interface {
	Start() error
	Stop() error
}

type handlers struct {
	ctx context.Context
}

func NewHandlers() Handlers {
	return &handlers{}
}

func (h *handlers) Start() error {
	return nil
}

func (h *handlers) Stop() error {
	return nil
}
