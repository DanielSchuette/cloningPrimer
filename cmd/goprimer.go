package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	cloningprimer "github.com/DanielSchuette/cloningPrimer"
	"github.com/fatih/color"
)

var (
	seqFile     = flag.String("seq_file", "../app/assets/tp53.seq", "valid file path to a *.seq file with correctly formatted DNA sequence information\ndefault is the file at 'github.com/DanielSchuette/app/assets/tp53.seq'")
	enzymeFile  = flag.String("enzyme_file", "../app/assets/enzymes.re", "valid file path to a *.re file with correctly formatted restriction enzyme information\ndefault is the file at 'github.com/DanielSchuette/app/assets/enzymes.re'")
	enzymeNameF = flag.String("enzyme_name_forward", "BamHI", "name of the enzyme you want to use for the 5' end (must be in the '--enzyme_file')")
	enzymeNameR = flag.String("enzyme_name_reverse", "EcoRI", "name of the enzyme you want to use for the 3' end (must be in the '--enzyme_file')")
	startPos    = flag.Int("5prime_start", 1, "5' position of the first complementary nucleotide in the provided sequence that the forward primer should bind to\nsee './doc' for more information on how to customize primer calculations")
	stopPos     = flag.Int("3prime_start", 1, "3' position of the first complementary nucleotide in the provided sequence that the reverse primer should bind to\nsee './doc' for more information on how to customize primer calculations")
	overhangF   = flag.Int("overhang_forward", 4, "number of random nucleotides added to the forward primer (an integer between 2 - 10)")
	overhangR   = flag.Int("overhang_reverse", 4, "number of random nucleotides added to the reverse primer (an integer between 2 - 10)")
	lengthF     = flag.Int("length_forward", 18, "length of the complementary part of the forward primer")
	lengthR     = flag.Int("length_reverse", 18, "length of the complementary part of the reverse primer")
	startCodon  = flag.Bool("start_codon", true, "set this flag to 'false' if the input sequence does not have a start codon (an ATG will be added automatically)")
	stopCodon   = flag.Bool("stop_codon", true, "set this flag to 'false' if the input sequence does not have a stop cdon (then, a TAA will be added automatically)")
	verbose     = flag.Bool("verbose", false, "enable verbose output (defaults to false)")
)

func main() {
	// parse command line arguments
	flag.Parse()

	// check if `seqFile' and `enzymeFile' arguments are provided
	if (*seqFile == "") || (*enzymeFile == "") {
		color.Set(color.FgRed) /* make output colorful */
		fmt.Println("arguments `--seq_file' and `--enzyme_file' are required (see `--help' for more information)")
		color.Unset() /* unset colorful output */
		os.Exit(1)
	}

	// load *.re file
	color.Set(color.FgGreen) /* make output colorful */
	enzymes, err := cloningprimer.ParseEnzymesFromFile(*enzymeFile)
	color.Unset() /* unset colorful output */
	color.Unset()
	if err != nil {
		color.Set(color.FgRed) /* make output colorful */
		log.Fatalf("error while loading *.re file: %v\n", err)
		color.Unset() /* unset colorful output */
	}
	if *verbose {
		color.Set(color.FgBlue) /* make output colorful */
		fmt.Println(enzymes)
		color.Unset() /* unset colorful output */
	}

	// load *.seq file
	color.Set(color.FgGreen) /* make output colorful */
	seq, err := cloningprimer.ParseSequenceFromFile(*seqFile)
	color.Unset() /* unset colorful output */
	if err != nil {
		color.Set(color.FgRed) /* make output colorful */
		log.Fatalf("error while loading *.seq file: %v\n", err)
		color.Unset() /* unset colorful output */
	}
	if *verbose {
		color.Set(color.FgBlue) /* make output colorful */
		fmt.Println(seq)
		color.Unset() /* unset colorful output */
	}

	// get forward and reverse primer recognition sequences from the `enzymes' map using regular expression matching
	// report an error if no or more then one enzyme was matched; forward primer:
	var enzymeF string /* variable to hold the 5' enzyme */
	var enzymeR string /* variable to hold the 3' enzyme */
	enzymeFMap, err := cloningprimer.FilterEnzymeMap(enzymes, *enzymeNameF)
	if err != nil {
		color.Set(color.FgRed) /* make output colorful */
		log.Fatalf("error filtering enzyme map: %v\n", err)
		color.Unset() /* unset colorful output */
	}
	if len(enzymeFMap) < 1 {
		color.Set(color.FgRed) /* make output colorful */
		log.Fatalf("invalid input: cannot find %v in '%s'\n", *enzymeNameF, *enzymeFile)
		color.Unset() /* unset colorful output */
	}
	if len(enzymeFMap) > 1 {
		color.Set(color.FgRed) /* make output colorful */
		log.Fatalf("invalid input: %v matches multiple enzyme names in '%s':\n%v\n", *enzymeNameF, *enzymeFile, enzymeFMap)
		color.Unset() /* unset colorful output */
	}
	for k, v := range enzymeFMap {
		color.Set(color.FgYellow) /* make output colorful */
		fmt.Printf("using %v as the 5' restriction enzyme (recognition sequence: %v)\n", k, v.RecognitionSite)
		color.Unset() /* unset colorful output */
		enzymeF = v.RecognitionSite
	}

	// reverse primer:
	enzymeRMap, err := cloningprimer.FilterEnzymeMap(enzymes, *enzymeNameR)
	if err != nil {
		color.Set(color.FgRed) /* make output colorful */
		log.Fatalf("error filtering enzyme map: %v\n", err)
		color.Unset() /* unset colorful output */
	}
	if len(enzymeRMap) < 1 {
		color.Set(color.FgRed) /* make output colorful */
		log.Fatalf("invalid input: cannot find %v in '%s'\n", *enzymeNameR, *enzymeFile)
		color.Unset() /* unset colorful output */
	}
	if len(enzymeRMap) > 1 {
		color.Set(color.FgRed) /* make output colorful */
		log.Fatalf("invalid input: %v matches multiple enzyme names in '%s':\n%v\n", *enzymeNameR, *enzymeFile, enzymeRMap)
		color.Unset() /* unset colorful output */
	}
	for k, v := range enzymeRMap {
		color.Set(color.FgYellow) /* make output colorful */
		fmt.Printf("using %v as the 3' restriction enzyme (recognition sequence: %v)\n", k, v.RecognitionSite)
		color.Unset() /* unset colorful output */
		enzymeR = v.RecognitionSite
	}

	// calculate primers based upon `seq', `enzymeF', and `enzymeR'
	primerF, err := cloningprimer.FindForward(seq, enzymeF, *startPos, *lengthF, *overhangF, *startCodon)
	if err != nil {
		color.Set(color.FgRed) /* make output colorful */
		log.Fatalf("error while computing forward primer: %v\n", err)
		color.Unset() /* unset colorful output */
	}
	primerR, err := cloningprimer.FindReverse(seq, enzymeR, *stopPos, *lengthR, *overhangR, *stopCodon)
	if err != nil {
		color.Set(color.FgRed) /* make output colorful */
		log.Fatalf("error while computing reverse primer: %v\n", err)
		color.Unset() /* unset colorful output */
	}

	// print input parameters and result of calculations
	color.Set(color.FgYellow, color.Bold) /* make output colorful */
	fmt.Println("computing primers...")
	color.Unset() /* unset colorful output */
	fmt.Printf("a forward primer was computed starting at position %d (from the 5' end of the sequence)\n%d random nucleotides were added before the enzyme recognition sequence (%s)\nthe length of the complementary part of the primer is %d\na start codon was added automatically: %v\n", *startPos, *overhangF, enzymeF, *lengthF, !*startCodon)
	color.Set(color.FgGreen, color.Bold) /* make output colorful */
	fmt.Printf("result: %s\n", primerF)
	color.Unset() /* unset colorful ouput */
	fmt.Printf("a reverse primer was computed starting at position %d (from the 3' end of the sequence)\n%d random nucleotides were added before the enzyme recognition sequence (%s)\nthe length of the complementary part of the primer is %d\na start codon was added automatically: %v\n", *stopPos, *overhangR, enzymeR, *lengthR, !*stopCodon)
	color.Set(color.FgGreen, color.Bold) /* make output colorful */
	fmt.Printf("result: %s\n", primerR)
	color.Unset() /* unset colorful ouput */

	// TODO: calculate statistics and output them to the user
}
