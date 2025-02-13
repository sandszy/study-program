package main

import "fmt"

func main() {
	//slice动态扩容，按照一定的算法扩容cap，扩容时会进行内存复制
	var s []int
	s = append(s, 11)
	fmt.Println(len(s), cap(s))
	s = append(s, 12)
	fmt.Println(len(s), cap(s))
	s = append(s, 13)
	fmt.Println(len(s), cap(s))
	s = append(s, 14)
	fmt.Println(len(s), cap(s))
	s = append(s, 15)
	fmt.Println(len(s), cap(s))

	u := []int{11, 12, 13, 14, 15}
	fmt.Println("array", u)
	x := u[1:3]
	fmt.Printf("slice(len=%d,cap=%d):%v\n", len(x), cap(x), x)

	x = append(x, 24)
	fmt.Println("after append 24, array", u)
	fmt.Printf("after append 24, slice(len=%d,cap=%d):%v\n", len(x), cap(x), x)

	x = append(x, 25)
	fmt.Println("after append 25, array", u)
	fmt.Printf("after append 25, slice(len=%d,cap=%d):%v\n", len(x), cap(x), x)

	//自动扩容后，x的底层数组变成了新的，在这之后，即便再修改切片中的元素值，原数组u的元素也没有发生任何改变，因为此时切片x与数组u已经解除了绑定关系，x已经不再是数组u的描述符了。
	x = append(x, 26)
	fmt.Println("after append 26, array", u)
	fmt.Printf("after append 26, slice(len=%d,cap=%d):%v\n", len(x), cap(x), x)

	x[0] = 22
	fmt.Println("after reassign 1st elem of slice, array", u)
	fmt.Printf("after reassign 1st elem of slice, slice(len=%d,cap=%d):%v\n", len(x), cap(x), x)
}
