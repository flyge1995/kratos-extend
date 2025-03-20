package crontab

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/robfig/cron/v3"
)

var _ transport.Server = new(Server)

type Server struct {
	crontab *cron.Cron
}

func NewServer(crontab *cron.Cron) *Server {
	return &Server{crontab: crontab}
}

func (s *Server) Start(ctx context.Context) error {
	s.crontab.Run()
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	select {
	case <-ctx.Done():
	case <-s.crontab.Stop().Done():
	}
	return nil
}
