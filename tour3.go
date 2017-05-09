/**
 * go 语言 select
 */
package main
import (
	"fmt"
)

// go 线程设置c 管道
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			//y = y
			fmt.Println("aaaaaaa")
			x, y = y, x + y
		case <-quit:
			fmt.Println("quit")
			return
		default:
			//fmt.Println("    .")
		}
	}
}


func main() {
	c := make(chan int)
	quit := make(chan int)
	// go 线程读取c 管道
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(i+ 999999)

			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}