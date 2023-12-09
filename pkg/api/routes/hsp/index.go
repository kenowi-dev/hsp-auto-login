package hsp

import (
	"github.com/kenowi-dev/hsp-auto-login/pkg/api/files"
	"github.com/kenowi-dev/hspscraper"
	"html/template"
	"log"
	"time"

	"net/http"
)

type Handler interface {
	Get(w http.ResponseWriter, r *http.Request)
	Courses(writer http.ResponseWriter, request *http.Request)
	Registrations(writer http.ResponseWriter, request *http.Request)
	Dev(writer http.ResponseWriter, request *http.Request)
	Login(writer http.ResponseWriter, request *http.Request)
}

type handler struct {
	registrationHandler RegistrationHandler
	sports              []*hspscraper.Sport
}

func New() (Handler, error) {

	db, err := newDB[*Registration]("db.json")
	if err != nil {
		return nil, err
	}
	regHandler := newRegistrationHandler(db)

	sports, err := hspscraper.GetAllSports()
	if err != nil {
		return nil, err
	}

	return &handler{
		registrationHandler: regHandler,
		sports:              sports,
	}, nil
}

func (h *handler) Get(w http.ResponseWriter, _ *http.Request) {
	tmpl := template.Must(template.ParseFS(files.Templates, "templates/index.html"))
	tmpl = template.Must(tmpl.ParseFS(files.Templates, "templates/fragments/sports.html"))
	err := tmpl.Execute(w, h.sports)
	if err != nil {
		log.Println(err.Error())
	}
}

func (h *handler) Courses(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("sport")
	for _, sport := range h.sports {
		if sport.Name == name {
			courses, err := hspscraper.GetAllCourses(sport)
			if err != nil {
				log.Println(err.Error())
				return
			}
			tmpl := template.Must(template.ParseFS(files.Templates, "templates/fragments/courses.html"))
			err = tmpl.Execute(w, courses)
			if err != nil {
				log.Println(err.Error())
				return
			}
			break
		}
	}

}

func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	registration := Registration{
		Sport:        r.FormValue("sport"),
		CourseNumber: r.FormValue("courseNumber"),
		Date:         time.Time{},
		Email:        email,
		Password:     password,
	}
	err := h.registrationHandler.AddRegistration(&registration)
	if err != nil {
		log.Println(err.Error())
		return
	}
	tmpl := template.Must(template.ParseFS(files.Templates, "templates/fragments/registrations.html"))
	err = tmpl.Execute(w, h.registrationHandler.GetRegistrationsByUser(email, password))
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func (h *handler) Registrations(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	tmpl := template.Must(template.ParseFS(files.Templates, "templates/fragments/registrations.html"))
	err := tmpl.Execute(w, h.registrationHandler.GetRegistrationsByUser(email, password))
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func (h *handler) Dev(w http.ResponseWriter, _ *http.Request) {
	tmpl := template.Must(template.ParseFS(files.Templates, "templates/fragments/registrations.html"))
	err := tmpl.Execute(w, h.registrationHandler.GetAllRegistrations())
	if err != nil {
		log.Println(err.Error())
		return
	}
}
