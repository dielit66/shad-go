//go:build !solution

// DONE

package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

type FetchProducer struct {
	client http.Client
	wg     sync.WaitGroup
}

type Fetcher interface {
	fetch(url string) error
}

func NewFetchProducer() *FetchProducer {
	return &FetchProducer{
		client: http.Client{},
	}
}

func main() {
	producer := NewFetchProducer()
	args := os.Args[1:]

	for i := 0; i < len(args); i++ {
		producer.wg.Add(1)
		go producer.fetch(args[i])
	}

	producer.wg.Wait()

}

func (fp *FetchProducer) fetch(url string) {
	defer fp.wg.Done()
	startTime := time.Now()
	r, err := fp.client.Get(url)

	if err != nil {
		fmt.Printf("Error while fetching url %s, error: %v\n", url, err)

		return
	}
	defer r.Body.Close()
	cl := r.Header.Get("content-length")

	fmt.Printf("%v     %s     %s\n", time.Since(startTime), cl, url)
}
