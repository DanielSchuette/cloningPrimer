package cloningprimer

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

// RestrictEnzyme is an internally used struct that holds data of a certain restriction enzyme
type RestrictEnzyme struct {
	Name            string   /* e.g. EcoRI */
	RecognitionSite string   /* e.g. AACGTT */
	NoPalinCleav    string   /* either "no" or "(...)(...)", see *.re specification in ./assets/enzymes.re */
	ID              string   /* the PDB ID of the enzyme */
	Isoschizomeres  []string /* common isoschizomeres */
}

// ParseEnzymesFromFile parses enzyme data (identifiers, recognition sequences, etc.) from
// a *.re file (see the example in ./assets/enzymes.re) and returns a map with enzyme names as
// keys and `restricEnzyme' structs as values
func ParseEnzymesFromFile(file string) (map[string]RestrictEnzyme, error) {
	// check validity of input
	// return an error if `file' is not a *.re
	if path.Ext(file) != ".re" {
		return nil, fmt.Errorf("invalid input: %v is not a *.re file (see ./doc.go for more information)", file)
	}

	// open file and read its contents
	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer func() {
		err = f.Close()
		if err != nil {
			log.Fatalf("error closing file: %v\n", err)
		}
	}()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("error reading from file: %v", err)
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
	for i, n := 0, len(b); i < n; i++ {
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
					case 3:
						itemContainer.ID = string(dataItem)
					case 4:
						itemContainer.Isoschizomeres = strings.Split(string(dataItem), ",")
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

	fmt.Printf("parsed %d enzymes from '%s'\n", line, file)
	return enzymesMap, nil
}

// ParseSequenceFromFile parses a plasmid or DNA sequence from a *.seq file (see the example
// in ./assets/tp53.seq) and returns the sequence as a string
func ParseSequenceFromFile(file string) (string, error) {
	// check validity of input
	// return an error if `file' is not a *.seq
	if path.Ext(file) != ".seq" {
		return "", fmt.Errorf("invalid input: %v is not a *.seq file (see ./doc.go for more information)", file)
	}

	// open file and read its contents
	f, err := os.Open(file)
	if err != nil {
		return "", fmt.Errorf("error opening file: %v", err)
	}
	defer func() {
		err = f.Close()
		if err != nil {
			log.Fatalf("error closing file: %v", err)
		}
	}()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return "", fmt.Errorf("error reading from file: %v", err)
	}

	// parse data line-wise into a `restrictEnzyme' struct
	var noNucleotides int /* variable to keep track of number of parsed nucleotides *.seq file (for user output) */
	var parse bool        /* variable to keep track of whether the current line should be parsed or not */
	var seq []byte        /* temporary variable to hold the growing nucleotide sequence as it is parsed */

Loop:
	for i, n := 0, len(b); i < n; i++ {
		// decide what to do next
		if i < len(b)-2 {
			// if current char is a new line delimiter, decide how to proceed
			if b[i] == '\n' {
				switch {
				// if next char is not a comment '/' and the next char not '*' => set parse to true and continue parsing the next line
				case (b[i+1] != '/') && (b[i+2] != '*'):
					parse = true
					continue Loop

				// if next char is a comment '/' or '*' => do not parse next line
				case (b[i+1] == '/') || (b[i+2] == '*'):
					parse = false
				}
			}
		}

		// edge case: if the current line is the first line, set parse to false if this line is a comment
		// otherwise, set parse to true and start parsing the sequence
		if i == 0 {
			if (b[i] == '/') && (b[i+1] == '*') {
				parse = false
			} else {
				parse = true
			}
		}

		// if parse is false, continue the loop until the switch triggers again
		if !parse {
			continue Loop
		}

		// if current char is not within a comment line, assume that it is part of a nucleotide sequence
		// if the current char is not a valid nucleotide char, return an error to the caller
		// white spaces and other funky characters are ignored
		if b[i] == 9 || b[i] == 10 || b[i] == 11 || b[i] == 12 || b[i] == 13 || b[i] == 32 {
			continue Loop
		}
		if !IsNucleotide(b[i]) {
			return string(seq), fmt.Errorf("invalid letter in nucleotide sequence: %s at position %d", string(b[i]), noNucleotides)
		}
		seq = append(seq, b[i])
		noNucleotides++
	}

	fmt.Printf("parsed %d nucleotides from '%s'\n", noNucleotides, file)
	return string(seq), nil
}
