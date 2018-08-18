package main

import (
	"html/template"
	"log"
	"net/http"

	cloningprimer "github.com/DanielSchuette/cloningPrimer"
)

var tmpl *template.Template
var enzymes map[string]cloningprimer.RestrictEnzyme
var err error

func init() {
	// parse templates
	tmpl = template.Must(template.ParseGlob("templates/*"))

	// parse `enzymes.re' and create map of restriction enzyme structs
	enzymes, err = cloningprimer.ParseEnzymesFromFile("../assets/enzymes.re")
	if err != nil {
		log.Fatalf("error loading enzymes: %v\n", err)
	}
}

func main() {
	// register handler funcs
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/index/", indexHandler)
	http.HandleFunc("/enzymesPage/", enzymesHandler)
	http.HandleFunc("/search/", enzymesSearchHandler)
	http.HandleFunc("/designPage/", designHandler)
	http.HandleFunc("/computePrimers/", computePrimersHandler)
	http.HandleFunc("/links/", linksHandler)
	http.HandleFunc("/license/", licenseHandler)
	http.HandleFunc("/contribute/", contributeHandler)

	// file server for static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// listen and serve locally
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// root directory is re-directed to '/index/'
	http.Redirect(w, r, "/index", http.StatusFound)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "index", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func enzymesHandler(w http.ResponseWriter, r *http.Request) {
	// execute template with map of restriction enzymes as input
	err := tmpl.ExecuteTemplate(w, "enzymes", enzymes)
	if err != nil {
		log.Fatal(err)
	}
}

func enzymesSearchHandler(w http.ResponseWriter, r *http.Request) {
	// parse request form and print query information on server site
	r.ParseForm()
	log.Printf("/search r.Form['Query']: %v\n", r.Form["Query"])

	// execute template with map of restriction enzymes as input
	// if user entered a search term, filter `enzymes' accordingly
	// otherwise, do not filter at all
	if len(r.Form["Query"]) > 0 {
		if r.Form["Query"][0] != "" {
			// filter enzyme map
			e, err := cloningprimer.FilterEnzymeMap(enzymes, r.Form["Query"][0])
			if err != nil {
				log.Fatalf("error filtering enzymes: %v\n", err)
			}

			// if at least one enzyme name matches the query, parse template with that result
			if len(e) > 0 {
				err := tmpl.ExecuteTemplate(w, "enzymessearch", e)
				if err != nil {
					log.Fatal(err)
				}
			}

			// if no enzyme name matches the query, parse template without data
			if len(e) == 0 {
				err := tmpl.ExecuteTemplate(w, "enzymessearch", nil)
				if err != nil {
					log.Fatal(err)
				}
			}
		}

		// if query is empty, return full list of enzymes
		if r.Form["Query"][0] == "" {
			err := tmpl.ExecuteTemplate(w, "enzymes", enzymes)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	if len(r.Form["Query"]) == 0 {
		err := tmpl.ExecuteTemplate(w, "enzymes", enzymes)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func designHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "design", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func computePrimersHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "designcompute", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func linksHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "links", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func licenseHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "license", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func contributeHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "contribute", nil)
	if err != nil {
		log.Fatal(err)
	}
}
