package statement

import (
	ov "github.com/go-ozzo/ozzo-validation"
	"github.com/neverhaveiever-io/api/internal/validation"
	"regexp"
)

func (s Statement) Validate() error {
	return validation.Reformat(ov.ValidateStruct(&s,
		ov.Field(
			&s.Statement,
			ov.Required,
			ov.Match(regexp.MustCompile(`Never\shave\sI\sever\s.+\.$`)),
		),
		ov.Field(&s.Category),
	))
}
