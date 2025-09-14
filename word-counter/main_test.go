package main

import (
	"bytes"
	"testing"
)

// TestCountWords test the count function set to count words
func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")

	exp := 4

	res := count(b, false, false)
	// 	If this function returns anything other than 4, the
	// test doesn’t pass and we raise an error that shows what we expected and
	// what we actually got instead.
	if res != exp {
		t.Errorf("Expected %d, got %d instead. \n", exp, res)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\nline2\nline3 word1")
	exp := 3
	res := count(b, true, false)
	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}

// testCountBytes test count function set to count bytes
func testCountBytes(t *testing.T) {
	b := bytes.NewBufferString("123456789")
	exp := 9
	res := count(b, false, true)
	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}
