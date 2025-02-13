package main

import "fmt"

const (
	a = 5
	s = "hello,Gopher"
)

func main() {
	n := a
	var i interface{} = a

	fmt.Printf("%T\n", n)
	fmt.Printf("%T\n", i)
	i = s
	fmt.Printf("%T\n", i)
}
