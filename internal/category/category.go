package category

import (
	"database/sql/driver"
	"fmt"
)

type Category string

const (
	Harmless  Category = "harmless"
	Delicate  Category = "delicate"
	Offensive Category = "offensive"
)

func (c *Category) Scan(src interface{}) error {
	switch src := src.(type) {
	case []byte: // text column
		*c = Category(src)
		return nil
	case string: // varchar column
		*c = Category(src)
		return nil
	default:
		return fmt.Errorf("scan: unable to scan type %T into Category", src)
	}
}

func (c Category) Value() (driver.Value, error) {
	return c.String(), nil
}

func (c Category) String() string {
	return string(c)
}
