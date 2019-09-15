package main

import (
	"net/http"
	"os"
	"io/ioutil"
	"icode.baidu.com/baidu/go-lib/log/log4go"
)

// 重新定义一个处理类，返回 error
type appHandler func(writer http.ResponseWriter, request * http.Request) error
// 专门处理 error 的函数，接收 handler 返回的 error
func errWrapper(handler appHandler) func(writer http.ResponseWriter, request * http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)
		if err != nil {
			log4go.Warn("Error Handling")
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			default:
				code = http.StatusInternalServerError

			}
			http.Error(writer,
				http.StatusText(code),
				code)
		}
	}
}

func main() {
	http.HandleFunc("/list/", errWrapper(handleFileList))
	err := http.ListenAndServe(":8080", nil )
	if err != nil {
		panic(err)
	}
}

func handleFileList(writer http.ResponseWriter, request *http.Request) error {
	path := request.URL.Path[len("/list/"):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	writer.Write(all)
	return nil
}