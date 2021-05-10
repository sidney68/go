package main

import (
    "fmt"
    "sync"
)

type Crawler struct {
    mutex sync.Mutex
    urls  map[string]bool
}

func New() *Crawler {
    return &Crawler{
        urls: make(map[string]bool),
    }
}

func (c *Crawler) visit(url string) bool {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    if _, ok := c.urls[url]; ok {
        return true
    }
    c.urls[url] = true
    return false
}

type Fetcher interface {
    // Fetch returns the body of URL and
    // a slice of URLs found on that page.
    Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func (c *Crawler) Crawl(url string, depth int, fetcher Fetcher) {
    var wg sync.WaitGroup

    b := c.visit(url)
    if b || depth <= 0 {
        return
    }
    body, urls, err := fetcher.Fetch(url)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("found: %s %q\n", url, body)
    for _, url := range urls {
        wg.Add(1)
        go func(url string) {
            defer wg.Done()
            c.Crawl(url, depth-1, fetcher)
        }(url)
    }

    wg.Wait()
}

func main() {
    crawler := New()
    crawler.Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
    body string
    urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
    if resp, ok := f[url]; ok {
        return resp.body, resp.urls, nil
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
