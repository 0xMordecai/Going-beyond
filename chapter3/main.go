package main

import (
	"errors"
)

func main() {}

func guess(number uint) (answer bool, err error) {
	if number > 99 {
		err = errors.New("Number is larger than 100")
	}
	// check if guess is correct
	return answer, err
}
