package main

import (
	"errors"
	"fmt"
)

func main() {}

func ErrWrap() error {
	err1 := errors.New("Oops Something happend.")
	err2 := fmt.Errorf("An error was encountered - %w", err1)
	return err2
}
func ErrUnwrap() error {
	err2 := ErrUnwrap()
	err1 := errors.Unwrap(err2)
	return err1
}

// Another way of wrapping an error with an error is to create a customized
// error struct like this
type ConnectionError struct {
	Host string
	Port int
	Err  error
}

func (err *ConnectionError) Error() string {
	return fmt.Sprintf("Error connecting to %s at port %d", err.Host, err.Port)
}
