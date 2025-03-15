package api

/*
func timeSlotHandler(w http.ResponseWriter, r *http.Request) {
	body, err := getBody(r)
	if err != nil {
		log.Println(err.Error())
		return
	}

	sport := body.Get("sport")
	courseNumber := body.Get("courseNumber")
	couse, err := hspscraper.FindCourse(sport, courseNumber)
	if err != nil {
		log.Println(err.Error())
		return
	}
	dates, err := hspscraper.GetCourseDates(couse)
	if err != nil {
		log.Println(err.Error())
		return
	}

	component := templates.Dates(dates)
	err = component.Render(r.Context(), w)
}


*/
