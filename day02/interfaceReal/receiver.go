package interfaceReal

import (
	"time"
	"net/http"
	"net/http/httputil"
)

type ReceiverReal struct {
	UserAgent string
	TimeOut time.Duration

}

// 使用指针接收者去实现的方法
func (r *ReceiverReal) Get(url string) string {
	/*
	实现 get 方法，就相当于实现了 Receiver 接口
	*/
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	result, err := httputil.DumpResponse(resp, true)
	resp.Body.Close()
	if err != nil {
		panic(err)
	}
	return string(result)
}

