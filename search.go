package cloningprimer

import (
	"fmt"
	"regexp"
	"strings"
)

// FilterEnzymeMap takes a map with keys of type `string' and values of type `RestrictEnzyme' and returns a slice of strings containing enzyme names that match a certain query string `query'
func FilterEnzymeMap(enzymeMap map[string]RestrictEnzyme, query string) (map[string]RestrictEnzyme, error) {
	// copy the input because maps are reference types
	inputCopy := make(map[string]RestrictEnzyme)
	for key, value := range enzymeMap {
		inputCopy[key] = value
	}

	// define a regex pattern and delete all non-matches from copied map
	pattern := query + ".*"
	for key := range inputCopy {
		matched, err := regexp.MatchString(strings.ToLower(pattern), strings.ToLower(key))
		if err != nil {
			return nil, fmt.Errorf("error matching pattern %v: %v", pattern, err)
		}
		if !matched {
			delete(inputCopy, key)
		}
	}
	return inputCopy, nil
}
