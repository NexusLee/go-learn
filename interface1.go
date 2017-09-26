package main

import (
	"fmt"
)

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func live() People {
	var stu *Student
	return stu
}

func main() {
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
}

/*
输出:
	BBBBBBB
 */

/*
这个考点是很多人忽略的interface内部结构。 go中的接口分为两种一种是空的接口类似这样：

var in interface{}
另一种如题目：

type People interface {
	Show()
}
*/