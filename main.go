package main

import (
	"github.com/bhakiyakalimuthu/hashapp/internal/app"
	"github.com/voi-go/svc"
	chiworker "github.com/voiapp/svc-workers/chi"
	"go.uber.org/zap"
)


func main()  {
	s, err := svc.New("hashapp","snap-shot")
	svc.MustInit(s,err)
	ctrl := app.
	rest := chiworker.New()
}

