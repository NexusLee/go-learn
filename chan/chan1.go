package main

import (
	"sync"
	"fmt"
)

type threadSafeSet struct {
	sync.RWMutex
	s []interface{}
}

func (set *threadSafeSet) Iter() <-chan interface{} {
	 ch := make(chan interface{}) // 解除注释看看！
	//ch := make(chan interface{},len(set.s))
	go func() {
		set.RLock()

		for elem,value := range set.s {
			ch <- elem
			println("Iter:",elem,value)
		}

		close(ch)
		set.RUnlock()

	}()
	return ch
}

func main(){
	th := threadSafeSet{
		s : []interface{}{"1","2"},
	}
	v := <- th.Iter()
	fmt.Sprintf("%s%v","ch", v)
}

/*
输出:
	Iter: 0 (0x48bd80,0xc42000e200)
 */

//chan 无缓存只能写入1次
