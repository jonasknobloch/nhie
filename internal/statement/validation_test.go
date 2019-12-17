package statement

import (
	"github.com/google/uuid"
	"github.com/neverhaveiever-io/api/internal/category"
	"testing"
	"time"
)

func TestStatement_Validate(t *testing.T) {
	cases := []struct {
		input      string
		shouldPass bool
	}{
		{"Never have I ever fucked a coconut.", true},
		{"Never have I ever fucked a coconut", false},
		{"never have I ever fucked a coconut.", false},
		{"Never have I ever.", false},
	}

	for _, testCase := range cases {
		statement := Statement{
			ID:        uuid.UUID{},
			Statement: testCase.input,
			Category:  category.Offensive,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: nil,
		}

		if err := statement.Validate(); (err == nil) != testCase.shouldPass {
			if testCase.shouldPass {
				t.Fatalf("Valiation should pass. %+v", statement)
			}
			t.Fatalf("Valiation should not pass. %+v", statement)
		}
	}
}
