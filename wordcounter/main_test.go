package main

import (
	"bytes"
	"testing"
)

// TestCountWords tests the count function set to count words
func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")

	exp := 4

	res := count(b,false,false)

	if res != exp {
		t.Errorf("Expected %d, got %d instead", exp, res)
	}

}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2\n line2 word3 word4\nline3 word1")

	exp := 3

	res := count(b,true,false)

	if res != exp {
		t.Errorf("Expected %d, got %d instead", exp, res)
	}

}

func TestByteCount(t *testing.T){
	b := bytes.NewBufferString("There are five monkeys")
	
	exp := 22

	res := count(b, false,true)

	if res != exp {
		t.Errorf("Expected %d, got %d instead", exp, res)
	}
}
