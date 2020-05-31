/*
 * @Author: gongluck
 * @Date: 2020-05-30 13:44:52
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-05-31 14:01:36
 */

package controller

import (
	"GoWeb/dao"
	"GoWeb/model"
	"html/template"
	"net/http"
	"strconv"
)

func Login(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	user, _ := dao.CheckUserNameAndPassword(username, password)
	if user.ID != 0 {
		t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
		t.Execute(w, "")
	} else {
		t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		t.Execute(w, "用户名或密码不正确！")
	}
}

func Regist(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")

	user, _ := dao.CheckUserName(username)
	if user.ID != 0 {
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w, "用户名已存在！")
	} else {
		dao.SaveUser(username, password, email)
		t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
		t.Execute(w, "")
	}
}

func CheckUserName(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	user, _ := dao.CheckUserName(username)
	if user.ID != 0 {
		w.Write([]byte("用户名已存在!"))
	} else {
		w.Write([]byte("用户名可用!"))
	}
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books, _ := dao.GetBooks()
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	t.Execute(w, books)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	price := r.PostFormValue("price")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")
	fprice, _:= strconv.ParseFloat(price, 64)
	isales, _ := strconv.ParseInt(sales, 10, 0)
	istock, _ := strconv.ParseInt(stock, 10, 0)
	book := &model.Book{
		Title:title,
		Author:author,
		Price:fprice,
		Sales:int(isales),
		Stock:int(istock),
		ImgPath:"/static/img/default.jpg",
	}
	dao.AddBook(book)
	GetBooks(w, r)
}
