package category

import (
	ov "github.com/go-ozzo/ozzo-validation"
	"github.com/nhie-io/api/internal/validation"
)

func (c Category) Validate() error {
	// https://github.com/go-ozzo/ozzo-validation/issues/81
	return validation.Reformat(ov.Validate(string(c),
		ov.Required,
		ov.In(string(Harmless), string(Delicate), string(Offensive)),
	))
}
