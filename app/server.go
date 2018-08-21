package main

import (
	"errors"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	cloningprimer "github.com/DanielSchuette/cloningPrimer"
)

var (
	tmpl    *template.Template
	enzymes map[string]cloningprimer.RestrictEnzyme
	err     error
	local   = flag.Bool("local", false, "set this argument to `true' to run the server locally at 127.0.0.1:8080")
)

// struct designForm is used by the server to hold data that was parsed from the
// designpage.html and computeprimers.html pages
type designForm struct {
	Sequence             string                                  /* the nucleotide sequence from the user input */
	ForwardEnzyme        string                                  /* the 5' restriction enzyme from the user input */
	ReverseEnzyme        string                                  /* the 3' restriction enzyme from the user input */
	ForwardComplementary string                                  /* length of 5' primer overlap with target sequence */
	ReverseComplementary string                                  /* length of 3' primer overlap with target sequence */
	ForwardOverhang      string                                  /* number of 5' random nucleotides from the user input */
	ReverseOverhang      string                                  /* number of 3' random nucleotides from the user input */
	Start                string                                  /* 'yes' or 'no', indicating presence of start codon */
	Stop                 string                                  /* 'yes' or 'no', indicating presence of stop codon */
	Enzymes              map[string]cloningprimer.RestrictEnzyme /* holds restriction enzyme information */
	ForwardPrimer        string                                  /* holds the computed forward primer */
	ReversePrimer        string                                  /* holds the computed reverse primer */
}

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
	printDesignFormData(r)
	d, err := parseDesignFormData(r)
	if err != nil {
		log.Printf("error while parsing form for primer computation: %v\n", err)
		err = tmpl.ExecuteTemplate(w, "design", enzymes)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	d.Enzymes = enzymes

	// if no input was received, return `designcompute' template without data
	// this particular input box returns a slice containing just one string
	if d.Sequence == "" || d.ForwardEnzyme == "" || d.ReverseEnzyme == "" {
		err = tmpl.ExecuteTemplate(w, "design", enzymes)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	// if input was received, validate input sequence
	d.Sequence, err = validateSequence([]byte(d.Sequence))
	if err != nil {
		log.Printf("error validating user input sequence: %v\n", err)

		// return `designpage' template to user and return from handler
		err = tmpl.ExecuteTemplate(w, "designcompute", d)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	// compute forward primer and append it to `d' struct
	restrictF := enzymes[d.ForwardEnzyme].RecognitionSite /* get recognition sequence of primer from `enzymes' map */
	overhangF, err := strconv.Atoi(d.ForwardOverhang)     /* get number of random nucleotides */
	if err != nil {
		log.Fatal(err)
	}
	var startBool bool
	switch d.Start {
	// if input sequence has a start codon, set `startCodon' to false (no start codon is going to be added)
	case "yes":
		startBool = false

	// if input sequence has no start codon, set `startCodon' to true (a start codon is going to be added automatically)
	case "no":
		startBool = true
	}
	compF, err := strconv.Atoi(d.ForwardComplementary) /* get length of 5' complementary sequence */
	if err != nil {
		log.Fatal(err)
	}
	d.ForwardPrimer, err = cloningprimer.FindForward(d.Sequence, restrictF, 1, compF, overhangF, startBool)
	if err != nil {
		d.ForwardPrimer = fmt.Sprintf("an error occured: %v", err)
		log.Printf("error calculating forward primer: %v\n", err)
	}

	// compute reverse primer and append it to `d' struct
	restrictR := enzymes[d.ReverseEnzyme].RecognitionSite /* get recognition sequence of primer from `enzymes' map */
	overhangR, err := strconv.Atoi(d.ReverseOverhang)     /* get number of random nucleotides */
	if err != nil {
		log.Fatal(err)
	}
	var stopBool bool
	switch d.Stop {
	// if input sequence has a stop codon, set `stopCodon' to false (no stop codon is going to be added)
	case "yes":
		stopBool = false

	// if input sequence has no stop codon, set `stopCodon' to true (a stop codon is going to be added automatically)
	case "no":
		stopBool = true
	}
	compR, err := strconv.Atoi(d.ReverseComplementary) /* get length of 5' complementary sequence */
	if err != nil {
		log.Fatal(err)
	}
	d.ReversePrimer, err = cloningprimer.FindReverse(d.Sequence, restrictR, 1, compR, overhangR, stopBool)
	if err != nil {
		d.ReversePrimer = fmt.Sprintf("an error occured: %v", err)
		log.Printf("error calculating reverse primer: %v\n", err)
	}

	// execute template with data
	err = tmpl.ExecuteTemplate(w, "designcompute", d)
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

func printDesignFormData(r *http.Request) {
	log.Printf("/computePrimers/ r.Form['sequenceQuery']: %v\n", r.Form["sequenceQuery"])
	log.Printf("/computePrimers/ r.Form['forwardEnzyme']: %v\n", r.Form["forwardEnzyme"])
	log.Printf("/computePrimers/ r.Form['reverseEnzyme']: %v\n", r.Form["reverseEnzyme"])
	log.Printf("/computePrimers/ r.Form['forwardComplementary']: %v\n", r.Form["forwardComplementary"])
	log.Printf("/computePrimers/ r.Form['reverseComplementary']: %v\n", r.Form["reverseComplementary"])
	log.Printf("/computePrimers/ r.Form['forwardOverhang']: %v\n", r.Form["forwardOverhang"])
	log.Printf("/computePrimers/ r.Form['reverseOverhang']: %v\n", r.Form["reverseOverhang"])
	log.Printf("/computePrimers/ r.Form['startRadio']: %v\n", r.Form["startRadio"])
	log.Printf("/computePrimers/ r.Form['stopRadio']: %v\n", r.Form["stopRadio"])
}

func parseDesignFormData(r *http.Request) (designForm, error) {
	// check validity of form field input to avoid `IndexOutOfRange' error when parsing field values
	if (len(r.Form["sequenceQuery"]) == 0) || (len(r.Form["forwardEnzyme"]) == 0) || (len(r.Form["reverseEnzyme"]) == 0) || (len(r.Form["forwardComplementary"]) == 0) || (len(r.Form["reverseComplementary"]) == 0) || (len(r.Form["forwardOverhang"]) == 0) || (len(r.Form["reverseOverhang"]) == 0) || (len(r.Form["startRadio"]) == 0) || (len(r.Form["stopRadio"]) == 0) {
		// if any of the form fields is empty, return an empty `designForm' struct and an error
		return designForm{}, errors.New("zero-length form element")
	}

	// populate `designForm' fields and return `d' to caller
	d := designForm{
		Sequence:             r.Form["sequenceQuery"][0],
		ForwardEnzyme:        r.Form["forwardEnzyme"][0],
		ReverseEnzyme:        r.Form["reverseEnzyme"][0],
		ForwardComplementary: r.Form["forwardComplementary"][0],
		ReverseComplementary: r.Form["reverseComplementary"][0],
		ForwardOverhang:      r.Form["forwardOverhang"][0],
		ReverseOverhang:      r.Form["reverseOverhang"][0],
		Start:                r.Form["startRadio"][0],
		Stop:                 r.Form["stopRadio"][0],
	}
	return d, nil
}

func validateSequence(seq []byte) (string, error) {
	// iterate over sequence and append bytes to `s' while ignoring ' ', '\n', '\r', and so forth
	// return the sequence and an error if a byte is not a valid nucleotide
	var s []byte
	for _, b := range seq {
		if b == 9 || b == 10 || b == 11 || b == 12 || b == 13 || b == 32 {
			continue
		}

		// if the current letter is not a valid nucleotide, return the sequence up to
		// this point and an error
		if !cloningprimer.IsNucleotide(b) {
			s = append(s, b)
			s = append(s, []byte(" ... this character is not a valid nucleotide (must be one of A,T,C,G)")...)
			return string(s), fmt.Errorf("invalid char in nucleotide sequence: %v", string(b))
		}
		s = append(s, []byte(strings.ToUpper(string(b)))...)
	}
	return string(s), nil
}
