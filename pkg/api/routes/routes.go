package routes

import (
	"github.com/kenowi-dev/hsp-auto-login/pkg/api/routes/hsp_auto_login"
	"github.com/kenowi-dev/hsp-auto-login/pkg/api/routes/index"
	"github.com/kenowi-dev/hsp-auto-login/pkg/api/routes/static"
	"github.com/kenowi-dev/hsp-auto-login/pkg/hsp"
	"net/http"
)

type Routes interface {
	Init(mux *http.ServeMux)
}

type routes struct {
	hspAutoLogin hsp_auto_login.Handler
	index        index.Handler
	static       static.Handler
}

func New(hsp hsp.HSP) Routes {
	return &routes{
		hspAutoLogin: hsp_auto_login.New(hsp),
		index:        index.New(hsp),
		static:       static.New(),
	}
}

func (r *routes) Init(mux *http.ServeMux) {
	mux.HandleFunc("/static/", r.static.Get)
	mux.HandleFunc("/", r.index.Get)
	mux.HandleFunc("/login", r.hspAutoLogin.Post)
}
