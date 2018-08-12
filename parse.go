package cloningprimer

type restrictEnzyme struct {
	// TODO: implement
}

// ParseEnyzmesFromFile parses enzyme data (identifiers, recognition sequences, etc.) from
// a *.re file (see the example in ./enzymes.re) and returns a map with enzyme names as
// keys and `restricEnzyme' structs as values
func ParseEnyzmesFromFile(file string) map[string]string {
	// f, err := os.OpenFile(file)
	// r := io.Reader(f)
	// fmt.Println(r)
	// TODO: implement
	m := make(map[string]string)
	return m
}

// ParseSequenceFromFile parses a plasmid or DNA sequence from a *.seq file (see the example
// in ./WFS1.seq) and returns the sequence as a string
func ParseSequenceFromFile(file string) string {
	// TODO: implement
	return ""
}
