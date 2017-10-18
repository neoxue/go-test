package main

import "fmt"

type Fetcher interface {
	Fetch(url string) (body string, item []string, err error)
}

func Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		Crawl(u, depth - 1, fetcher)
	}
	return
}

/*
func main() {
	Crawl("http://sina.cn/", 4, fetcher)
}
*/

type fakeFetcher map[string] *fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) {

}



