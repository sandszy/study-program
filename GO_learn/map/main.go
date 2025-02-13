package main

import (
	"fmt"
)

func main() {
	//var m map[string]int // m = nil 不支持零值可用
	//m["key"] = 1         // panic: assignment to entry nil map

	// 使用复合字面值创建
	var n = map[string]int{
		"k1": 1,
		"k2": 2,
	}
	println(n["k1"])
	println(n["k2"])

	//和切片一样，map也是引用类型，将map类型变量作为函数参数传入不会有很大的性能损耗，并且在函数内部对map变量的修改在函数外部也是可见的
	foo(n)
	println(n["k1"])
	println(n["k2"])

	//删除数据
	m := map[string]int{
		"key1": 1,
		"key2": 2,
	}
	fmt.Println(m)
	delete(m, "key1")
	fmt.Println(m)

	//遍历数据
	//同一map多次遍历,顺序并不相同,因为Go运行时在初始化map迭代器时对起始位置做了随机处理
	//千万不要依赖遍历map所得到的元素次序
	o := map[int]int{
		1: 11,
		2: 22,
		3: 33,
	}
	fmt.Printf("{")
	for k, v := range o {
		fmt.Printf("[%d,%d]", k, v) //Printf不可以直接传入int，Println可以
	}
	fmt.Printf("}")
	//如果需要稳定的次序,用另外一种数据结构，按需要的次序保存key
	fmt.Println("\n如果需要稳定的次序,用另外一种数据结构，按需要的次序保存key")
	var s1 []int
	p := map[int]int{
		1: 11,
		2: 22,
		3: 33,
	}
	for k, _ := range p {
		s1 = append(s1, k) //将元素按初始次序保存在切片中
	}
	for i := 0; i < 100; i++ {
		doIteration(s1, p)
	}
}

func foo(m map[string]int) {
	m["k1"] = 3
	m["k2"] = 4
}

func doIteration(s1 []int, m map[int]int) {
	fmt.Printf("{")
	for _, k := range s1 {
		v, ok := m[k]
		if !ok {
			continue
		}
		fmt.Printf("[%d,%d]", k, v)
	}
	fmt.Printf("}\n")
}
