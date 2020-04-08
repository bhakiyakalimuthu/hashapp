package main

import (
	"github.com/bhakiyakalimuthu/hashapp/internal/app"
	"github.com/voi-go/svc"
	"go.uber.org/zap"
)


func main()  {
	s, err := svc.New("hashapp","snap-shot")
	svc.MustInit(s,err)
	ctrl := app.NewController(zap.L())
	rest := NewWorker(ctrl)
	s.AddWorker("http",rest)
	s.Run()
}

