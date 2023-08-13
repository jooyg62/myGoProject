package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mu sync.Mutex
	m  map[string]bool
}

func (sm SafeMap) Set(key string, value bool) {
	sm.mu.Lock()
	sm.m[key] = value
	sm.mu.Unlock()
}

func (sm SafeMap) Value(key string) (bool, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	v, ok := sm.m[key]
	return v, ok
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, c chan string, cm SafeMap) {

	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.

	defer close(c)

	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}

	_, ok := cm.Value(url)
	if ok {
		return
	} else {
		cm.Set(url, true)
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		c <- err.Error()
		return
	}

	c <- fmt.Sprintf("found: %s %q", url, body)

	results := make([]chan string, len(urls))
	for i, u := range urls {
		results[i] = make(chan string)
		go Crawl(u, depth-1, fetcher, results[i], cm)
	}

	for i := range results {
		for s := range results[i] {
			c <- s
		}
	}

	return
}

func main() {
	cm := SafeMap{m: make(map[string]bool)}

	c := make(chan string)
	go Crawl("https://golang.org/", 4, fetcher, c, cm)
	for s := range c {
		fmt.Println(s)
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
