package routes

import (
	"github.com/kenowi-dev/hsp-auto-login/pkg/api/routes/hsp"
	"github.com/kenowi-dev/hsp-auto-login/pkg/api/routes/static"
	"net/http"
)

type Routes interface {
	Init(mux *http.ServeMux)
}

type routes struct {
	index  hsp.Handler
	static static.Handler
}

func New() (Routes, error) {
	idx, err := hsp.New()
	if err != nil {
		return nil, err
	}
	return &routes{
		index:  idx,
		static: static.New(),
	}, nil
}

func (r *routes) Init(mux *http.ServeMux) {
	mux.HandleFunc("/static/", r.static.Get)
	mux.HandleFunc("/", r.index.Get)
	mux.HandleFunc("/courses", r.index.Courses)
	mux.HandleFunc("/registration", r.index.Registrations)
	mux.HandleFunc("/dev", r.index.Dev)
	mux.HandleFunc("/login", r.index.Login)
}
