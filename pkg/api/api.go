package api

import (
	"github.com/kenowi-dev/hsp-auto-login/pkg/api/routes"
	"log"
	"net/http"
	"strconv"
)

type Api interface {
	SetupAndRun(port int)
}

type api struct {
}

func New() Api {
	return &api{}
}

func (r *api) SetupAndRun(port int) {
	mux := http.NewServeMux()
	apiRoutes, err := routes.New()
	if err != nil {
		log.Fatal(err.Error())
	}

	apiRoutes.Init(mux)

	log.Printf("App running on %s...\n", strconv.Itoa(port))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), mux))
}
