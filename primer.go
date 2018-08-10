package cloningprimer

import "fmt"

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
// position and up to (starting position + length). E.g. if `length = 10` and `start = 1`,
// a primer will be returned that binds to nucleotides 1 - 10.
func FindForward(seq, restrict string, start, length, random int) (string, error) {
	// check validity of input
	// return an error if `seq` contains invalid letters (anything except for A,T,C,G)
	// TODO: implement
	// return an error if `random` > 6 or < 3
	// TODO: implement

	// loop over letters in sequence and append the appropriate ones to a slice of bytes
	b := make([]byte, 0)
	for i, l := range []byte(seq) {
		if (i >= (start - 1)) && !(i > (length - 1)) {
			b = append(b, l)
		}
		if i > (length - 1) {
			break
		}
	}
	// append restriction side
	result := restrict + string(b)

	// apppend random nucleotides
	for i := 2; i < (random + 2); i++ {
		switch {
		case i%4 == 0:
		case i%3 == 0:
		case i%4 == 0:
		case i%5 == 0:
		}
	}
	// TODO: implement

	// diagnostic output
	fmt.Printf("forward primer: %s\n", result)
	return result, nil
}

// FindReverse finds a reverse primer of specified length binding at the specified end position
// TODO: documentation
func FindReverse(seq string, end, length int) (string, error) {
	return "", nil
}
