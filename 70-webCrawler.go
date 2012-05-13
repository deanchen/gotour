package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "regexp"
    "net/url"
    "strings"
)
var CRAWL_ROOT = "http://golang.org"
var LINKS_PER_PAGE = 2000
var DEPTH = 2
var anchorRegex, _ = regexp.Compile("<a.*?href=['\"](.*?)['\"].*?>.*?</a>")

type Fetcher interface {
    // Fetch returns the body of URL and
    // a slice of URLs found on that page.
    Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(parenturl string, crawlurl string, depth int, fetcher *MyFetcher, parentChannel chan *Result) {
    // Don't fetch the same URL twice.
    if _, exists := (*fetcher)[crawlurl]; exists || depth <= 0 {
        parentChannel <- nil
        return
    }


    fmt.Println("Fetching: " + crawlurl)
    // Fetch URLs in parallel.
    _, urls, err := fetcher.Fetch(crawlurl)
    if err != nil {
        fmt.Println(err)
        parentChannel <- &Result{crawlurl, "", urls, false, parenturl}
        return
    }

    channel := make(chan *Result)

    sent := 0
    for _, u := range urls {
        // don't crawl too far past site boundaries
        if strings.Index(crawlurl, CRAWL_ROOT) != -1 {
            sent++
            go Crawl(crawlurl, u, depth-1, fetcher, channel)
        }
    }

    for i := 0; i < sent; i++ {
        result := <-channel
        if result != nil {
            (*fetcher)[result.url] = result
        }
    }

    parentChannel <- &Result{crawlurl, "", urls, true, parenturl}
    return
}

func main() {
    masterChannel := make(chan *Result)
    go Crawl("", CRAWL_ROOT, DEPTH, fetcher, masterChannel)
    result := <-masterChannel
    if result != nil {
        (*fetcher)[result.url] = result
    }
    fmt.Println("\n" + "####### Crawl Results ########")
    for rooturl, results := range (*fetcher) {
        if results.status {
            fmt.Println("-", rooturl)
        } else {
            fmt.Println("*", rooturl, results.parentUrl)
        }
        for _, leafurl := range results.urls {
            //fmt.Println("-", leafurl)
            leafurl = leafurl
        }
    }
}


type MyFetcher map[string]*Result

type Result struct {
    url string
    body string
    urls []string
    status bool
    parentUrl string
}

func (f *MyFetcher) Fetch(requestUrl string) (string, []string, error) {
    baseurl, _ := url.ParseRequestURI(requestUrl)
    if res, err := http.Get(requestUrl); err == nil {
        defer res.Body.Close()
        bytes, ioerr := ioutil.ReadAll(res.Body)
        if ioerr == nil {
            body := string(bytes)
            results := anchorRegex.FindAllStringSubmatch(body, LINKS_PER_PAGE)
            urls := make([]string, len(results))
            for i, result := range results {
                resultUrl := result[1]
                if len(resultUrl) > 0 {
                    if strings.Index(resultUrl, "http") != 0 {
                        if resultUrl[0] != 47 {
                            resultUrl = baseurl.Scheme + "://" + baseurl.Host + "/" + resultUrl
                        } else {
                            resultUrl = baseurl.Scheme + "://" + baseurl.Host + resultUrl
                        }
                    }
                urls[i] = resultUrl
                }
            }
            return body, urls, ioerr
        }
        return "", nil, ioerr
    } else {
        return "", nil, err
    }
    return "", nil, fmt.Errorf("request failed", requestUrl)
}

var fetcher = &MyFetcher{}
/*
// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
    body string
    urls []string
}

func (f *fakeFetcher) Fetch(url string) (string, []string, error) {
    if res, ok := (*f)[url]; ok {
        return res.body, res.urls, nil
    }
    return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = &fakeFetcher{
    "http://golang.org/": &fakeResult{
        "The Go Programming Language",
        []string{
            "http://golang.org/pkg/",
            "http://golang.org/cmd/",
        },
    },
    "http://golang.org/pkg/": &fakeResult{
        "Packages",
        []string{
            "http://golang.org/",
            "http://golang.org/cmd/",
            "http://golang.org/pkg/fmt/",
            "http://golang.org/pkg/os/",
        },
    },
    "http://golang.org/pkg/fmt/": &fakeResult{
        "Package fmt",
        []string{
            "http://golang.org/",
            "http://golang.org/pkg/",
        },
    },
    "http://golang.org/pkg/os/": &fakeResult{
        "Package os",
        []string{
            "http://golang.org/",
            "http://golang.org/pkg/",
        },
    },
}
*/
