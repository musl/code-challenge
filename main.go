package main

import (
	"bufio"
	"io"
)

// TopNTriples returns the top N most frequently occuring triplets of
// three words in a given body of text.
func TopNTriples(n int, r io.Reader) (map[string]int, error) {
	triples := make(map[string]int)

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		// TODO: count triples found in this stream
	}

	if err := scanner.Err(); err != nil {
		return map[string]int{}, err
	}

	return triples, nil
}

func main() {
	// TODO: handle input, processing, and output
}
