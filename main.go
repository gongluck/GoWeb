/*
 * @Author: gongluck
 * @Date: 2020-05-25 22:50:32
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-01 13:47:16
 */

package main

import (
	//"GoWeb/model"
	"fmt"
	"net/http"

	//"encoding/json"
	"GoWeb/controller"
	"html/template"
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
	//fmt.Fprintf(w, "通过自己创建的处理器处理请求！")

	// w.Header().Set("Content-Type", "application/json")
	// user := model.User{
	// 	ID:       1,
	// 	Username: "gongluck",
	// 	Password: "123456",
	// 	Email:    "1039994845@qq.com",
	// }
	// json, _ := json.Marshal(user)
	// w.Write(json)

	//重定向
	// w.Header().Set("Location", "https://gongluck.github.io")
	// w.WriteHeader(http.StatusFound)

	//模板
	t, _ := template.ParseFiles("views/index.html")
	//t.Execute(w, r.FormValue("name"))
	//t := template.Must(template.ParseFiles("index.html", "index2.html"))
	t.ExecuteTemplate(w, "index.html", "")
}

func handlerAction(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html", "index2.html")
	// t.ExecuteTemplate(w, "index.html", []model.User{
	// 	{ID : 1, Username : "gongluck", Password : "123456", Email : "1039994845@qq.com"},
	// 	{ID : 2, Username : "gong", Password : "888888", Email : "11111111@qq.com"},
	// 	{ID : 3, Username : "luck", Password : "654321", Email : "888888888@qq.com"},
	// })
	t.ExecuteTemplate(w, "model", "")
}

func handlerCookie(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "user",
		Value:    "test",
		HttpOnly: true,
		MaxAge:   10,
	}
	// cookie2 := http.Cookie{
	// 	Name:     "cookie2",
	// 	Value:    "more test",
	// 	HttpOnly: true,
	// }
	// w.Header().Set("Set-Cookie", cookie.String())
	// w.Header().Add("Set_Cookie", cookie2.String())

	http.SetCookie(w, &cookie)
}

func handlerGetCookie(w http.ResponseWriter, r *http.Request) {
	//cookies := r.Header["Cookie"]
	cookies, _ := r.Cookie("user")
	fmt.Println("Get all cookies:", cookies)
}

func main() {
	//http.HandleFunc("/", handler)

	http.HandleFunc("/action", handlerAction)
	http.HandleFunc("/cookie", handlerCookie)
	http.HandleFunc("/getcookie", handlerGetCookie)

	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/logout", controller.Logout)
	http.HandleFunc("/regist", controller.Regist)
	http.HandleFunc("/checkUserName", controller.CheckUserName)
	http.HandleFunc("/getPageBooks", controller.GetPageBooks)
	http.HandleFunc("/getPageBooksByPrice", controller.GetPageBooksByPrice)

	http.HandleFunc("/deleteBook", controller.DeleteBook)
	http.HandleFunc("/toUpdateBookPage", controller.ToUpdateBookPage)
	http.HandleFunc("/updateBook", controller.UpdateBook)

	http.HandleFunc("/main", controller.IndexHandler)

	//设置静态资源
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static/"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages/"))))

	// myHandler := MyHandler{}
	// http.Handle("/main", &myHandler)

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
