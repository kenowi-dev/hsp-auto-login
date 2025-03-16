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
	mux.HandleFunc("/api/coursesDates", api.CoursesDatesHandlerFunc)

	log.Printf("App running on %s...\n", strconv.Itoa(port))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), mux))
}
