package cloningprimer

import (
	"errors"
	"testing"
)

type testCaseGC struct {
	in   string
	want float64
	err  error
}

type testCaseTm struct {
	in   tmInput
	want float64
	err  error
}

type tmInput struct {
	primer        string
	complementary int
}

func TestCalculateGC(t *testing.T) {
	cases := []testCaseGC{
		// test primer with GC content of 50%
		{
			in:   "GGCCTTAA",
			want: 0.5,
			err:  nil,
		},
		// test primer with GC content of 0%
		{
			in:   "TTAA",
			want: 0.0,
			err:  nil,
		},
		// test primer with GC content of 100%
		{
			in:   "GGCC",
			want: 1.0,
			err:  nil,
		},
		// test invalid input: empty `primer' argument
		{
			in:   "",
			want: 0.0,
			err:  errors.New("input sequence `primer' cannot be empty"),
		},
		// test invalid input: non-nucleotide letter
		{
			in:   "QTAG",
			want: 0.0,
			err:  errors.New("error while calculating GC content: invalid char in nucleotide sequence: Q"),
		},
	}

	// loop over test cases
	for _, c := range cases {
		got, err := CalculateGC(c.in)

		// test similarity of expected and received value
		if got != c.want {
			t.Errorf("CalculateGC(%v) == %v, want %v\n", c.in, got, c.want)
		}

		// if no error is returned, test if none is expected
		if err == nil && c.err != nil {
			t.Errorf("CalculateGC(%v) == %v, want %v\n", c.in, err, c.err)
		}

		// if error is returned, test if an error is expected
		if err != nil {
			// if c.err is nil, print wanted and received error
			// else if an error is wanted and received but error messages are not the same
			// print wanted and received error
			if c.err == nil {
				t.Errorf("CalculateGC(%v) == %v, want %v\n", c.in, err, c.err)
			} else if err.Error() != c.err.Error() {
				t.Errorf("CalculateGC(%v) == %v, want %v\n", c.in, err, c.err)
			}
		}
	}
}

func TestCalculateTm(t *testing.T) {
	cases := []testCaseTm{
		// test invalid input for argument `complementary'
		{
			in:   tmInput{"GGCCTTAA", -3},
			want: 0.0,
			err:  errors.New("invalid input: `complementary' should be >= 0, not -3"),
		},
		// test valid input
		{
			in:   tmInput{"GGCCTTAA", 0},
			want: 24.0,
			err:  nil,
		},
		// test input with more than 15 nucleotides (should throw an error)
		{
			in:   tmInput{"AGAGCGAGCGATTGATAGCACCGTGAC", 0},
			want: 0.0,
			err:  errors.New("this method should only be used for sequences with less than 15 nucleotides"),
		},
		// test invalid input: empty `primer' argument
		{
			in:   tmInput{"", 0},
			want: 0.0,
			err:  errors.New("input sequence `primer' cannot be empty"),
		},
		// test invalid input: non-nucleotide letter in `primer'
		{
			in:   tmInput{"AGAGACGCGAQ", 0},
			want: 0.0,
			err:  errors.New("error while calculating Tm: invalid char in nucleotide sequence: Q"),
		},
	}

	// loop over test cases
	for _, c := range cases {
		got, err := CalculateTm(c.in.primer, c.in.complementary)

		// test similarity of expected and received value
		if got != c.want {
			t.Errorf("CalculateTm(%v, %v) == %v, want %v\n", c.in.primer, c.in.complementary, got, c.want)
		}

		// if no error is returned, test if none is expected
		if err == nil && c.err != nil {
			t.Errorf("CalculateTm(%v, %v) == %v, want %v\n", c.in.primer, c.in.complementary, got, c.want)
		}

		// if error is returned, test if an error is expected
		if err != nil {
			// if c.err is nil, print wanted and received error
			// else if an error is wanted and received but error messages are not the same
			// print wanted and received error
			if c.err == nil {
				t.Errorf("CalculateTm(%v, %v) == %v, want %v\n", c.in.primer, c.in.complementary, err, c.err)
			} else if err.Error() != c.err.Error() {
				t.Errorf("CalculateTm(%v, %v) == %v, want %v\n", c.in.primer, c.in.complementary, err, c.err)
			}
		}
	}
}
