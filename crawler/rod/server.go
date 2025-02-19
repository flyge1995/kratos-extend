package crawler

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/robfig/cron/v3"
)

var _ transport.Server = new(Server)

func NewServer(logger log.Logger) *Server {
	crontab := cron.New(cron.WithParser(cron.NewParser(
		cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
	)))

	return &Server{crontab: crontab, logger: log.NewHelper(logger)}
}

type Server struct {
	crontab *cron.Cron
	logger  *log.Helper
}

func (s *Server) Start(ctx context.Context) error {
	s.crontab.Run()
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.crontab.Stop()
	return nil
}

func (s *Server) AddTask(task Task) error {
	task, err := NewTaskDecorator(task)
	if err != nil {
		return err
	}
	_, err = s.crontab.AddFunc(task.Metadata().Crontab, func() {
		err := task.OnHandler(nil)
		if err != nil {
			s.logger.Error(err.Error())
		}
	})
	return err
}

//type Service struct {
//}
//
//func (s *Service) Run(task Task) {
//	go func() {
//		s.Start()
//	}()
//}
//
//func (s *Service) Start(task Task) error {
//	//
//
//	//
//
//	return nil
//}
