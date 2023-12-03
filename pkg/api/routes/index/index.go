package index

import (
	"github.com/kenowi-dev/hsp-auto-login/pkg/api/files"
	"github.com/kenowi-dev/hsp-auto-login/pkg/hsp"
	"html/template"

	"net/http"
)

type Handler interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	hsp hsp.HSP
}

func New(hsp hsp.HSP) Handler {
	return &handler{hsp: hsp}
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFS(files.Templates, "templates/index.html"))

	tmpl = template.Must(tmpl.ParseFS(files.Templates, "templates/fragments/*.html"))
	err := tmpl.Execute(w, h.hsp.GetRegistrationListeners())
	if err != nil {
		panic(err)
	}
}
