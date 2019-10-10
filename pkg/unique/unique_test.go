package unique

import "testing"

func TestStrings(t *testing.T) {
	cases := []struct {
		input, output []string
	}{
		{[]string{"foo", "foo", "bar"}, []string{"foo", "bar"}},
		{[]string{"foo", "bar"}, []string{"foo", "bar"}},
		{[]string{}, []string{}},
	}

	for key, testCase := range cases {
		if strings := Strings(testCase.input); len(strings) != len(testCase.output) {
			t.Fatalf("%d: Unexpected length on the returned slice: %v", key, len(strings))
		} else if strings := Strings(testCase.input); !equals(strings, testCase.output) {
			t.Fatalf("%d: Unexpected slice contents: %v", key, strings)
		}
	}
}

func equals(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
