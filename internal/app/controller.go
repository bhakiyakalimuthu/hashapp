package app

import (
	"github.com/go-chi/chi"
	"go.uber.org/zap"
	"net/http"

	"github.com/bhakiyakalimuthu/hashapp/pkg"
)

var _ pkg.Controller = (*Controller)(nil)
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
	router.Get("/home", c.home)
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