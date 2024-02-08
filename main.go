package main

import (
	"fmt"
	"hackathon/hackathon"
	"log"
	"net/http"
)

func main() {
	// Create a new instance of the hackathon application
	app := hackathon.New()

	// Configure HTTP routes
	http.HandleFunc("/", app.HandleIndex)
	http.HandleFunc("/sort", app.Sort)
	http.HandleFunc("/search", app.HandleSearch)
	http.HandleFunc("/filter", app.HandleFilter)
	http.HandleFunc("/add", app.Add)
	http.HandleFunc("/edit", app.Edit)
	http.HandleFunc("/delete", app.Delete)
	http.HandleFunc("/update", app.Update)
	http.HandleFunc("/get", app.Get)
	http.HandleFunc("/get_all", app.GetAll)
	http.HandleFunc("/get_all_sorted", app.GetAllSorted)
	http.HandleFunc("/get_all_filtered", app.GetAllFiltered)
	http.HandleFunc("/get_all_searched", app.GetAllSearched)

	// Start the HTTP server on port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Println("Server started on port 8080...")
}
