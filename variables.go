package main

import "fmt"

func main() {
	var a = 15 // auto type
	fmt.Printf("a=%d\n", a)

	var b int // default: 0
	fmt.Printf("b=%d\n", b)

	var c bool // default: false
	fmt.Printf("c=%t\n", c)

	var d, e int // 同时声明多个，使用默认值
	fmt.Printf("d=%d, e=%d\n", d, e)

	var f, g int = 1, 2 // 同时声明多个，使用类型赋值
	fmt.Printf("f=%d, g=%d\n", f, g)

	var h, i = 3, 4.5 // 同时声明多个，自动识别类型
	fmt.Printf("h=%d, i=%f\n", h, i)

	j := 6 // 简短形式
	fmt.Printf("j=%d\n", j)

	k, l := 7, 8 // 简短形式多赋值
	fmt.Printf("k=%d, l=%d\n", k, l)

	k, l = l, k // 交换变量
	fmt.Printf("k=%d, l=%d\n", k, l)

	const HELLO string = "Hello world!"
	fmt.Printf("HELLO=%s\n", HELLO)
}
