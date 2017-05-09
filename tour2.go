/**
 * go 语言信道
 */
package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) // 关闭队列
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	fmt.Println("1111111")
	go sum(s[:len(s)/2], c)
	fmt.Println("2222222")
	go sum(s[len(s)/2:], c)
	fmt.Println("3333333")
	x, y := <-c, <-c // 从 c 中获取数据
	fmt.Println(x, y, x+y)

	// range 与 close
	f := make(chan int, 10) // 创建带有缓冲区的管道
	go fibonacci(cap(f), f)
	for i := range f {
		fmt.Println(i)
	}
}