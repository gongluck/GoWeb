/*
 * @Author: gongluck
 * @Date: 2020-05-25 22:50:32
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-05-25 22:55:43
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

func main() {
	http.HandleFunc("/", handler)

	// 创建路由
	http.ListenAndServe(":8080", nil)
}
