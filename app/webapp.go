package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("index.html").ParseFiles("app/templates/index.html"))
	err := t.Execute(w, nil)
	if err != nil {
		fmt.Fprintf(w, "error: %v\n", err)
		log.Fatal(err)
	}
}
