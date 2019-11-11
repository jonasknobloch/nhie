package validation

import (
	"encoding/json"
	ov "github.com/go-ozzo/ozzo-validation"
)

type errors struct {
	errs ov.Errors
}

func (es errors) Error() string {
	return es.errs.Error()
}

func (es errors) MarshalJSON() ([]byte, error) {
	var errs []interface{}
	for key, err := range es.errs {
		if ms, ok := err.(json.Marshaler); ok {
			errs = append(errs, struct {
				Name   string         `json:"name"`
				Reason json.Marshaler `json:"reason"`
			}{
				Name:   key,
				Reason: ms,
			})
		} else {
			errs = append(errs, struct {
				Name   string `json:"name"`
				Reason string `json:"reason"`
			}{
				Name:   key,
				Reason: err.Error(),
			})
		}
	}
	return json.Marshal(errs)
}

func Reformat(err error) error {
	if errs, ok := err.(ov.Errors); ok {
		return errors{errs}
	}
	return err
}
