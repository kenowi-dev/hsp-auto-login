package api

import (
	"github.com/kenowi-dev/hsp-auto-login/pkg/api/routes"
	"github.com/kenowi-dev/hsp-auto-login/pkg/hsp"
	"log"
	"net/http"
	"strconv"
)

type Api interface {
	SetupAndRun(port int)
}

type api struct {
	hsp hsp.HSP
}

func New(hsp hsp.HSP) Api {
	return &api{hsp: hsp}
}

func (r *api) SetupAndRun(port int) {
	mux := http.NewServeMux()
	apiRoutes := routes.New(r.hsp)

	apiRoutes.Init(mux)

	log.Printf("App running on %s...\n", strconv.Itoa(port))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), mux))

}
