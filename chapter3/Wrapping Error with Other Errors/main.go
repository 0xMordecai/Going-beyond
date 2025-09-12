package main

import (
	"errors"
	"fmt"
)

func main() {}

func ErrWrap() any {
	err1 := errors.New("Oops Something happend.")
	err2 := fmt.Errorf("An error was encountered - %w", err1)
	return err2
}
