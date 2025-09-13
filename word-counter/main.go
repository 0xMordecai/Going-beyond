package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Calling the count function to count the number of words
	// received from the Standard Input and printing it out
	fmt.Println(count(os.Stdin))
}

func count(r io.Reader) int {
	return 1
}
