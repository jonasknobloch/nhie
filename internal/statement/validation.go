package statement

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"regexp"
)

func (s Statement) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(
			&s.Statement,
			validation.Required,
			validation.Match(regexp.MustCompile(`Never\shave\sI\sever\s.+\.$`)),
		),
		validation.Field(&s.Category),
	)
}
