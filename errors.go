package env

import (
	"errors"
	"os"
)

var (
	ErrMissingValue  = errors.New("missing value")
	ErrParsingFailed = errors.New("parsing failed")
)

func DefaultErrorHandler(err error) {
	os.Stderr.WriteString(err.Error() + "\n")
	os.Exit(1)
}

type missingError struct {
	Name string
}

func (e *missingError) Error() string {
	return "env: missing value for '" + e.Name + "'"
}

func (e *missingError) Unwrap() error {
	return ErrMissingValue
}

type parsingError struct {
	Name string
	err  error
}

func (e *parsingError) Error() string {
	return "env: failed to parse '" + e.Name + "': " + e.err.Error()
}

func (e *parsingError) Unwrap() error {
	return wrap(ErrParsingFailed, e.Error())
}

type wrappedError struct {
	err error
	msg string
}

func wrap(err error, msg string) error {
	return &wrappedError{err: err, msg: msg}
}

func (e *wrappedError) Error() string {
	return e.msg + ": " + e.err.Error()
}

func (e *wrappedError) Unwrap() error {
	return e.err
}
