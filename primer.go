package cloningprimer

import (
	"fmt"
	"strings"
)

/*
* TODO:
*  find forward primer: func FindForward()
*  find complement/reverse complement of a sequence: func Reverse(), func Complement()
*  find reverse primer: func FindReverse()
*  add restriction site + random bp's: AddOverhang
*  calculate stats, e.g. GC content (read up on that): func CalcStats()
* TODO:
*  pull restriction sites from the internet: GetRSites()
*  pull plasmid information from file and test for restriction sites: func CheckPlasmid ()
*  offer access to NCBI blast API (if available?): Blast()
 */

// FindForward finds a forward primer of specified length binding at the specified starting
// position and up to (starting position + length). E.g. if `length' = 10 and `start' = 1,
// a primer will be returned that binds to nucleotides 1 - 10.
func FindForward(seq, restrict string, start, length, random int) (string, error) {
	// check validity of input
	// return an error if `seq' contains invalid letters (anything except for A,T,C,G)
	for i := 0; i < len(seq); i++ {
		if !IsNucleotide(seq[i]) {
			return "", fmt.Errorf("invalid input %s at position %d, expected sequence of lower or upper case A,T,C,G", string(seq[i]), i+1)
		}
	}

	// return an error if `random' > 6 or < 3
	if (random < 3) && (random > 6) {
		return "", fmt.Errorf("invalid input random = %v, expected integer value between 3 and 6", random)
	}

	// loop over letters in sequence and append the appropriate ones to a slice of bytes
	b := make([]byte, 0)
	for i, l := range []byte(seq) {
		if (i >= (start - 1)) && !(i > (length - 1)) {
			l = byte(strings.ToUpper(string(l))) /* make current letter a string, upper case, and byte again */
			b = append(b, l)
		}
		if i > (length - 1) {
			break
		}
	}
	// append restriction side
	result := restrict + string(b) /* concatenate the newly assembled string and the user input `restrict' */

	// append pseudo random nucleotides
	for i := 2; i < (random + 2); i++ {
		switch {
		case i%4 == 0:
			result = "G" + result
		case i%3 == 0:
			result = "C" + result
		case i%2 == 0:
			result = "T" + result
		default:
			result = "A" + result
		}
	}
	return result, nil
}

// FindReverse finds a reverse primer of specified length binding at the specified end position
// TODO: documentation
func FindReverse(seq string, end, length int) (string, error) {
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
