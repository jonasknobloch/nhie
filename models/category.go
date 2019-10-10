package models

import (
	"database/sql/driver"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Category string

const (
	Harmless  Category = "harmless"
	Delicate  Category = "delicate"
	Offensive Category = "offensive"
)

func (c *Category) Scan(value interface{}) error {
	*c = Category(value.([]byte))
	return nil
}

func (c Category) Value() (driver.Value, error) {
	return string(c), nil
}

func (c Category) Validate() error {
	// https://github.com/go-ozzo/ozzo-validation/issues/81
	return validation.Validate(string(c),
		validation.Required,
		validation.In(string(Harmless), string(Delicate), string(Offensive)),
	)
}
