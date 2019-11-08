package category

import (
	"database/sql/driver"
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
