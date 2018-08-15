// Package cloningprimer ...
//
//
// Resources:
//
// https://www.neb.com/protocols/1/01/01/primer-design-e6901
//
// https://www.protocols.io/view/Primer-Design-for-Restriction-Enzyme-Cloning-E6901-imsvmd?guidelines
//
// https://www.addgene.org/mol-bio-reference/restriction-enzymes/
//
// https://www.addgene.org/protocols/pcr-cloning/
//
// https://www.embl.de/pepcore/pepcore_services/cloning/pcr_strategy/primer_design/amplification/index.html
//
// https://www.embl.de/pepcore/pepcore_services/cloning/pcr_strategy/primer_design/
//
//
// TODO: Methylation can block cleavage by some restriction enzymes. In E. coli, Dam methylase affects the sequence GATC and Dcm methylase affects the sequence CCAGG or CCTGG. If these sequences are present, you will need to use a dam-, dcm- E. coli strain to grow your plasmid.
//
// TODO: check 3' end of primer (3 or more G or C bases at this position - may stabilize nonspecific annealing; 3' thymidine - more prone to mispriming)
//
// TODO: primer length (18 - 30 nucleotides)
//
// TODO: self-complementarity (primer dimers)
//
// TODO: GC content (40 - 60%)
//
// TODO: annealing temperature between overlap of target and primer (delta 2 - 4°C and > 60°C, according to: Tm = 2°C * (A + T) + 4°C * (C + G))
//
// TODO: N-terminal ATG in-frame with ORF
//
// TODO: C-terminal TAA (preferred because less prone to read-through than TAG and TGA; 2 or 3 stop codons are possible) in-frame with ORF
//
//
// More TODO's:
//
// - add more unit tests
//
// - Web interface
//
// - CLI
//
// - API documentation
//
// - calculate stats (see above): func CalcStats()
//
// - check sequence and recognition sites for restriction sites, start codons, stop codon, ...
//
// - pull plasmid information from file and test for restriction sites: func CheckPlasmid ()
//
// - offer access to NCBI blast API (if available?): Blast()
//
// - check spaces in sequence and silently delete them
//
//
// Code Example:
//
/*
	// find forward primer with EcoRI restriction site and 5 random nucleotides as an overhang
	input := "ATGCAAAAACGGGCGATTTATCCGGGTACTTTCGATCCCATTACCAATGGTCATATCGATATCGTGACGCGCGCCACGCAGATGTTCGATCACGTTATTCTGGCGATTGCCGCCAGCCCCAGTAAAAAACCGATGTTTACCCTGGAAGAGCGTGTGGCACTGGCACAGCAGGCAACCGCGCATCTGGGGAACGTGGAAGTGGTCGGGTTTAGTGATTTAATGGCGAACTTCGCCCGTAATCAACACGCTACGGTGCTGATTCGTGGCCTGCGTGCGGTGGCAGATTTTGAATATGAAATGCAGCTGGCGCATATGAATCGCCACTTAATGCCGGAACTGGAAAGTGTGTTTCTGATGCCGTCGAAAGAGTGGTCGTTTATCTCTTCATCGTTGGTGAAAGAGGTGGCGCGCCATCAGGGCGATGTCACCCATTTCCTGCCGGAGAATGTCCATCAGGCGCTGATGGCGAAGTTAGCGTAG"
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
*/
package cloningprimer
