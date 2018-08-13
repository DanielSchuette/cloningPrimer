package cloningprimer

import (
	"errors"
	"testing"
)

type testCasePrimer struct {
	in   inputForPrimer
	want string
	err  error
}

type inputForPrimer struct {
	seq      string
	restrict string
	seqStart int
	length   int
	random   int
	addCodon bool
}

func TestFindForward(t *testing.T) {
	cases := []testCasePrimer{
		// test invalid input
		// test `seq' with non-nucleotide letters
		{
			in: inputForPrimer{
				seq:      "QVDASTGASD", /* first invalid letter should result in an error */
				restrict: "GAATTC",
				seqStart: 3,
				length:   3,
				random:   4,
				addCodon: true,
			},
			want: "",
			err:  errors.New("invalid input Q at position 1, expected sequence of lower or upper case A,T,C,G"),
		},
		{
			in: inputForPrimer{
				seq:      "ATGCCGVDASTGASD", /* first invalid letter should result in an error */
				restrict: "GAATTC",
				seqStart: 3,
				length:   3,
				random:   4,
				addCodon: true,
			},
			want: "",
			err:  errors.New("invalid input V at position 7, expected sequence of lower or upper case A,T,C,G"),
		},
		// test `seq' that is exactly of length (`length' + `seqStart' - 1)
		{
			in: inputForPrimer{
				seq:      "ATGCCGTCGCATTCTG",
				restrict: "GAATTC",
				seqStart: 1,
				length:   16,
				random:   4,
				addCodon: true,
			},
			want: "AGCTGAATTCATGCCGTCGCATTCTG", /* the expected output is the entire sequence plus overhang */
			err:  nil,
		},
		// test `seq' that is shorter than (`length' + `seqStart' - 1)
		{
			in: inputForPrimer{
				seq:      "ATGCCGTCGCATTGTCCATCT",
				restrict: "GAATTC",
				seqStart: 10,
				length:   16,
				random:   4,
				addCodon: true,
			},
			want: "",
			err:  errors.New("invalid input, the given sequence (21 nucleotides) is not long enough for a primer of length = 16 starting at nucleotide 10 (25 > 21)"),
		},
		// test valid input for `random'
		{
			in: inputForPrimer{
				seq:      "ATGCCGTCGCATTGTCCATCT",
				restrict: "GAATTC",
				seqStart: 10,
				length:   16,
				random:   1,
				addCodon: true,
			},
			want: "",
			err:  errors.New("invalid input random = 1, expected integer value between 2 and 10"),
		},
		{
			in: inputForPrimer{
				seq:      "ATGCCGTCGCATTGTCCATCT",
				restrict: "GAATTC",
				seqStart: 10,
				length:   16,
				random:   22,
				addCodon: true,
			},
			want: "",
			err:  errors.New("invalid input random = 22, expected integer value between 2 and 10"),
		},
		// sequence does not start with an 'ATG' (start codon) and `startCodon' is false
		{
			in: inputForPrimer{
				seq:      "CTGCCGTCGCATTGTCCATCTTACTGACCTGATGTGCCA",
				restrict: "GAATTC",
				seqStart: 10,
				length:   16,
				random:   8,
				addCodon: false,
			},
			want: "",
			err:  errors.New("input sequence does not begin with a start codon ('ATG')\nmake sure to automatically add a start codon by setting `startCodon' to `true'"),
		},
		// addition of 'ATG' start codon
		{
			in: inputForPrimer{
				seq:      "CTGCCGTCGCATTGTCCATCTTACTGACCTGATGTGCCA",
				restrict: "GAATTC",
				seqStart: 8,
				length:   16,
				random:   3,
				addCodon: true,
			},
			want: "GCTGAATTCATGCGCATTGTCCATCTTA", /* start codon should be inserted right after recognition sequence */
			err:  nil,
		},
		// invalid `seqStart'
		{
			in: inputForPrimer{
				seq:      "CTGCCGTCGCATTGTCCATCTTACTGACCTGATGTGCCA",
				restrict: "GAATTC",
				seqStart: 0,
				length:   16,
				random:   3,
				addCodon: true,
			},
			want: "",
			err:  errors.New("invalid input: primer start point must be an integer > 0 (not 0)"),
		},
		// experimentally validated primer
		{
			in: inputForPrimer{
				seq:      "ATGGACTCCAACACTGCTCCGCTGGGCCCCTCCTGCCCACAGCCCCCGCCAGCACCGCAGCCCCAGGCGCGTTCCCGACTCAATGCCAC",
				restrict: "GGATCC",
				seqStart: 1,
				length:   18,
				random:   4,
				addCodon: false,
			},
			want: "AGCTGGATCCATGGACTCCAACACTGCT",
			err:  nil,
		},
	}

	// loop over test cases
	for _, c := range cases {
		got, err := FindForward(c.in.seq, c.in.restrict, c.in.seqStart, c.in.length, c.in.random, c.in.addCodon)

		// test similarity of expected and received value
		if got != c.want {
			t.Errorf("FindForward(%v, %v, %v, %v, %v, %v) == %v, want %v\n", c.in.seq, c.in.restrict, c.in.seqStart, c.in.length, c.in.random, c.in.addCodon, got, c.want)
		}

		// if no error is returned, test if none is expected
		if err == nil && c.err != nil {
			t.Errorf("FindForward(%v, %v, %v, %v, %v, %v) == %v, want %v\n", c.in.seq, c.in.restrict, c.in.seqStart, c.in.length, c.in.random, c.in.addCodon, err, c.err)
		}

		// if error is returned, test if an error is expected
		if err != nil {
			// if c.err is nil, print wanted and received error
			// else if an error is wanted and received but error messages are not the same
			// print wanted and received error
			if c.err == nil {
				t.Errorf("FindForward(%v, %v, %v, %v, %v, %v) == %v, want %v\n", c.in.seq, c.in.restrict, c.in.seqStart, c.in.length, c.in.random, c.in.addCodon, err, c.err)
			} else if err.Error() != c.err.Error() {
				t.Errorf("FindForward(%v, %v, %v, %v, %v, %v) == %v, want %v\n", c.in.seq, c.in.restrict, c.in.seqStart, c.in.length, c.in.random, c.in.addCodon, err, c.err)
			}
		}
	}
}

func TestFindReverse(t *testing.T) {
	cases := []testCasePrimer{
		// test invalid input
		// test `seq' with non-nucleotide letters
		{
			in: inputForPrimer{
				seq:      "QVDASTGASD", /* first invalid letter should result in an error */
				restrict: "GAATTC",
				seqStart: 3,
				length:   3,
				random:   4,
				addCodon: true,
			},
			want: "",
			err:  errors.New("invalid input Q at position 1, expected sequence of lower or upper case A,T,C,G"),
		},
		{
			in: inputForPrimer{
				seq:      "ATGCCGVDASTGASD", /* first invalid letter should result in an error */
				restrict: "GAATTC",
				seqStart: 3,
				length:   3,
				random:   4,
				addCodon: true,
			},
			want: "",
			err:  errors.New("invalid input V at position 7, expected sequence of lower or upper case A,T,C,G"),
		},
		// test `seq' that is exactly of length (`length' + `seqStart' - 1)
		{
			in: inputForPrimer{
				seq:      "ATGCCGTCGCATTCTG",
				restrict: "GAATTC",
				seqStart: 1,
				length:   16,
				random:   4,
				addCodon: true,
			},
			want: "AGCTGAATTCTTACAGAATGCGACGGCAT", /* the expected output is the entire sequence plus overhang */
			err:  nil,
		},
		// test `seq' that is shorter than (`length' + `seqStart' - 1)
		{
			in: inputForPrimer{
				seq:      "ATGCCGTCGCATTGTCCATCT",
				restrict: "GAATTC",
				seqStart: 10,
				length:   16,
				random:   4,
				addCodon: true,
			},
			want: "",
			err:  errors.New("invalid input, the given sequence (21 nucleotides) is not long enough for a primer of length = 16 starting at nucleotide 10 (25 > 21)"),
		},
		// test valid input for `random'
		{
			in: inputForPrimer{
				seq:      "ATGCCGTCGCATTGTCCATCT",
				restrict: "GAATTC",
				seqStart: 10,
				length:   16,
				random:   1,
				addCodon: true,
			},
			want: "",
			err:  errors.New("invalid input random = 1, expected integer value between 2 and 10"),
		},
		{
			in: inputForPrimer{
				seq:      "ATGCCGTCGCATTGTCCATCT",
				restrict: "GAATTC",
				seqStart: 10,
				length:   16,
				random:   22,
				addCodon: true,
			},
			want: "",
			err:  errors.New("invalid input random = 22, expected integer value between 2 and 10"),
		},
		// sequence does not start with an 'ATG' (start codon) and `startCodon' is false
		{
			in: inputForPrimer{
				seq:      "CTGCCGTCGCATTGTCCATCTTACTGACCTGATGTGCCA",
				restrict: "GAATTC",
				seqStart: 10,
				length:   16,
				random:   8,
				addCodon: false,
			},
			want: "",
			err:  errors.New("input sequence does not begin with a stop codon ('TAA', 'TAG', 'TGA')\nmake sure to automatically add a start codon by setting `startCodon' to `true'"),
		},
		// recognition of valid 'TGA' stop codon at position 8 from 3' end
		{
			in: inputForPrimer{
				seq:      "CTGCCGTCGCATTGTCCATCTTACTGACCTGATGTGCCA",
				restrict: "GAATTC",
				seqStart: 8,
				length:   16,
				random:   3,
				addCodon: true,
			},
			want: "GCTGAATTCTCAGGTCAGTAAGATG", /* stop codon should be correctly recognized and inserted */
			err:  nil,
		},
		{
			in: inputForPrimer{
				seq:      "AAAAATTTTTTTCCATCAGGCGCTGATGGCGAAGTTAGCGTAG", /* has 'TAG' stop codon */
				restrict: "GAATTC",
				seqStart: 1,
				length:   20,
				random:   3,
				addCodon: true,
			},
			want: "GCTGAATTCCTACGCTAACTTCGCCATCA",
			err:  nil,
		},
		// insertion of 'TAA' stop codon
		{
			in: inputForPrimer{
				seq:      "AAAAATTTTTTTCCATCAGGCGCTGATGGCGAAGTTAGCG", /* same `seq' as previous example, but w/o stop */
				restrict: "GAATTC",
				seqStart: 1,
				length:   20,
				random:   3,
				addCodon: true,
			},
			want: "GCTGAATTCTTACGCTAACTTCGCCATCAGCG",
			err:  nil,
		},
		// invalid `seqStart'
		{
			in: inputForPrimer{
				seq:      "CTGCCGTCGCATTGTCCATCTTACTGACCTGATGTGCCA",
				restrict: "GAATTC",
				seqStart: 0,
				length:   16,
				random:   3,
				addCodon: true,
			},
			want: "",
			err:  errors.New("invalid input: primer start point must be an integer > 0 (not 0)"),
		},
		// experimentally validated primer
		{
			in: inputForPrimer{
				seq:      "CGTCATCCCCAGCAGCCTGTTCCTGCAGGACGACGAAGATGATGACGAGCTGGCGGGGAAGAGCCCTGAGGACCTGCCACTGCGT",
				restrict: "GAATTC",
				seqStart: 1,
				length:   18,
				random:   4,
				addCodon: true,
			},
			want: "AGCTGAATTCTTAACGCAGTGGCAGGTCCTC",
			err:  nil,
		},
	}

	// loop over test cases
	for _, c := range cases {
		got, err := FindReverse(c.in.seq, c.in.restrict, c.in.seqStart, c.in.length, c.in.random, c.in.addCodon)

		// test similarity of expected and received value
		if got != c.want {
			t.Errorf("FindReverse(%v, %v, %v, %v, %v, %v) == %v, want %v\n", c.in.seq, c.in.restrict, c.in.seqStart, c.in.length, c.in.random, c.in.addCodon, got, c.want)
		}

		// if no error is returned, test if none is expected
		if err == nil && c.err != nil {
			t.Errorf("FindReverse(%v, %v, %v, %v, %v, %v) == %v, want %v\n", c.in.seq, c.in.restrict, c.in.seqStart, c.in.length, c.in.random, c.in.addCodon, err, c.err)
		}

		// if error is returned, test if an error is expected
		if err != nil {
			// if c.err is nil, print wanted and received error
			// else if an error is wanted and received but error messages are not the same
			// print wanted and received error
			if c.err == nil {
				t.Errorf("FindReverse(%v, %v, %v, %v, %v, %v) == %v, want %v\n", c.in.seq, c.in.restrict, c.in.seqStart, c.in.length, c.in.random, c.in.addCodon, err, c.err)
			} else if err.Error() != c.err.Error() {
				t.Errorf("FindReverse(%v, %v, %v, %v, %v, %v) == %v, want %v\n", c.in.seq, c.in.restrict, c.in.seqStart, c.in.length, c.in.random, c.in.addCodon, err, c.err)
			}
		}
	}
}

func TestComplement(t *testing.T) {
	// TODO: implement unit tests
}

func TestReverse(t *testing.T) {
	// TODO: implement unit tests
}

func TestIsNucleotide(t *testing.T) {
	// TODO: implement unit tests
}

func TestAddOverhang(t *testing.T) {
	// TODO: implement unit tests
}

func TestHasStartCodon(t *testing.T) {
	// TODO: implement unit tests
}

func TestHasStopCodon(t *testing.T) {
	// TODO: implement unit tests
}
