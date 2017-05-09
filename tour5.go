/**
 * daisy chain
 */

package main

import (
	"fmt"
)

func f(left, right chan int) {
	// 这个函数就把right的输出和left的输入联系起来了。
	left <- 1 + <-right
}

func main() {
	const n = 10000
	leftmost := make(chan int)
	right := leftmost
	left := leftmost
	fmt.Println("-----1")
	// 创建长度为n的daisy链
	for i := 0; i < n; i++ {

		right = make(chan int)
		go f(left, right)
		left = right
	}
	fmt.Println("-----2")
	// 在链的最右端输入1，那么最左端就会得到10001
	go func(c chan int) {
//		fmt.Println(<-c)
		c <- 1
	}(right)
	fmt.Println(<-leftmost)
}