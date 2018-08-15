package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	cloningprimer "github.com/DanielSchuette/cloningPrimer"
)

var (
	seqFile     = flag.String("seq_file", "", "valid file path to a ")
	enzymeFile  = flag.String("enzyme_file", "", "valid file path to a *.re file with correctly formatted restriction enzyme information\ndefault is the file at 'github.com/DanielSchuette/assets/enzymes.re'")
	enzymeNameF = flag.String("enzyme_name_forward", "", "name of the enzyme you want to use for the 5' end (must be in the '--enzyme_file')")
	enzymeNameR = flag.String("enzyme_name_reverse", "", "name of the enzyme you want to use for the 3' end (must be in the '--enzyme_file')")
	overhangF   = flag.Int("overhang_forward", 4, "number of random nucleotides added to the forward primer (an integer between 2 - 10)")
	overhangR   = flag.Int("overhang_reverse", 4, "number of random nucleotides added to the reverse primer (an integer between 2 - 10)")
)

func main() {
	// parse command line arguments
	flag.Parse()

	// check if `seqFile' and `enzymeFile' arguments are provided
	if (*seqFile == "") || (*enzymeFile == "") {
		fmt.Println("arguments `--seq_file' and `--enzyme_file' are required (see `--help' for more information)")
		os.Exit(1)
	}

	// load *.re file
	enzymes, err := cloningprimer.ParseEnyzmesFromFile(*enzymeFile)
	if err != nil {
		log.Fatalf("error while loading *.re file: %v\n", err)
	}
	fmt.Println(enzymes)

	// TODO: implement
}
