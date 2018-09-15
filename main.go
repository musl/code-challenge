package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

// Triple is
type Triple struct {
	Value string
	Count int
}

// TripleHeap is a reverse-ordered heap.
type TripleHeap []Triple

// Len returns the number of elements in a heap.
func (t TripleHeap) Len() int { return len(t) }

// Less returns true if the value is greater, allowing the heap to sort
// in descending order.
func (t TripleHeap) Less(i, j int) bool { return t[i].Count > t[j].Count }

// Swap swaps the elements in a heap.
func (t TripleHeap) Swap(i, j int) { t[i], t[j] = t[j], t[i] }

// Push pushes a Triple onto a TripleHeap.
func (t *TripleHeap) Push(nt interface{}) {
	*t = append(*t, nt.(Triple))
}

// Pop pops a Triple off of a TripleHeap
func (t *TripleHeap) Pop() interface{} {
	old := *t
	n := len(old)
	x := old[n-1]
	*t = old[0 : n-1]
	return x
}

// TopNTriples returns a map containing the top zero to n most
// frequently occurring sequences of three words or triples.
func TopNTriples(n int, readers []io.Reader) ([]Triple, error) {

	pattern, err := regexp.Compile(`[^a-z]+`)
	if err != nil {
		return []Triple{}, err
	}

	tripleMap := make(map[string]int)

	for _, r := range readers {
		words := []string{}
		// Make use of bufio to efficiently process large files without
		// having to do the work ourselves.
		scanner := bufio.NewScanner(r)
		scanner.Split(bufio.ScanWords)

		for scanner.Scan() {
			word := pattern.ReplaceAllString(strings.ToLower(scanner.Text()), ``)
			if word == `` {
				continue
			}

			if len(words) == 3 {
				words = words[1:]
			}

			words = append(words, word)

			if len(words) == 3 {
				tripleMap[strings.Join(words, ` `)]++
			}
		}

		// On error, return the error so that consuming code can decide what
		// to do.
		if err := scanner.Err(); err != nil {
			return []Triple{}, err
		}
	}

	tripleHeap := &TripleHeap{}
	heap.Init(tripleHeap)

	for k, v := range tripleMap {
		heap.Push(tripleHeap, Triple{Value: k, Count: v})
	}

	// Limit the number of returned triples to the number of triples
	// found or n, whichever is smaller.
	o := len(tripleMap)
	if o > n {
		o = n
	}

	triples := make([]Triple, o)
	for i := 0; i < o; i++ {
		triples[i] = heap.Pop(tripleHeap).(Triple)
	}

	return triples, nil
}

func main() {

	// Note that program output goes to stout while logging goes to stderr.
	log.SetOutput(os.Stderr)

	// Make debugging nicer when looking at error messages, etc.
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	readers := []io.Reader{}

	// The first element of os.Args is the command name so we remove it.
	args := os.Args[1:]

	for _, f := range args {
		log.Printf("Reading from: %s", f)
		reader, err := os.Open(f)
		if err != nil {
			log.Fatal(err)
		}
		defer reader.Close()
		readers = append(readers, reader)
	}

	if len(readers) == 0 {
		log.Printf("Reading from standard input.")
		readers = []io.Reader{os.Stdin}
	}

	topk, err := TopNTriples(100, readers)
	if err != nil {
		log.Fatal(err)
	}

	for _, t := range topk {
		fmt.Printf("%d\t%v\n", t.Count, t.Value)
	}
}
