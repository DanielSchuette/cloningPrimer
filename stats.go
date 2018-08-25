package cloningprimer

import (
	"errors"
	"fmt"
)

// CalculateGC takes a `primer' as an input and returns the GC nucleotide content as a floating point number between 0.0 and 1.0
func CalculateGC(primer string) (float64, error) {
	// check validity of input
	seq, err := ValidateSequence([]byte(primer))
	if err != nil {
		return 0.0, fmt.Errorf("error while calculating GC content: %v", err)
	}

	// iterate over sequence and count the occurances of 'G' and 'C'
	// the sequence should be all upper cases at this point
	var counter float64
	for _, e := range seq {
		if (e == 'G') || (e == 'C') {
			counter++
		}
	}

	// return the percentage of 'G' and 'C' in the target sequence and no error
	return counter / float64(len(seq)), nil
}

// CalculateTm takes a `primer' (5' -> 3') as an input and returns the melting temperature
// (or Tm) as a floating point number; it uses the formula: Tm = 2°C * (A + T) + 4°C * (C + G)
// `complementary' indicates how many nucleotides (from the 3' end of the `primer') should be considered
// if `complementary' is `0', this argument is ignored and the entire `primer' is used for calculations
func CalculateTm(primer string, complementary int) (float64, error) {
	// check validity of input
	if complementary < 0 {
		return 0.0, fmt.Errorf("invalid input: `complementary' should be >= 0, not %d", complementary)
	}
	seq, err := ValidateSequence([]byte(primer))
	if err != nil {
		return 0.0, fmt.Errorf("error while calculating Tm: %v", err)
	}
	if len(primer) > 15 || ((complementary != 0) && (complementary > 15)) {
		return 0.0, errors.New("this method should only be used for sequences with less than 15 nucleotides")
	}

	// iterate over sequence and keep track of two sums: ('A' + 'T') and ('G' + 'C')
	// all chars in the sequence should be upper case by now
	var gcSum float64
	var atSum float64
	var counter int
	for i, n := (len(seq) - 1), 0; i >= n; i-- {
		e := seq[i]
		switch e {
		case 'G', 'C':
			gcSum++
		case 'A', 'T':
			atSum++
		}
		counter++
		if complementary != 0 {
			if counter == complementary {
				break
			}
		}
	}
	return (2 * atSum) + (4 * gcSum), nil
}
