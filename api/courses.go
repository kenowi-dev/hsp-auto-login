package api

import (
	"net/http"
	"svelte-go/hspscraper"
)

func CoursesHandlerFunc(w http.ResponseWriter, r *http.Request) {
	sportName := r.URL.Query().Get("sport")
	if sportName == "" {
		http.Error(w, "sportName is required", http.StatusBadRequest)
	}

	sport, err := hspscraper.FindSport(sportName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	courses, err := hspscraper.GetAllCourses(sport)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJson(w, courses)
}

func CoursesDatesHandlerFunc(w http.ResponseWriter, r *http.Request) {
	var course hspscraper.Course
	err := parseBody(r, &course)
	if err != nil {
		http.Error(w, "course cannot be parsed", http.StatusBadRequest)
	}

	dates, err := hspscraper.GetCourseDates(&course)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	writeJson(w, dates)
}
