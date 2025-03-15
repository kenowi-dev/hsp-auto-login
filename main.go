package main

import (
	"log"
	"net/http"
	"strconv"
	"svelte-go/api"
	"svelte-go/frontend"
)

func main() {
	port := 8080
	mux := http.NewServeMux()

	mux.HandleFunc("/", frontend.StaticIndexHandlerFunc)
	mux.HandleFunc("/api/sports", api.SportsHandlerFunc)
	mux.HandleFunc("/api/courses", api.CoursesHandlerFunc)

	log.Printf("App running on %s...\n", strconv.Itoa(port))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), mux))
}

type Sport struct {
	Name        string `json:"name,omitempty"`
	Href        string `json:"href,omitempty"`
	InFlexiCard bool   `json:"in_flexi_card,omitempty"`
	ExtraInfo   string `json:"extra_info,omitempty"`
}

func mockGetSports() []Sport {
	var sports []Sport
	sports = append(sports, Sport{
		Name:        "Sport",
		Href:        "https://sports.svelte.com/sports/",
		InFlexiCard: true,
		ExtraInfo:   "",
	})

	sports = append(sports, Sport{
		Name:        "Sport2",
		Href:        "https://sports.svelte.com/sports/",
		InFlexiCard: true,
		ExtraInfo:   "",
	})

	sports = append(sports, Sport{
		Name:        "Sport3",
		Href:        "https://sports.svelte.com/sports/",
		InFlexiCard: true,
		ExtraInfo:   "",
	})
	return sports
}
