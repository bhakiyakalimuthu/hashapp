package app

import (
	"github.com/go-chi/chi"
	chiworker "github.com/voiapp/svc-workers/chi"
	"go.uber.org/zap"
	"net/http"
)

var _ chiworker.Controller = (*Controller)(nil)
type Controller struct {
	logger *zap.Logger
}

func NewController(logger *zap.Logger) *Controller {
	return &Controller{logger:logger}
}

func (c *Controller) Init(logger *zap.Logger) error {
	c.logger = logger
	return nil
}

func (c *Controller) SetupRouter(router chi.Router) error {
	router.Get("/", c.home)
	if err := c.SetupRouter(router);err !=nil {
		c.logger.Error("setup router failed in controller",zap.Error(err))
		return err
	}
	return nil
}

func (c *Controller) Terminate() error {
	return nil
}

func (c *Controller) home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if _,err:= w.Write([]byte(`hello, hashapp home`));err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	return

}