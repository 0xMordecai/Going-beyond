package main

import (
	"fmt"

	"github.com/dustin/go-humanize"
)

func main() {
	var number uint64 = 123456789
	fmt.Println("Size of file is", humanize.Bytes(number))
}
