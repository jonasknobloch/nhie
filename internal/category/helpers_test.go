package category

import "testing"

func TestGetRandom(t *testing.T) {
	categories := []Category{
		Harmless,
		Delicate,
		Offensive,
	}

	category := GetRandom()

	for _, c := range categories {
		if c == category {
			return
		}
	}

	t.Fatalf("Unexpected category. %+v", category)
}
