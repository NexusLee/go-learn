/**
 * go 语言线程
 */
package main
import (
	"fmt"
)

func say(s string) {
	for i := 0; i < 2; i++ {
		fmt.Println(i, s)
	}
}


func main() {
	go say("world")
	say("hello")
}