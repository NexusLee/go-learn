package main

import(
	"runtime"
	"sync"
	"fmt"
)

func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

/*
输出:
	i:  9
	i:  10
	i:  10
	i:  10
	i:  10
	i:  10
	i:  10
	i:  10
	i:  10
	i:  10
	i:  10
	i:  0
	i:  1
	i:  2
	i:  3
	i:  4
	i:  5
	i:  6
	i:  7
	i:  8
*/


//A WaitGroup waits for a collection of goroutines to finish. The main goroutine calls Add to set the number of goroutines to wait for. Then each of the goroutines runs and calls Done when finished. At the same time, Wait can be used to block until all goroutines have finished.