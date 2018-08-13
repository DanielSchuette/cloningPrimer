package cloningprimer

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

const (
	// Codon is a constant of length 3
	Codon = 3
)

// FindForward finds a forward primer with a `length' number of complementary nucleotides, binding to the specified starting position (`seqStart'), counting from the 5' end) and up to (`seqStart' + `length' - 1); e.g. if `length' = 10 and `start' = 1, a primer will be returned that binds to nucleotides 1 - 10; the boolean `startCodon' indicates if an 'ATG' should be added and is only evaluated if no 'ATG' is found in the input `seq' (if that is the case, 'ATG' adds three nucleotides to the total length of the primer); `random' indicates how many random nucleotides should be added as an overhang; `restrict' is a string giving the recognition sequence of a restriction enzyme
func FindForward(seq, restrict string, seqStart, length, random int, startCodon bool) (string, error) {
	// check validity of input
	// return an error if `seqStart' < 1
	if seqStart < 1 {
		return "", fmt.Errorf("invalid input: primer start point must be an integer > 0 (not %d)", seqStart)
	}
	// return an error if `seq' contains invalid letters (anything except for A,T,C,G)
	for i := 0; i < len(seq); i++ {
		if !IsNucleotide(seq[i]) {
			return "", fmt.Errorf("invalid input %s at position %d, expected sequence of lower or upper case A,T,C,G", string(seq[i]), i+1)
		}
	}

	// return an error if `random' > 10 or < 2
	if (random < 2) || (random > 10) {
		return "", fmt.Errorf("invalid input random = %v, expected integer value between 2 and 10", random)
	}

	// a `length' < 16, > 30 and > `seq' returns an error
	if (length < 16) || (length > 30) || (length > len(seq)) {
		return "", fmt.Errorf("invalid input length = %d, must be an integer value larger than 15 and smaller than the length of the given sequence", length)
	}

	// if (`seqStart' + `length' -1) > length of `seq' an error is returned
	if (seqStart + length - 1) > len(seq) { /* subtract 1 because the nucleotide at `seqStart' is part of the sequence */
		return "", fmt.Errorf("invalid input, the given sequence (%d nucleotides) is not long enough for a primer of length = %d starting at nucleotide %d (%d > %d)", len(seq), length, seqStart, seqStart+length-1, len(seq))
	}

	// loop over letters in sequence and append the appropriate ones to a slice of bytes
	var b []byte
	for i, l := range []byte(seq) {
		if (i >= (seqStart - 1)) && !(i >= (seqStart + length - 1)) {
			l := []byte(strings.ToUpper(string(l))) /* make current letter a string, upper case, and byte again */
			b = append(b, l...)
		}
		if i >= (seqStart + length - 1) {
			break
		}
	}

	// if the selected part of `seq' does not have a start codon, check how to proceed
	if !HasStartCodon(string(b)) {
		switch startCodon {
		case true:
			b = append([]byte("ATG"), b...)
		case false:
			return "", errors.New("input sequence does not begin with a start codon ('ATG')\nmake sure to automatically add a start codon by setting `startCodon' to `true'")
		}
	}
	result := restrict + string(b)             /* concatenate the newly assembled string and the user input `restrict' */
	result = AddOverhang(result, random, true) /* add `random' number of nucleotides to front of `result' */
	return result, nil
}

// FindReverse finds a reverse primer with a `length' number of complementary nucleotides, binding to the specified start position measured from the 3' end of `seq' up to nucleotide (`seqStart' + `length' -1); `random' indicates the number of random nucleotides to be added to the primer; `restrict' indicates the restriction enzyme recognition site; the boolean `stopCodon' indicates if a stop codon should be added to the primer (only evaluated if the last 3 nucleotides of the sequence underlying the primer do not make up a valid stop codon - in that case, the stop codon adds three nucleotides to the total length of the primer)
func FindReverse(seq, restrict string, seqStart, length, random int, stopCodon bool) (string, error) {
	// check validity of input
	// TODO: implement

	// compute the reverse of the input sequence and the complementary sequence of the reversed sequence `seqRev'
	seqRev := Reverse(seq)
	var complement []byte
	for i := 0; i < len(seq); i++ {
		c, err := Complement([]byte(seqRev)[i])
		if err != nil {
			log.Fatalf("cannot compute complement of %v: %v\n", []byte(seqRev)[i], err)
		}
		complement = append(complement, c)
	}

	// loop over letters in `complement' sequence and append the appropriate ones to a slice of bytes
	var b []byte
	for i, l := range []byte(complement) {
		if (i >= (seqStart - 1)) && !(i >= (seqStart + length - 1)) {
			l := []byte(strings.ToUpper(string(l))) /* make current letter a string, upper case, and byte again */
			b = append(b, l...)
		}
		if i >= (seqStart + length - 1) {
			break
		}
	}

	// if the selected part of `seq' does not have a start codon, check how to proceed
	if !HasStopCodon1(string(b)) && !HasStopCodon2(string(b)) && !HasStopCodon3(string(b)) {
		switch stopCodon {
		case true:
			b = append([]byte("TTA"), b...)
		case false:
			return "", errors.New("input sequence does not begin with a start codon ('TAA', 'TAG', 'TGA')\nmake sure to automatically add a start codon by setting `startCodon' to `true'")
		}
	}
	result := restrict + string(b)             /* concatenate the newly assembled string and the user input `restrict' */
	result = AddOverhang(result, random, true) /* add `random' number of nucleotides to front of `result' */
	return result, nil
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
	var seqRev []byte
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

// HasStopCodon1 (see also ...2, ...3) returns true if the first 3 characters of a given input sequence `seq' are reversed complements of the stop codon TAA
func HasStopCodon1(seq string) bool {
	// for TAA check TTA
	if []byte(seq)[0] != 'T' {
		return false
	}
	if []byte(seq)[1] != 'T' {
		return false
	}
	if []byte(seq)[2] != 'A' {
		return false
	}
	return true
}

// HasStopCodon2 tests reverse complement of TAG (see HasStopCodon1)
func HasStopCodon2(seq string) bool {
	// for TAG test CTA
	if []byte(seq)[0] != 'C' {
		return false
	}
	if []byte(seq)[1] != 'T' {
		return false
	}
	if []byte(seq)[2] != 'A' {
		return false
	}
	return true
}

// HasStopCodon3 tests reverse complement of TGA (see HasStopCodon1)
func HasStopCodon3(seq string) bool {
	// for TGA test TCA
	if []byte(seq)[0] != 'T' {
		return false
	}
	if []byte(seq)[1] != 'C' {
		return false
	}
	if []byte(seq)[2] != 'A' {
		return false
	}
	return true
}
