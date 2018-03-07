package main

import(
	"sync"
	"fmt"
)

func main() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}
	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		go func(url string) {
			// Fetch the URL.
			fmt.Println(url)
			// Decrement the counter.
			wg.Done()
		}(url)
	}
	fmt.Println("1")
	// Wait for all HTTP fetches to complete.
	wg.Wait()
	fmt.Println("2")
}

/*
输出:
	1
	http://www.somestupidname.com/
	http://www.golang.org/
	http://www.google.com/
	2
*/