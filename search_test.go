package cloningprimer

import "testing"

type searchInput struct {
	m       map[string]RestrictEnzyme
	pattern string
}

type testCaseSearch struct {
	in   searchInput
	want string
	err  error
}

func TestFilterEnzymeMap(t *testing.T) {
	cases := []testCaseSearch{
		// test matching of pattern 'ba' against 'BamHI'
		{
			in:   searchInput{map[string]RestrictEnzyme{"BamHI": {}}, "ba"},
			want: "BamHI",
			err:  nil,
		},
		// test matching of pattern 'cor' against 'EcoRI'
		{
			in:   searchInput{map[string]RestrictEnzyme{"EcoRI": {}}, "cor"},
			want: "EcoRI",
			err:  nil,
		},
		// test matching of pattern 'ba' against 'EcoRI'
		{
			in:   searchInput{map[string]RestrictEnzyme{"EcoRI": {}}, "ba"},
			want: "", /* do not expect a match for this test */
			err:  nil,
		},
	}

	// loop over test cases
	for _, c := range cases {
		m, err := FilterEnzymeMap(c.in.m, c.in.pattern)

		// test similarity of expected and received value
		// get first key of `FilterEnzymeMap' result `m' and assign it to `got'
		var got string
		for key := range m {
			got = key
			break
		}
		if got != c.want {
			t.Errorf("FilterEnzymeMap(%v, %v) == %v, want %v\n", c.in.m, c.in.pattern, got, c.want)
		}

		// if no error is returned, test if none is expected
		if err == nil && c.err != nil {
			t.Errorf("FilterEnzymeMap(%v, %v) == %v, want %v\n", c.in.m, c.in.pattern, got, c.want)
		}

		// if error is returned, test if an error is expected
		if err != nil {
			// if c.err is nil, print wanted and received error
			// else if an error is wanted and received but error messages are not the same
			// print wanted and received error
			if c.err == nil {
				t.Errorf("FilterEnzymeMap(%v, %v) == %v, want %v\n", c.in.m, c.in.pattern, got, c.want)
			} else if err.Error() != c.err.Error() {
				t.Errorf("FilterEnzymeMap(%v, %v) == %v, want %v\n", c.in.m, c.in.pattern, got, c.want)
			}
		}
	}
}
