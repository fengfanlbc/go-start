package main

import (
	"go-start/day02/interfaceMock"
	"fmt"
	"go-start/day02/interfaceReal"
	"time"
	"go-start/day02/queue"
)

type Receiver interface {
	Get(url string) string
}

func download(r Receiver, url string)  string{
	return r.Get(url)
}

func main() {
	var r Receiver
	r = interfaceMock.ReceiverMock{Content:"HELLO"}
	fmt.Printf("%T : %v \n", r, r)
	//url := "http://www.baidu.com"
	r = &interfaceReal.ReceiverReal{UserAgent:"Mozilla/5.0", TimeOut:time.Minute}
	fmt.Printf("%T : %v \n", r, r)
	//fmt.Println(download(r, url))
	q := queue.Queue{}
	q.Push(1)
	fmt.Println(q)
	q.Push("d")
	fmt.Println(q)
	fmt.Println(q.IsEmpty())
}
