package category

import validation "github.com/go-ozzo/ozzo-validation"

func (c Category) Validate() error {
	// https://github.com/go-ozzo/ozzo-validation/issues/81
	return validation.Validate(string(c),
		validation.Required,
		validation.In(string(Harmless), string(Delicate), string(Offensive)),
	)
}
