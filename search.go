package cloningprimer

// FilterEnzymeMap takes a map with keys of type `string' and values of type `RestrictEnzyme' and returns a slice of strings containing enzyme names that match a certain query string `query'
func FilterEnzymeMap(enzymeMap map[string]RestrictEnzyme, query string) ([]string, error) {
	var result []string
	for key := range enzymeMap {
		result = append(result, key)
	}
	return result, nil
}
