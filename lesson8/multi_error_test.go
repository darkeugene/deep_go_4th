package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiError(t *testing.T) {
	var err error
	err = Append(err, errors.New("error 1"))
	err = Append(err, errors.New("error 2"))

	expectedMessage := "2 errors occured:\n\t* error 1\t* error 2\n"
	assert.EqualError(t, err, expectedMessage)
}

func TestMultiError1(t *testing.T) {
	var err error
	err = Append(err, errors.New("error 1"), errors.New("error 2"))
	err = Append(err, errors.New("error 3"))

	expectedMessage := "3 errors occured:\n\t* error 1\t* error 2\t* error 3\n"
	assert.EqualError(t, err, expectedMessage)
}

func TestMultiErrorUnwrap(t *testing.T) {
	var multErr1 error
	err1 := errors.New("error 1")
	err2 := errors.New("error 2")
	err3 := errors.New("error 3")
	multErr1 = Append(multErr1, err1, err2)
	multErr1 = Append(multErr1, err3)

	errs := errors.Unwrap(multErr1)
	_ = errs
}

func TestMultiErrorIS(t *testing.T) {
	var multErr1, multErr2 error
	err1 := errors.New("error 1")
	err2 := errors.New("error 2")
	err3 := errors.New("error 3")
	multErr1 = Append(multErr1, err1, err2)
	multErr1 = Append(multErr1, err3)

	expectedMessage1 := "3 errors occured:\n\t* error 1\t* error 2\t* error 3\n"
	assert.EqualError(t, multErr1, expectedMessage1)

	err4 := errors.New("error 4")
	err5 := errors.New("error 5")
	multErr2 = Append(multErr2, err4, err5)

	expectedMessage2 := "2 errors occured:\n\t* error 4\t* error 5\n"
	assert.EqualError(t, multErr2, expectedMessage2)

	assert.False(t, errors.Is(multErr1, errors.New("error 1")))
	assert.False(t, errors.Is(multErr2, errors.New("error 2")))

	assert.True(t, errors.Is(multErr1, err1))
	assert.True(t, errors.Is(multErr1, err2))
	assert.True(t, errors.Is(multErr1, err3))

	assert.False(t, errors.Is(multErr1, multErr2))

	multErr1 = Append(multErr1, multErr2)

	expectedMessageComp := "4 errors occured:\n\t* error 1\t* error 2\t* error 3\t* 2 errors occured:\n\t* error 4\t* error 5\n\n"
	assert.EqualError(t, multErr1, expectedMessageComp)

	assert.True(t, errors.Is(multErr1, multErr2))
}

func TestMultiErrorAS(t *testing.T) {
	var multErr1 error
	err1 := errors.New("error 1")
	err2 := &customError{}

	multErr1 = Append(multErr1, err1, err2)

	expectedMessage1 := "2 errors occured:\n\t* error 1\t* custom error\n"
	assert.EqualError(t, multErr1, expectedMessage1)

	var typeMul *MultiError
	var typeCustom *customError

	assert.True(t, errors.As(multErr1, &typeMul))
	assert.True(t, errors.As(multErr1, &typeCustom))
}

type customError struct{}

func (e *customError) Error() string {
	return "custom error"
}
