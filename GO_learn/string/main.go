package main

import "fmt"

func main() {
	//原始字符串
	var s string = "hello"
	fmt.Println("original string:", s)
	//切片化后试图改变原字符串
	s1 := []byte(s)
	s1[0] = 't'
	fmt.Println("slice:", string(s1))
	fmt.Println("after reslice,the original string is:", string(s)) //没有变化

}
