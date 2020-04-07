package main

import (
	"github.com/voi-go/svc"
	"go.uber.org/zap"
	"github.com/go-chi/chi"
	"golang.org/x/net/context"
	"net"
	"net/http"
	"go.opencensus.io/plugin/ochttp"
	"strconv"
)

var _ svc.Worker = (*ChiWorker)(nil)

type ChiWorker struct {
	port int
	logger *zap.Logger
	router chi.Router
	controller Controller
	server  *http.Server
}

func NewWorker(ctrl Controller)  *ChiWorker {
	return &ChiWorker{
		port:       8080,
		controller: ctrl,
	}
}
func (c *ChiWorker) Init(logger *zap.Logger) error {
	c.logger = logger
	c.router = chi.NewRouter()
	if err := c.controller.Init(c.logger); err !=nil {
		logger.Error("failed to init controller")
		return err
	}
	c.server = &http.Server{
		Addr: net.JoinHostPort("", strconv.Itoa(c.port)),
		Handler:     &ochttp.Handler{
			Handler:   c.router,
		}      ,
	}
	r := chi.NewRouter()
	if err:= c.controller.SetupRouter(r);err !=nil {
		logger.Error("failed to setup router")
		return err
	}
	return nil
}

func (c *ChiWorker) Run() error {
	c.logger.Info("HTTP server running",zap.String("host name",c.server.Addr))
	if err := c.server.ListenAndServe();err!=nil && err != http.ErrServerClosed {
		c.logger.Error("HTTP server failed to RUN")
		return err
	}
	return nil
}

func (c *ChiWorker) Terminate() error {
	if err := c.server.Shutdown(context.Background());err != nil {
		c.logger.Warn("server shutdown failed")
	}
	return c.controller.Terminate()
}

type Controller interface {
	Init(logger *zap.Logger) error
	SetupRouter(router chi.Router) error
	Terminate() error
}
func main()  {
	s,err := svc.New()
}

