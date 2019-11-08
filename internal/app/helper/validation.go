package helper

import (
	"encoding/json"
)

type ValidationError struct {
	Name   string `json:"name"`
	Reason string `json:"reason"`
}

func JSONtoValidationErrors(data []byte) ([]ValidationError, error) {
	var objectMap map[string]string

	if err := json.Unmarshal(data, &objectMap); err != nil {
		return nil, err
	}

	var errors []ValidationError

	for key, value := range objectMap {
		errors = append(errors, ValidationError{
			Name:   key,
			Reason: value,
		})
	}

	return errors, nil
}

func ExtractValidationErrors(err error) ([]ValidationError, error) {
	// marshal err into JSON
	data, err := json.Marshal(err)

	if err != nil {
		return nil, err
	}

	// JSON to validation errors
	errors, err := JSONtoValidationErrors(data)

	if err != nil {
		return nil, err
	}

	return errors, nil
}
