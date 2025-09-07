package main

import (
	"errors"
	"fmt"
)

func main() {
	var guess_num uint = 55
	answer, err := guess(guess_num)
	if err != nil {
		panic(err)
	} else if answer == true {
		fmt.Printf("your guess is correct")
	} else {
		fmt.Printf("your guess is incorrect")
	}

}

func guess(number uint) (answer bool, err error) {
	if number > 99 {
		err = errors.New("Number is larger than 100")
	}
	// check if guess is correct
	return answer, err
}

// Using -->helpers<-- to simplify Repetitive Error Handling

// func check(err error, msg string) {
// 	if err != nil {
// 		log.Println("Error encountered:", msg)
// 	}
// }
