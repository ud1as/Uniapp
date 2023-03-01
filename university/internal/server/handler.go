package server

import (
	"github.com/Studio56School/university/internal/handler"
	"go.uber.org/zap"
)

type ServerHandlers struct {
	university handler.IHandler
}

func newHandlers(svc *ServerServices, logger *zap.Logger) (*ServerHandlers, error) {

	h := &ServerHandlers{}
	h.university = handler.NewHandler(svc.Srv, logger)

	return h, nil
}
