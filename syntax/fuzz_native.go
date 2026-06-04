package syntax

import "testing"

// FuzzSyntax is the native Go fuzzing entry point.
func FuzzSyntax(f *testing.F) {
	f.Add("hello")
	f.Add("(a|b)+")
	f.Fuzz(func(t *testing.T, s string) {
		tree, err := Parse(s, ParseOptions{})
		if err != nil {
			return
		}
		if _, err := Write(tree); err != nil {
			t.Fatal(err)
		}
	})
}
