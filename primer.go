package cloningprimer

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

/*
* TODO:
*  UNIT TESTS!
*  find complement/reverse complement of a sequence: func Reverse(), func Complement()
*  find reverse primer: func FindReverse()
*  add restriction site + random bp's: AddOverhang
*  calculate stats, e.g. GC content (read up on that): func CalcStats()
*  check proper start and stop codon (ATC and ... stops)
*  check for frame shifts if primer does not start in the beginning
*  check sequence for restriction sites, start codons, stop codon, ...
* TODO:
*  pull restriction sites from the internet: GetRSites()
*  pull plasmid information from file and test for restriction sites: func CheckPlasmid ()
*  offer access to NCBI blast API (if available?): Blast()
*  check spaces in sequence and just silently delete them
 */

// FindForward finds a forward primer of specified length binding at the specified starting
// position and up to (starting position + length). E.g. if `length' = 10 and `start' = 1,
// a primer will be returned that binds to nucleotides 1 - 10.
func FindForward(seq, restrict string, seqStart, length, random int, startCodon bool) (string, error) {
	// check validity of input
	// return an error if `seq' contains invalid letters (anything except for A,T,C,G)
	for i := 0; i < len(seq); i++ {
		if !IsNucleotide(seq[i]) {
			return "", fmt.Errorf("invalid input %s at position %d, expected sequence of lower or upper case A,T,C,G", string(seq[i]), i+1)
		}
	}

	// return an error if `random' > 6 or < 3
	if (random < 3) || (random > 6) {
		return "", fmt.Errorf("invalid input random = %v, expected integer value between 3 and 6", random)
	}

	// following https://www.neb.com/protocols/1/01/01/primer-design-e6901, a `length' < 16 and > `seq'
	// returns an error
	if (length < 16) || (length > len(seq)) {
		return "", fmt.Errorf("invalid input length = %d, must be an integer value larger than 15 and smaller than the length of the given sequence", length)
	}

	// loop over letters in sequence and append the appropriate ones to a slice of bytes
	b := make([]byte, 0)
	for i, l := range []byte(seq) {
		if (i >= (seqStart - 1)) && !(i > (length - 1)) {
			l := []byte(strings.ToUpper(string(l))) /* make current letter a string, upper case, and byte again */
			b = append(b, l...)
		}
		if i > (length - 1) {
			break
		}
	}
	if !HasStartCodon(string(b)) {
		return "", errors.New("input sequence does not begin with a start codon ('ATG')")
		// TODO: automatically add start codon (`startCodon' bool)
	}
	result := restrict + string(b)             /* concatenate the newly assembled string and the user input `restrict' */
	result = AddOverhang(result, random, true) /* add `random' number of nucleotides to front of `result' */
	return result, nil
}

// FindReverse finds a reverse primer of specified length binding at the specified end position
// TODO: documentation
func FindReverse(seq, restrict string, start, length, random int) (string, error) {
	seqRev := Reverse(seq)
	complement := make([]byte, 0)
	for i := 0; i < len(seq); i++ {
		c, err := Complement([]byte(seqRev)[i])
		if err != nil {
			log.Fatalf("cannot compute complement of %v: %v\n", []byte(seqRev)[i], err)
		}
		complement = append(complement, c)
	}
	// fmt.Printf("sequence: %v\nreversed: %v\nreverse complement: %v\n", seq, seqRev, string(complement))
	// TODO: select appropriate letters from seqRev and return them after prepending restriction site and overhang
	return "", nil
}

// IsNucleotide returns a boolean if input rune is a valid nucleotide letter (i.e. one of A/a/T/t/G/g/C/c)
func IsNucleotide(letter byte) bool {
	// if input `letter' is a valid nucleotide, return true
	switch letter {
	case 'A', 'a', 'T', 't', 'G', 'g', 'C', 'c':
		return true
	default:
		return false
	}
}

// Reverse finds the reverse of a nucleotide sequence; it requires prior checking of possible sources of errors (for example, it does not check if the input sequence contains invalid nucleotide letters); thus, `Reverse' should be called in the context of a valid `seq' input argument
func Reverse(seq string) string {
	seqRev := make([]byte, 0)
	for i := len(seq) - 1; i >= 0; i-- {
		seqRev = append(seqRev, []byte(seq)[i])
	}
	return string(seqRev)
}

// Complement finds the complement of a nucleotide sequence (i.e. Watson-Crick base pairs)
func Complement(nucleotide byte) (byte, error) {
	if !IsNucleotide(nucleotide) {
		return 0, fmt.Errorf("invalid input: %v is not a nucleotide", nucleotide)
	}
	switch nucleotide {
	case 'A', 'a':
		return 'T', nil
	case 'T', 't':
		return 'A', nil
	case 'C', 'c':
		return 'G', nil
	case 'G', 'g':
		return 'C', nil
	}
	return 0, fmt.Errorf("could not find a complementary nucleotide for input: %v", nucleotide)
}

// AddOverhang appends pseudo-random nucleotides as an overhang to the front (`front' = True) or back (`front' = False) of the input nucleotide sequence `seq' (overhang is of length `len')
func AddOverhang(seq string, len int, front bool) string {
	overhang := ""
	for i := 2; i < (len + 2); i++ {
		switch {
		case i%4 == 0:
			overhang = "G" + overhang
		case i%3 == 0:
			overhang = "C" + overhang
		case i%2 == 0:
			overhang = "T" + overhang
		default:
			overhang = "A" + overhang
		}
	}
	if front {
		return overhang + seq
	}
	return seq + overhang
}

// HasStartCodon returns true if the first 3 characters of a given input sequence `seq' are 'ATG'
func HasStartCodon(seq string) bool {
	if []byte(seq)[0] != 'A' {
		return false
	}
	if []byte(seq)[1] != 'T' {
		return false
	}
	if []byte(seq)[2] != 'G' {
		return false
	}
	return true
}

// HasStopCodon returns true if the first 3 characters of a given input sequence `seq' are TODO: complete
func HasStopCodon(seq string) bool {
	return false
}
