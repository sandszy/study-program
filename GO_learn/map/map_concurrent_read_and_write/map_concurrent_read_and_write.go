package main

import (
	"fmt"
	"time"
)

//从上面的实现原理来看，充当map描述符角色的hmap实例自身是有状态的（hmap.flags）且对状态的读写是没有并发保护的，因此map实例不是并发写安全的，不支持并发读写。如果对map实例进行并发读写，程序运行时会发生panic

func doIteration(m map[int]int) {
	for k, v := range m {
		_ = fmt.Sprintf("[%d, %d] ", k, v)
	}
}
func doWrite(m map[int]int) {
	for k, v := range m {
		m[k] = v + 1
	}
}

func main() {
	m := map[int]int{
		1: 11,
		2: 12,
		3: 13,
	}

	go func() {
		for i := 0; i < 1000; i++ {
			doIteration(m)
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			doWrite(m)
		}
	}()

	time.Sleep(5 * time.Second)
}
