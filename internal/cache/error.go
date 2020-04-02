package cache

import "errors"

var ErrKeyNotFound = errors.New("key not found")

type Error struct {
	error
}

func newError(err error) *Error {
	return &Error{err}
}

func (e *Error) Error() string {
	return "cache -> " + e.error.Error()
}

func (e *Error) Unwrap() error {
	return e.error
}

type SerializeError struct {
	error
}

func newSerializeError(err error) *SerializeError {
	return &SerializeError{err}
}

func (e *SerializeError) Error() string {
	return "serialize -> " + e.error.Error()
}

func (e *SerializeError) Unwrap() error {
	return e.error
}
