package cloningprimer

import (
	"errors"
	"testing"
)

// test case struct for functions that compute primers
type testCasePrimer struct {
	in   inputForPrimer
	want string
	err  error
}

// struct that is used in primer test case struct and that holds
// input for functions that compute primers
type inputForPrimer struct {
	seq      string
	restrict string
	seqStart int
	length   int
	random   int
	addCodon bool
}

// test case struct for functions that do computations on codons
type testCaseCodon struct {
	in   hasCodon
	want bool
}

// struct that is used in codon test case struct and that holds
// input for respective functions
type hasCodon struct {
	in    string
	exact bool
}

type testCaseComplement struct {
	in   byte
	want byte
	err  error
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
		// length of primer shorter than `MinimumPrimerLength'
		{
			in: inputForPrimer{
				seq:      "ATGGACTCCAACACTGCTCCGCTGGGCCCCTCCTGCCC",
				restrict: "GGATTC",
				seqStart: 1,
				length:   8,
				random:   4,
				addCodon: false,
			},
			want: "",
			err:  errors.New("invalid input length = 8, must be an integer value >= 10 and smaller than the length of the given sequence (as well as <= the maximum primer length of 30)"),
		},
		// length of primer larger than `MaximumPrimerLength'
		{
			in: inputForPrimer{
				seq:      "ATGGACTCCAACACTGCTCCGCTGGGCCCCTCCTGCCC",
				restrict: "GGATTC",
				seqStart: 1,
				length:   32,
				random:   4,
				addCodon: false,
			},
			want: "",
			err:  errors.New("invalid input length = 32, must be an integer value >= 10 and smaller than the length of the given sequence (as well as <= the maximum primer length of 30)"),
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
		// length of primer shorter than `MinimumPrimerLength'
		{
			in: inputForPrimer{
				seq:      "ATGGACTCCAACACTGCTCCGCTGGGCCCCTCCTGCCC",
				restrict: "GGATTC",
				seqStart: 1,
				length:   8,
				random:   4,
				addCodon: false,
			},
			want: "",
			err:  errors.New("invalid input length = 8, must be an integer value >= 10 and smaller than the length of the given sequence (as well as <= the maximum primer length of 30)"),
		},
		// length of primer larger than `MaximumPrimerLength'
		{
			in: inputForPrimer{
				seq:      "ATGGACTCCAACACTGCTCCGCTGGGCCCCTCCTGCCC",
				restrict: "GGATTC",
				seqStart: 1,
				length:   32,
				random:   4,
				addCodon: false,
			},
			want: "",
			err:  errors.New("invalid input length = 32, must be an integer value >= 10 and smaller than the length of the given sequence (as well as <= the maximum primer length of 30)"),
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

func TestHasStartCodon(t *testing.T) {
	cases := []testCaseCodon{
		// test just a start codon
		{
			in:   hasCodon{"ATG", true},
			want: true,
		},
		// test start codon at beginning of a sequence
		{
			in:   hasCodon{"ATGCCGAGACAGT", true},
			want: true,
		},
		// test start codon at end of a sequence (in this case, `exact' must be false)
		{
			in:   hasCodon{"GAGAGCCCACGCGAGATG", false},
			want: true,
		},
		// test sequence without a start codon
		{
			in:   hasCodon{"GAGAGCCACGAGCAGCG", true},
			want: false,
		},
		// if the length of the input is < 3, false should always be returned
		{
			in:   hasCodon{"AT", true},
			want: false,
		},
	}

	// loop over test cases
	for _, c := range cases {
		got := HasStartCodon(c.in.in, c.in.exact)

		// test similarity of expected and received value
		if got != c.want {
			t.Errorf("HasStartCodon(%v, %v) == %v, want %v\n", c.in.in, c.in.exact, got, c.want)
		}
	}
}

func TestHasStopCodon1(t *testing.T) {
	cases := []testCaseCodon{
		// test just a stop codon
		{
			in:   hasCodon{"TTA", true},
			want: true,
		},
		// test stop codon at beginning of a sequence
		{
			in:   hasCodon{"TTACCGAGACAGT", true},
			want: true,
		},
		// test stop codon at end of a sequence (in this case, `exact' must be false)
		{
			in:   hasCodon{"GAGAGCCCACGCGAGTTA", false},
			want: true,
		},
		// test sequence without a stop codon
		{
			in:   hasCodon{"GAGAGCCACGAGCAGCG", true},
			want: false,
		},
		// if the length of the input is < 3, false should always be returned
		{
			in:   hasCodon{"AT", true},
			want: false,
		},
	}

	// loop over test cases
	for _, c := range cases {
		got := HasStopCodon1(c.in.in, c.in.exact)

		// test similarity of expected and received value
		if got != c.want {
			t.Errorf("HasStopCodon1(%v, %v) == %v, want %v\n", c.in.in, c.in.exact, got, c.want)
		}
	}
}

func TestHasStopCodon2(t *testing.T) {
	cases := []testCaseCodon{
		// test just a stop codon
		{
			in:   hasCodon{"CTA", true},
			want: true,
		},
		// test stop codon at beginning of a sequence
		{
			in:   hasCodon{"CTACCGAGACAGT", true},
			want: true,
		},
		// test stop codon at end of a sequence (in this case, `exact' must be false)
		{
			in:   hasCodon{"GAGAGCCCACGCGAGCTA", false},
			want: true,
		},
		// test sequence without a stop codon
		{
			in:   hasCodon{"GAGAGCCACGAGCAGCG", true},
			want: false,
		},
		// if the length of the input is < 3, false should always be returned
		{
			in:   hasCodon{"AT", true},
			want: false,
		},
	}

	// loop over test cases
	for _, c := range cases {
		got := HasStopCodon2(c.in.in, c.in.exact)

		// test similarity of expected and received value
		if got != c.want {
			t.Errorf("HasStopCodon2(%v, %v) == %v, want %v\n", c.in.in, c.in.exact, got, c.want)
		}
	}
}

func TestHasStopCodon3(t *testing.T) {
	cases := []testCaseCodon{
		// test just a stop codon
		{
			in:   hasCodon{"TCA", true},
			want: true,
		},
		// test stop codon at beginning of a sequence
		{
			in:   hasCodon{"TCACCGAGACAGT", true},
			want: true,
		},
		// test stop codon at end of a sequence (in this case, `exact' must be false)
		{
			in:   hasCodon{"GAGAGCCCACGCGAGTCA", false},
			want: true,
		},
		// test sequence without a stop codon
		{
			in:   hasCodon{"GAGAGCCACGAGCAGCG", true},
			want: false,
		},
		// if the length of the input is < 3, false should always be returned
		{
			in:   hasCodon{"AT", true},
			want: false,
		},
	}

	// loop over test cases
	for _, c := range cases {
		got := HasStopCodon3(c.in.in, c.in.exact)

		// test similarity of expected and received value
		if got != c.want {
			t.Errorf("HasStopCodon3(%v, %v) == %v, want %v\n", c.in.in, c.in.exact, got, c.want)
		}
	}
}

func TestComplement(t *testing.T) {
	cases := []testCaseComplement{
		// test nucleotide 'A'
		{
			in:   'A',
			want: 'T',
			err:  nil,
		},
		// test nucleotide 'T'
		{
			in:   'T',
			want: 'A',
			err:  nil,
		},
		// test nucleotide 'G'
		{
			in:   'G',
			want: 'C',
			err:  nil,
		},
		// test nucleotide 'C'
		{
			in:   'C',
			want: 'G',
			err:  nil,
		},
		// test invalid nucleotide
		{
			in:   'Q',
			want: 0,
			err:  errors.New("invalid input: Q is not a nucleotide"),
		},
	}

	// loop over test cases
	for _, c := range cases {
		got, err := Complement(c.in)

		// test similarity of expected and received value
		if got != c.want {
			t.Errorf("Complement(%v) == %v, want %v\n", c.in, got, c.want)
		}

		// if no error is returned, test if none is expected
		if err == nil && c.err != nil {
			t.Errorf("Complement(%v) == %v, want %v\n", c.in, err, c.err)
		}

		// if error is returned, test if an error is expected
		if err != nil {
			// if c.err is nil, print wanted and received error
			// else if an error is wanted and received but error messages are not the same
			// print wanted and received error
			if c.err == nil {
				t.Errorf("Complement(%v) == %v, want %v\n", c.in, err, c.err)
			} else if err.Error() != c.err.Error() {
				t.Errorf("Complement(%v) == %v, want %v\n", c.in, err, c.err)
			}
		}
	}
}

// isEqualByteSlice tests if two byte slices are equal
func isEqualByteSlice(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
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

func TestValidateSequence(t *testing.T) {
	// TODO: implement unit tests
}
