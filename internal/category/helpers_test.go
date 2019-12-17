package category

import "testing"

func TestGetRandom(t *testing.T) {
	categories := []Category{
		Harmless,
		Delicate,
		Offensive,
	}

	cases := []struct {
		input, output []Category
	}{
		{nil, categories},
		{categories, categories},
		{categories[:1], categories[:1]},
		{categories[:2], categories[:2]},
	}

	for key, testCase := range cases {
		category := GetRandom(testCase.input...)

		// category found in expected output
		found := false

		for _, c := range testCase.output {
			if c == category {
				found = true
				break
			}
		}

		if !found {
			t.Fatalf("%d: Unexpected category: %+v", key, category)
		}
	}
}
