package crontab

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/robfig/cron/v3"
)

var _ transport.Server = new(Server)

type Server struct {
	crontab *cron.Cron
	logger  *log.Helper
}

func NewServer(crontab *cron.Cron, logger log.Logger) *Server {
	return &Server{crontab: crontab, logger: log.NewHelper(logger)}
}

func (s *Server) Start(ctx context.Context) error {
	s.logger.Info("[crontab]server start")
	s.crontab.Run()
	s.logger.Info("[crontab]server stop")
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.logger.Info("[crontab]server stopping")
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-s.crontab.Stop().Done():
	}
	return nil
}
