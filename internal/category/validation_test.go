package category

import "testing"

func TestCategory_Validate(t *testing.T) {
	cases := []struct {
		input      Category
		shouldPass bool
	}{
		{Category("harmless"), true},
		{Category("delicate"), true},
		{Category("offensive"), true},
		{Category("foo"), false},
		{Category("hArmLesS"), false},
	}

	for _, testCase := range cases {
		if err := testCase.input.Validate(); (err == nil) != testCase.shouldPass {
			if testCase.shouldPass {
				t.Fatalf("Valiation should pass. %+v", testCase.input)
			}
			t.Fatalf("Valiation should not pass. %+v", testCase.input)
		}
	}
}
