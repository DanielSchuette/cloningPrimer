// Package cloningprimer is a software tool that facilitates the design of primer pairs for gene cloning.
// The project lives at https://github.com/DanielSchuette/cloningPrimer and has a web interface and CLI
// in addition to its Go API.
//
// Code Example:
//
/*
	package main

	import (
		"fmt"
		"log"

		cloningprimer "github.com/DanielSchuette/cloningPrimer"
	)

	func main() {
		// define an input string (must be a valid nucleotide sequence)
		input := "ATGCAAAAACGGGCGATTTATCCGGGTACTTTCGATCCCATTACCAATGGTCATATCGATATCGTGACGCGCGCCACGCAGATGTTCGATCACGTTATTCTGGCGATTGCCGCCAGCCCCAGTAAAAAACCGATGTTTACCCTGGAAGAGCGTGTGGCACTGGCACAGCAGGCAACCGCGCATCTGGGGAACGTGGAAGTGGTCGGGTTTAGTGATTTAATGGCGAACTTCGCCCGTAATCAACACGCTACGGTGCTGATTCGTGGCCTGCGTGCGGTGGCAGATTTTGAATATGAAATGCAGCTGGCGCATATGAATCGCCACTTAATGCCGGAACTGGAAAGTGTGTTTCTGATGCCGTCGAAAGAGTGGTCGTTTATCTCTTCATCGTTGGTGAAAGAGGTGGCGCGCCATCAGGGCGATGTCACCCATTTCCTGCCGGAGAATGTCCATCAGGCGCTGATGGCGAAGTTAGCGTAG"

		// find forward primer with EcoRI restriction site and 5 random nucleotides as an overhang
		forward, err := cloningprimer.FindForward(input, "GAATTC", 1, 18, 4, false)
		if err != nil {
			log.Fatalf("error finding forward primer: %s\n", err)
		}
		if forward == "" {
			log.Fatalf("no forward primer found\n")
		}

		// find reverse primer with BamHI restriction site and 3 random nucleotides as an overhang
		reverse, err := cloningprimer.FindReverse(input, "GGATCC", 1, 20, 3, true)
		if err != nil {
			log.Fatalf("error finding reverse primer: %s\n", err)
		}
		if reverse == "" {
			log.Fatalf("no reverse primer found\n")
		}

		// print results
		fmt.Printf("input: %s\nforward primer: %s\nreverse primer: %s\n", input, forward, reverse)
	}
*/
package cloningprimer
