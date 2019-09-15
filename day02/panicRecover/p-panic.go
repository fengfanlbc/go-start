package main

import "fmt"

func tryRecover() int {
	defer func() {
		r := recover()
		if err, ok := r.(error) ; ok{
			fmt.Println(err)
		}else{
			panic(err)
		}
	}()
	b := 0
	var a = 5 / b
	return a
}

func main() {
	tryRecover()
}
