package cloningprimer

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

// internally used struct that holds data regarding a certain restriction enzyme
type restrictEnzyme struct {
	name            string /* e.g. EcoRI */
	recognitionSite string /* e.g. AACGTT */
	noPalinCleav    string /* either "no" or "(...)(...)", see *.re specification in ./assets/enzymes.re */
}

// ParseEnyzmesFromFile parses enzyme data (identifiers, recognition sequences, etc.) from
// a *.re file (see the example in ./enzymes.re) and returns a map with enzyme names as
// keys and `restricEnzyme' structs as values
func ParseEnyzmesFromFile(file string) (map[string]string, error) {
	// check validity of input
	// return an error if `file' is not a *.re
	if path.Ext(file) != ".re" {
		return nil, fmt.Errorf("", file)
	}

	// TODO: implement more error checking

	// open file and read its contents
	f, err := os.Open(file)
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(b)

	// TODO: implement

	m := make(map[string]string)
	return m, nil
}

// ParseSequenceFromFile parses a plasmid or DNA sequence from a *.seq file (see the example
// in ./WFS1.seq) and returns the sequence as a string
func ParseSequenceFromFile(file string) string {
	// TODO: implement
	return ""
}
