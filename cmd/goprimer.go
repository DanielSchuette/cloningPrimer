package main

import (
	"flag"
)

var (
	seqFile     = flag.String("seq_file", "", "valid file path to a ...")
	enzymeFile  = flag.String("enzyme_file", "", "valid file path to a ...\ndefault is the file at 'github.com/DanielSchuette/enzymes.re'")
	enzymeNameF = flag.String("enzyme_name_forward", "", "name of the enzyme you want to use for the 5' end (must be in the '--enzyme_file')")
	enzymeNameR = flag.String("enzyme_name_reverse", "", "name of the enzyme you want to use for the 3' end (must be in the '--enzyme_file')")
	overhangF   = flag.Int("overhang_forward", 4, "number of random nucleotides added to the forward primer (an integer between 2 - 10)")
	overhangR   = flag.Int("overhang_reverse", 4, "number of random nucleotides added to the reverse primer (an integer between 2 - 10)")
)

func main() {
	// parse command line arguments
	flag.Parse()

	// TODO: implement
}
