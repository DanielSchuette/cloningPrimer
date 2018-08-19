package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"

	cloningprimer "github.com/DanielSchuette/cloningPrimer"
)

var (
	tmpl    *template.Template
	enzymes map[string]cloningprimer.RestrictEnzyme
	err     error
	local   = flag.Bool("local", false, "set this argument to `true' to run the server locally at 127.0.0.1:8080")
)

func init() {
	// parse templates
	tmpl = template.Must(template.ParseGlob("templates/*"))

	// parse `enzymes.re' and create map of restriction enzyme structs
	enzymes, err = cloningprimer.ParseEnzymesFromFile("assets/enzymes.re")
	if err != nil {
		log.Fatalf("error loading enzymes: %v\n", err)
	}
}

func main() {
	// parse command line flags
	flag.Parse()

	// get port
	port := getPort()

	// register handler funcs
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/index/", indexHandler)
	http.HandleFunc("/documentation/", documentationHandler)
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
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Get the Port from the environment for Heroku (only if `--local' is not set)
func getPort() string {
	if *local {
		return ":8080"
	}

	// get port from the environment when running Heroku app
	var port = os.Getenv("PORT")

	// Set a default port if `$PORT' is not set
	if port == "" {
		port = "8080"
		log.Printf("no $PORT environment variable detected, defaulting to %v\n" + port)
	}
	return ":" + port
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// root directory is re-directed to '/index/'
	http.Redirect(w, r, "/index/", http.StatusFound)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "index", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func documentationHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "documentation", nil)
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
	log.Printf("/search/ r.Form['Query']: %v\n", r.Form["Query"])

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
	err := tmpl.ExecuteTemplate(w, "design", enzymes)
	if err != nil {
		log.Fatal(err)
	}
}

func computePrimersHandler(w http.ResponseWriter, r *http.Request) {
	// parse request form and print query information on server site
	r.ParseForm()
	log.Printf("/computePrimers/ r.Form['sequenceQuery']: %v\n", r.Form["sequenceQuery"])

	// if no input was received, return `designcompute' template without data
	// this particular input box returns a slice containing just one string
	if len(r.Form["sequenceQuery"]) > 0 {
		if r.Form["sequenceQuery"][0] == "" {
			err := tmpl.ExecuteTemplate(w, "design", enzymes)
			if err != nil {
				log.Fatal(err)
			}
			return
		}
		err := tmpl.ExecuteTemplate(w, "designcompute", enzymes)
		if err != nil {
			log.Fatal(err)
		}
		return
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
