package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	cloningprimer "github.com/DanielSchuette/cloningPrimer"
)

func main() {
	http.HandleFunc("/index/", handler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// parse template from file
	t := template.Must(template.New("index.html").ParseFiles("templates/index.html"))

	// create map of restriction enzyme structs
	enzymes, err := cloningprimer.ParseEnzymesFromFile("../assets/enzymes.re")
	if err != nil {
		fmt.Fprintf(w, "error loading enzymes: %v\n", err)
		log.Fatalf("error loading enzymes: %v\n", err)
	}

	// execute template with map of restriction enzymes as input
	err = t.Execute(w, enzymes)
	if err != nil {
		fmt.Fprintf(w, "error executing template: %v\n", err)
		log.Fatal(err)
	}
}
