package main

import (
	"bytes"
	"testing"
)

// TestCountWords test the count function set to count words
func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")

	exp := 4

	res := count(b)
	// 	If this function returns anything other than 4, the
	// test doesn’t pass and we raise an error that shows what we expected and
	// what we actually got instead.
	if res != exp {
		t.Errorf("Expected %d, got %d instead. \n", exp, res)
	}
}
