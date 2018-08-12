package main

import (
	"fmt"
	"log"

	cloningprimer "github.com/DanielSchuette/cloningPrimer"
)

func main() {
	// find forward primer with EcoRI restriction site and 5 random nucleotides
	input := "ATGCAATGTGAGCTTAGCCTGATCCGTAATCGTAAGT"
	forward, err := cloningprimer.FindForward(input, "GAATTC", 1, 18, 4, false)
	if err != nil {
		log.Fatalf("error finding forward primer: %s\n", err)
	}
	if forward == "" {
		log.Fatalf("no forward primer found\n")
	}
	reverse, err := cloningprimer.FindReverse(input, "GAATTC", 1, 18, 4)
	if err != nil {
		log.Fatalf("error finding reverse primer: %s\n", err)
	}
	/*
		if reverse == "" {
			log.Fatalf("no reverse primer found\n")
		}
	*/
	fmt.Printf("input: %s\nforward primer: %s\nreverse primer: %s\n", input, forward, reverse)
}
