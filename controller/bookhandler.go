/*
 * @Author: gongluck
 * @Date: 2020-05-31 15:49:12
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-01 14:55:08
 */

package controller

import (
	"GoWeb/dao"
	"GoWeb/model"
	"GoWeb/utils"
	"html/template"
	"net/http"
	"strconv"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	page, _ := dao.GetPageBooks(pageNo)
	t := template.Must(template.ParseFiles("views/index.html"))
	t.Execute(w, page)
}

func GetPageBooks(w http.ResponseWriter, r *http.Request) {
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	page, _ := dao.GetPageBooks(pageNo)
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	t.Execute(w, page)
}

func GetPageBooksByPrice(w http.ResponseWriter, r *http.Request) {
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	minPrice := r.FormValue("min")
	maxPrice := r.FormValue("max")
	var page *model.Page
	if minPrice == "" && maxPrice == "" {
		page, _ = dao.GetPageBooks(pageNo)
	} else {
		page, _ = dao.GetPageBooksByPrice(pageNo, minPrice, maxPrice)
		page.MinPrice = minPrice
		page.MaxPrice = maxPrice
	}
	flag, session := dao.IsLogin(r)
	if flag {
		page.IsLogin = true
		page.Username = session.UserName
	}

	t := template.Must(template.ParseFiles("views/index.html"))
	t.Execute(w, page)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	price := r.PostFormValue("price")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")
	fprice, _ := strconv.ParseFloat(price, 64)
	isales, _ := strconv.ParseInt(sales, 10, 0)
	istock, _ := strconv.ParseInt(stock, 10, 0)
	book := &model.Book{
		Title:   title,
		Author:  author,
		Price:   fprice,
		Sales:   int(isales),
		Stock:   int(istock),
		ImgPath: "/static/img/default.jpg",
	}
	dao.AddBook(book)
	GetPageBooks(w, r)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookID := r.FormValue("bookId")
	dao.DeleteBook(bookID)
	GetPageBooks(w, r)
}

func ToUpdateBookPage(w http.ResponseWriter, r *http.Request) {
	bookID := r.FormValue("bookId")
	book, _ := dao.GetBookByID(bookID)
	if book.ID > 0 {
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		t.Execute(w, book)
	} else {
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
	fprice, _ := strconv.ParseFloat(price, 64)
	isales, _ := strconv.ParseInt(sales, 10, 0)
	istock, _ := strconv.ParseInt(stock, 10, 0)
	book := &model.Book{
		ID:      int(ibookID),
		Title:   title,
		Author:  author,
		Price:   fprice,
		Sales:   int(isales),
		Stock:   int(istock),
		ImgPath: "/static/img/default.jpg",
	}
	if book.ID > 0 {
		dao.UpdateBook(book)
	} else {
		dao.AddBook(book)
	}

	GetPageBooks(w, r)
}

func AddBook2Cart(w http.ResponseWriter, r *http.Request) {
	bookID := r.FormValue("bookId")
	book, _ := dao.GetBookByID(bookID)
	_, session := dao.IsLogin(r)
	userID := session.UserID
	cart, _ := dao.GetCartByUserID(userID)
	if cart != nil {
		cartItem, _ := dao.GetCartItemByBookIDAndCartID(bookID, cart.CartID)
		if cartItem != nil {
			cts := cart.CartItems
			for _, v := range cts {
				if cartItem.Book.ID == v.Book.ID {
					v.Count = v.Count + 1
					dao.UpdateBookCount(v.Count, v.Book.ID, cart.CartID)
				}
			}
		} else {
			cartItem := &model.CartItem{
				Book:   book,
				Count:  1,
				CartID: cart.CartID,
			}
			cart.CartItems = append(cart.CartItems, cartItem)
			dao.AddCartItem(cartItem)
		}
		dao.UpdateCart(cart)
	} else {
		cart := &model.Cart{
			CartID: utils.CreateUUID(),
			UserID: userID,
		}
		var cartItems []*model.CartItem
		cartItem := &model.CartItem{
			Book:   book,
			Count:  1,
			CartID: cart.CartID,
		}
		cartItems = append(cartItems, cartItem)
		cart.CartItems = cartItems
		dao.AddCart(cart)
	}
	w.Write([]byte(book.Title))
}
