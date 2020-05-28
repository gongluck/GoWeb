/*
 * @Author: gongluck
 * @Date: 2020-05-25 22:50:32
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-05-25 23:30:24
 */

package main

import (
	"fmt"
	"net/http"
)

// 创建处理器函数
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!", r.URL.Path)
}

type MyHandler struct{}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "通过自己创建的处理器处理请求！")
}

func main() {
	http.HandleFunc("/", handler)

	//myHandler := MyHandler{}
	//http.Handle("/myHandler", &myHandler)

	// 创建Server结构，并详细配置里面的字段
	// server := http.Server{
	// 	Addr:        ":8080",
	// 	Handler:     &myHandler,
	// 	ReadTimeout: 2 * time.Second,
	// }
	// server.ListenAndServe()

	// 创建多路复用器
	//mux := http.NewServeMux()
	//mux.HandleFunc("/", handler)

	// 创建路由
	//http.ListenAndServe(":8080", mux)
	http.ListenAndServe(":8080", nil)
}
