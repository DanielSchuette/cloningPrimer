package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	cloningprimer "github.com/DanielSchuette/cloningPrimer"
)

var tmpl *template.Template

func init() {
	// parse templates
	tmpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	// register handler funcs
	http.HandleFunc("/index/", indexHandler)
	http.HandleFunc("/enzymesPage/", enzymesHandler)
	http.HandleFunc("/designPage/", designHandler)
	http.HandleFunc("/links/", linksHandler)
	http.HandleFunc("/license/", licenseHandler)
	http.HandleFunc("/contribute/", contributeHandler)

	// file server for static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// listen and serve locally
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("under construction"))
}

func enzymesHandler(w http.ResponseWriter, r *http.Request) {
	// create map of restriction enzyme structs
	enzymes, err := cloningprimer.ParseEnzymesFromFile("../assets/enzymes.re")
	if err != nil {
		fmt.Fprintf(w, "error loading enzymes: %v\n", err)
		log.Fatalf("error loading enzymes: %v\n", err)
	}

	// execute template with map of restriction enzymes as input
	err = tmpl.ExecuteTemplate(w, "index", enzymes)
	if err != nil {
		fmt.Fprintf(w, "error executing template: %v\n", err)
		log.Fatal(err)
	}
}

func designHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("under construction"))
}

func linksHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("under construction"))
}

func licenseHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("under construction"))
}

func contributeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("under construction"))
}
