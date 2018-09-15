package cloningprimer

import (
	"errors"
	"testing"
)

type testCaseEnzyme struct {
	in   string
	want map[string]RestrictEnzyme
	err  error
}

type testCaseSequence struct {
	in   string
	want string
	err  error
}

func TestParseEnzymesFromFile(t *testing.T) {
	cases := []testCaseEnzyme{
		// test invalid file extension
		{
			in:   "tests/parse1.seq",
			want: nil,
			err:  errors.New("invalid input: tests/parse1.seq is not a *.re file (see ./doc.go for more information)"),
		},
		// test non-existing file
		{
			in:   "tests/doesnotexist.re",
			want: nil,
			err:  errors.New("error opening file: open tests/doesnotexist.re: no such file or directory"),
		},
		// test correct parsing of enzymes from a file without comments: `parse1.re'
		{
			in: "tests/parse1.re",
			want: map[string]RestrictEnzyme{
				"AclI": {
					Name:            "AclI",
					RecognitionSite: "aslkfhsdf",
					NoPalinCleav:    "sdlfkj",
					ID:              "sldkfj",
					Isoschizomeres:  []string{"sdfklj", "sdlfkj"},
				},
			},
			err: nil,
		},
		// test correct parsing of enzymes from a file with comments: `parse2.re'
		{
			in: "tests/parse2.re",
			want: map[string]RestrictEnzyme{
				"AclI": {
					Name:            "AclI",
					RecognitionSite: "aslkfhsdf",
					NoPalinCleav:    "sdlfkj",
					ID:              "sdfk",
					Isoschizomeres:  []string{"sdlkfj", "sdflkj"},
				},
			},
			err: nil,
		},
		// test correct parsing of enzymes from a file with comments but no column labels: `parse2.re'
		{
			in: "tests/parse3.re",
			want: map[string]RestrictEnzyme{
				"AclI": {
					Name:            "AclI",
					RecognitionSite: "slfkj",
					NoPalinCleav:    "sdlkfj",
					ID:              "sldkfj",
					Isoschizomeres:  []string{"sldkjf"},
				},
			},
			err: nil,
		},
	}

	// loop over test cases
	for _, c := range cases {
		got, err := ParseEnzymesFromFile(c.in)

		// test similarity of expected and received value
		if !isSimilarMap(got, c.want) {
			t.Errorf("ParseEnzymesFromFile(%v) == %v, want %v\n", c.in, got, c.want)
		}

		// if no error is returned, test if none is expected
		if err == nil && c.err != nil {
			t.Errorf("ParseEnzymesFromFile(%v) == %v, want %v\n", c.in, err, c.err)
		}

		// if error is returned, test if an error is expected
		if err != nil {
			// if c.err is nil, print wanted and received errors
			// else if an error is wanted and received but error messages are not the same
			// print wanted and received error
			if c.err == nil {
				t.Errorf("ParseEnzymesFromFile(%v) == %v, want %v\n", c.in, err, c.err)
			} else if err.Error() != c.err.Error() {
				t.Errorf("ParseEnzymesFromFile(%v) == %v, want %v\n", c.in, err, c.err)
			}
		}
	}
}

func TestParseSequenceFromFile(t *testing.T) {
	cases := []testCaseSequence{
		// test invalid file extension
		{
			in:   "tests/parse1.re",
			want: "",
			err:  errors.New("invalid input: tests/parse1.re is not a *.seq file (see ./doc.go for more information)"),
		},
		// test non-existing file
		{
			in:   "tests/doesnotexist.seq",
			want: "",
			err:  errors.New("error opening file: open tests/doesnotexist.seq: no such file or directory"),
		},
		// test correct parsing of sequence from a file without comments: `parse1.seq'
		{
			in:   "tests/parse1.seq",
			want: "ATGGCCGCGT",
			err:  nil,
		},
		// test correct parsing of enzymes from a file with comments: `parse2.seq'
		{
			in:   "tests/parse2.seq",
			want: "ATGGCCGCGT",
			err:  nil,
		},
		// test correct parsing of sequence from a file with additional white spaces and returns: `parse3.seq'
		{
			in:   "tests/parse3.seq",
			want: "ATGGCCGCGTTGACGAGTGAGCATAGGCA",
			err:  nil,
		},
		// test parsing of sequence with non-nucleotide letters from a file: `parse4.seq'
		{
			in:   "tests/parse4.seq",
			want: "ATTATGA",
			err:  errors.New("invalid letter in nucleotide sequence: Q at position 7"),
		},
	}

	// loop over test cases
	for _, c := range cases {
		got, err := ParseSequenceFromFile(c.in)

		// test similarity of expected and received value
		if got != c.want {
			t.Errorf("ParseSequenceFromFile(%v) == %v, want %v\n", c.in, got, c.want)
		}

		// if no error is returned, test if none is expected
		if err == nil && c.err != nil {
			t.Errorf("ParseSequenceFromFile(%v) == %v, want %v\n", c.in, err, c.err)
		}

		// if error is returned, test if an error is expected
		if err != nil {
			// if c.err is nil, print wanted and received errors
			// else if an error is wanted and received but error messages are not the same
			// print wanted and received error
			if c.err == nil {
				t.Errorf("ParseSequenceFromFile(%v) == %v, want %v\n", c.in, err, c.err)
			} else if err.Error() != c.err.Error() {
				t.Errorf("ParseSequenceFromFile(%v) == %v, want %v\n", c.in, err, c.err)
			}
		}
	}
}

// isSimilarMap only tests if all keys in `m1' exist in `m2' and if all mapped `RestrictEnzyme'
// structs have matching values for their fields
// this function is not robust if invalid input is provided and should only be used for testing purposes!
func isSimilarMap(m1, m2 map[string]RestrictEnzyme) bool {
	if (m1 == nil) && (m2 == nil) {
		return true
	}
	for k, v := range m1 {
		// test fields of type string
		if val, ok := m2[k]; !ok {
			return false
		} else if val.Name != v.Name {
			return false
		} else if val.RecognitionSite != v.RecognitionSite {
			return false
		} else if val.NoPalinCleav != v.NoPalinCleav {
			return false
		} else if val.ID != v.ID {
			return false
		} else if len(val.Isoschizomeres) != len(v.Isoschizomeres) {
			return false
		} else {
			for i := 0; i < len(val.Isoschizomeres); i++ {
				if val.Isoschizomeres[i] != v.Isoschizomeres[i] {
					return false
				}
			}
		}
	}
	return true
}
