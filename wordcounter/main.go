package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// defining a boolean flag
	lines := flag.Bool("l", false, "Count lines in input")
	byteCount := flag.Bool("b", false, "Count byte size of text")

	//parsing flags provided by the user
	flag.Parse()

	//received from STDIN and printing it out.
	fmt.Println(count(os.Stdin, *lines, *byteCount))
}

func count(r io.Reader, countLines bool, countByte bool) int {
	//scanner used to read text from a Reader such as files
	scanner := bufio.NewScanner(r)

	//if the count lines flag is not set, we want to count words
	//so we define the scanner split type to words (default is split by lines)

	if !countLines {
		//define the scanner split type to words: default is lines
		scanner.Split(bufio.ScanWords)
	}

	// defining a counter
	wc := 0

	if countByte {
		scanner.Split(bufio.ScanBytes)
	}

	// for every word scanned, increase the counter
	for scanner.Scan() {
		wc++
		
	}

	return wc
}

//Compiling for different OS
// GOOS=windows go build
