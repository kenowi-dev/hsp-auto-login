package api

import (
	"log"
	"net/http"
	"svelte-go/hspscraper"
)

func SportsHandlerFunc(w http.ResponseWriter, _ *http.Request) {
	sports, err := hspscraper.GetAllFlexiCardSports()
	if err != nil {
		log.Fatal(err.Error())
	}
	writeJson(w, sports)
}
