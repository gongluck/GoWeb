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
	fmt.Fprintln(w, r.Method, r.URL.Path, r.Proto)
	fmt.Fprintln(w, r.URL.RawQuery)

	fmt.Fprintln(w, r.Header)
	fmt.Fprintln(w, r.Header["User-Agent"])
	fmt.Fprintln(w, r.Header.Get("User-Agent"))

	// bodylen := r.ContentLength
	// body := make([]byte, bodylen)
	// r.Body.Read(body)
	// fmt.Fprintln(w, string(body))

	//和以上的方法二选一,r.Body应该是只能Read一次
	// r.ParseForm()
	// fmt.Fprintln(w, r.Form)
	// fmt.Fprintln(w, r.PostForm)

	fmt.Fprintln(w, "name:", r.FormValue("name"))
	fmt.Fprintln(w, "password:", r.FormValue("password"))
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
