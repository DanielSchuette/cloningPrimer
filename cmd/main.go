package main

import (
	"fmt"
	"log"

	cloningprimer "github.com/DanielSchuette/cloningPrimer"
)

func main() {
	// find forward primer with EcoRI restriction site and 5 random nucleotides
	res, err := cloningprimer.FindForward("CAATGTGAGCTTAGCCTGATCCGTAATCGTAAGT", "GAATTC", 1, 10, 12)
	if err != nil {
		log.Fatalf("error finding forward primer: %s\n", err)
	}
	if res == "" {
		log.Fatalf("no forward primer returned\n")
	}
	fmt.Printf("result: %s\n", res)

	// find forward primer with wrong sequence
	// TODO: implement unit test for validity checks
	res2, err := cloningprimer.FindForward("tAATGTGACTTAGCCTGATCCGTAATCGTAAGT", "GAATTC", 1, 10, 12)
	if err != nil {
		log.Fatalf("error finding forward primer: %s\n", err)
	}
	if res2 == "" {
		log.Fatalf("no forward primer returned\n")
	}
	fmt.Printf("result: %s\n", res2)
}
