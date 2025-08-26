package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	str := "abcdefg"
	// str := "hello" // <----- will cause an error
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Fatalln("Cannot parse string", err)
	}
	fmt.Println("Number is", num)
}
