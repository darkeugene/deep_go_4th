package main

import (
	"fmt"
	"strings"
)

type MultiError struct {
	errs []error
}

func (e *MultiError) Error() string {
	if len(e.errs) == 0 {
		return ""
	}

	var b strings.Builder

	_, _ = fmt.Fprintf(&b, "%d errors occured:\n", len(e.errs))

	for _, err := range e.errs {
		_, _ = fmt.Fprintf(&b, "\t* %s", err.Error())
	}

	b.WriteString("\n")

	return b.String()
}

func (e *MultiError) Unwrap() []error {
	return e.errs
}

func Append(err error, errs ...error) *MultiError {
	mult, ok := err.(*MultiError)
	if !ok {
		mult = &MultiError{}
	}

	mult.errs = append(mult.errs, errs...)

	return mult
}
