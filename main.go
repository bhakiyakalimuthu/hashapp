package main

import (
	"github.com/voi-go/svc"
	"go.uber.org/zap"
	"github.com/go-chi/chi"
)

var _ svc.Worker = (*ChiWorker)(nil)

type ChiWorker struct {
	logger *zap.Logger
	router chi.Router
	controller Controller
}

func NewWorker()  *ChiWorker {
	return &ChiWorker{}
}
func (c ChiWorker) Init(logger *zap.Logger) error {
	c.logger = logger
	c.router = chi.NewRouter()
	if err := c.controller.Init(c.logger); err !=nil {
		logger.Error("failed to init controller")
		return err
	}
	r := chi.NewRouter()
	if err:= c.controller.SetupRouter(r);err !=nil {
		logger.Error("failed setup router")
		return err
	}
	return nil
}

func (c ChiWorker) Run() error {
	panic("implement me")
}

func (c ChiWorker) Terminate() error {
	panic("implement me")
}

type Controller interface {
	Init(logger *zap.Logger) error
	SetupRouter(router chi.Router) error
	Terminate() error
}
func main()  {
	s,err := svc.New()
}

