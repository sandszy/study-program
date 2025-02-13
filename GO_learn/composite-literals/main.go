package main

import (
	"fmt"
)

// 复合字面值由两部分组成：
// 一部分是类型，比如上述示例代码中赋值操作符右侧的myStruct、​[5]int、​[​]int和map[int]string；
// 另一部分是由大括号{}包裹的字面值。
// Go推荐使用field:value的复合字面值形式对struct类型变量进行值构造，这种值构造方式可以降低结构体类型使用者与结构体类型设计者之间的耦合，这也是Go语言的惯用法。
func main() {
	a := [5]int{13, 14, 15, 16, 17}
	m := map[int]string{1: "hello", 2: "gopher", 3: "!"}

	fmt.Println(a)
	fmt.Println(m)

	//结构体复合字面值
	type person struct {
		id   int
		name string
	}

	p := person{
		id:   330483198810260513,
		name: "沈震宇",
	}

	fmt.Println(p)

	//素组和切片的复合字面值
	//好像是按照字母的ascii码来代表下标
	numbers := [256]int{'a': 8, 'b': 7, 'c': 4, 'd': 3, 'e': 2, 'y': 1, 'x': 5}

	// [10]float{-1, 0, 0, 0, -0.1, -0.1, 0, 0.1, 0, -1}
	fnumbers := [...]float64{-1, 4: -0.1, -0.1, 7: 0.1, 9: -1}

	// $GOROOT/src/sort/search_test.go
	var data = []int{0: -10, 1: -5, 2: 0, 3: 1, 4: 2, 5: 3, 6: 5, 7: 7,
		8: 11, 9: 100, 10: 100, 11: 100, 12: 1000, 13: 10000}
	var sdata = []string{0: "f", 1: "foo", 2: "foobar", 3: "x"}

	fmt.Println(numbers)
	fmt.Println(fnumbers)
	fmt.Println(data)
	fmt.Println(sdata)

	//map复合字面值
	type Duration int64
	const (
		Nanosecond  Duration = 1
		Microsecond          = 1000 * Nanosecond
		Millisecond          = 1000 * Microsecond
		Second               = 1000 * Millisecond
		Minute               = 60 * Second
		Hour                 = 60 * Minute
	)
	var unitMap = map[string]int64{
		"ns": int64(Nanosecond),
		"us": int64(Millisecond),
		"ms": int64(Microsecond),
	}
	fmt.Println(unitMap)

	type Point struct {
		x float64
		y float64
	}

	s1 := []Point{
		{1.2345, 6.2789},
		{2.2345, 16.2789},
	}
	fmt.Println(s1)

}
