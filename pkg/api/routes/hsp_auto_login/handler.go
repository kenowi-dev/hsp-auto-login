package hsp_auto_login

import (
	"github.com/kenowi-dev/hsp-auto-login/pkg/api/files"
	"github.com/kenowi-dev/hsp-auto-login/pkg/hsp"
	"html/template"

	"log"
	"net/http"
)

type Handler interface {
	Post(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	hsp hsp.HSP
}

func New(hsp hsp.HSP) Handler {
	return &handler{
		hsp: hsp,
	}
}

func (h *handler) Post(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		data := hsp.Data{
			Sport:        r.PostFormValue("sport"),
			CourseNumber: r.PostFormValue("courseId"),
			Email:        r.PostFormValue("email"),
			Password:     r.PostFormValue("password"),
		}

		err := h.hsp.AutoLogin(&data)
		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}

		tmpl := template.Must(template.ParseFS(files.Templates, "templates/fragments/auto_login.html"))
		err = tmpl.ExecuteTemplate(w, "auto_login.html", h.hsp.GetRegistrationListeners())
		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}
	}
}
