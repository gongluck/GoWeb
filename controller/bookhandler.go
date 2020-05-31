/*
 * @Author: gongluck 
 * @Date: 2020-05-31 15:49:12 
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-05-31 16:01:29
 */

package controller

import (
	"GoWeb/dao"
	"GoWeb/model"
	"html/template"
	"net/http"
	"strconv"
)

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

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookID := r.FormValue("bookId")
	dao.DeleteBook(bookID)
	GetBooks(w, r)
}

func ToUpdateBookPage(w http.ResponseWriter, r *http.Request) {
	bookID := r.FormValue("bookId")
	book, _ := dao.GetBookByID(bookID)
	if book.ID > 0{
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		t.Execute(w, book)
	}else{
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		t.Execute(w, "")
	}
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	bookID := r.PostFormValue("bookId")
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	price := r.PostFormValue("price")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")
	ibookID, _ := strconv.ParseInt(bookID, 10, 0)
	fprice, _:= strconv.ParseFloat(price, 64)
	isales, _ := strconv.ParseInt(sales, 10, 0)
	istock, _ := strconv.ParseInt(stock, 10, 0)
	book := &model.Book{
		ID:int(ibookID),
		Title:title,
		Author:author,
		Price:fprice,
		Sales:int(isales),
		Stock:int(istock),
		ImgPath:"/static/img/default.jpg",
	}
	if book.ID > 0{
		dao.UpdateBook(book)
	}else{
		dao.AddBook(book)
	}
	
	GetBooks(w, r)
}
