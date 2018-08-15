package cloningprimer

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

// RestrictEnzyme is an internally used struct that holds data of a certain restriction enzyme
type RestrictEnzyme struct {
	Name            string /* e.g. EcoRI */
	RecognitionSite string /* e.g. AACGTT */
	NoPalinCleav    string /* either "no" or "(...)(...)", see *.re specification in ./assets/enzymes.re */
}

// ParseEnyzmesFromFile parses enzyme data (identifiers, recognition sequences, etc.) from
// a *.re file (see the example in ./enzymes.re) and returns a map with enzyme names as
// keys and `restricEnzyme' structs as values
func ParseEnyzmesFromFile(file string) (map[string]RestrictEnzyme, error) {
	// check validity of input
	// return an error if `file' is not a *.re
	if path.Ext(file) != ".re" {
		return nil, fmt.Errorf("invalid input: %v is not a *.re file (see ./doc.go for more information)", file)
	}

	// TODO: implement more error checking

	// open file and read its contents
	f, err := os.Open(file)
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	// create map that will ultimately be returned
	enzymesMap := make(map[string]RestrictEnzyme)

	// parse data line-wise into a `restrictEnzyme' struct
	var column int                    /* variable to keep track of current column in *.re file */
	var line int                      /* variable to keep track of current line in *.re file (for user output) */
	var parse bool                    /* variable to keep track of whether the current line should be parsed or not */
	var openQuote bool                /* variable to keep track of whether a certain "'" is currently open or not */
	var dataItem []byte               /* temporary variable to keep track of current data item */
	var itemContainer *RestrictEnzyme /* temporary variable to hold current data item before adding it to map */

Loop:
	for i := 0; i < len(b); i++ {
		// decide what to do next
		if i < len(b)-2 {
			// if current char is a new line delimiter, decide how to proceed
			if b[i] == '\n' {
				switch {
				// next char is a valid data item delimiter and next line not a comment
				// set parse to true, reset column counter, and increment line counter
				// also, create a new `RestrictEnzyme' struct `itemContainer' after adding
				// the current `itemContainer' to the `enzymesMap'
				case (b[i+1] == '\'') && (b[i+2] != '*'):
					line++
					column = 0
					parse = true
					if itemContainer != nil {
						if _, ok := enzymesMap[itemContainer.Name]; !ok {
							enzymesMap[itemContainer.Name] = *itemContainer
						}
					}
					itemContainer = new(RestrictEnzyme)

				// next char is not a valid data item delimiter or a comment => do not parse next line
				case (b[i+1] != '\'') || (b[i+2] == '*'):
					parse = false
				}
			}

			// if parse is false, continue the loop until the switch triggers again
			if !parse {
				continue Loop
			}

			// if current char is a valid item delimiter '\'', perform the appropriate action
			if b[i] == '\'' {
				// if this '\'' is delimiting the end of a data item, add the current `dataItem'
				// to the `itemContainer' field that corresponds to the current `column'
				// then, increment column count and continue loop after resetting the temporary
				// data item variable `dataItem' and set `openQuote' to false
				if openQuote {
					switch column {
					case 0:
						itemContainer.Name = string(dataItem)
					case 1:
						itemContainer.RecognitionSite = string(dataItem)
					case 2:
						itemContainer.NoPalinCleav = string(dataItem)
					}
					column++
					dataItem = make([]byte, 0)
					openQuote = false
					continue Loop
				}

				// if this '\'' is delimiting the start of a data item set `openQuote' to true
				// and continue loop to not add the opening '\'' to the respective string
				openQuote = true
				continue Loop
			}
		}

		// if parser is inbetween quotes, append byte to current `dataItem'
		if openQuote {
			dataItem = append(dataItem, b[i])
		}
	}

	fmt.Printf("parsed %d enzymes\n", line)
	return enzymesMap, nil
}

// ParseSequenceFromFile parses a plasmid or DNA sequence from a *.seq file (see the example
// in ./tp53.seq) and returns the sequence as a string
func ParseSequenceFromFile(file string) string {
	// check validity of input
	// return an error if `file' is not a *.seq
	if path.Ext(file) != ".seq" {
		return nil, fmt.Errorf("invalid input: %v is not a *.seq file (see ./doc.go for more information)", file)
	}

	// TODO: implement more error checking

	// open file and read its contents
	f, err := os.Open(file)
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: implement commenting in *.seq files

	return string(b)
}
