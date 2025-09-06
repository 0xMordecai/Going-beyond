package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "123456789"
	// str := "hello" // <----- will cause an error
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println("Number is", num)
}
