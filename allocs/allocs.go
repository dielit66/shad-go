//go:build !solution

// DONE *
package allocs

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"sort"
)

// implement your Counter below

type EnhacedCounter struct {
	counts map[string]int
}

func NewEnhancedCounter() Counter {
	return EnhacedCounter{
		counts: map[string]int{},
	}
}

func (ec EnhacedCounter) Count(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()
		ec.counts[word]++
	}

	return scanner.Err()
}

func (ec EnhacedCounter) String() string {
	// Предварительно выделяем память для всех ключей
	keys := make([]string, 0, len(ec.counts))
	for word := range ec.counts {
		keys = append(keys, word)
	}
	sort.Strings(keys)

	var buf bytes.Buffer

	for _, key := range keys {
		fmt.Fprintf(&buf, "word '%s' has %d occurrences\n", key, ec.counts[key])
	}
	return buf.String()
}
