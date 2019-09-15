package main

import "fmt"

func main() {
	tryDefer()
	fmt.Println(3)
}

// 函数 return 前 defer 会执行
func tryDefer()  {
	defer fmt.Println(1)
	defer fmt.Println(4)
	fmt.Println(2)
}
