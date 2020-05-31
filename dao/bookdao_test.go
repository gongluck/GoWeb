/*
 * @Author: gongluck
 * @Date: 2020-05-30 21:49:00
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-05-31 15:00:05
 */

package dao

import (
	"fmt"
	"testing"
	"GoWeb/model"
)

func TestBook(t *testing.T) {
	t.Run("测试获取所有图书", testGetBooks)
	t.Run("测试添加图书", testAddBook)
	t.Run("测试删除图书", testDeleteBook)
	t.Run("测试获取图书", testGetBook)
}

func testGetBooks(t *testing.T) {
	books, _ := GetBooks()
	for k, v := range books {
		fmt.Printf("第%v本书信息：%v\n", k+1, v)
	}
}

func testAddBook(t *testing.T) {
	book := &model.Book{
		Title:"三国演义",
		Author:"罗贯中",
		Price:88.88,
		Sales:100,
		Stock:100,
		ImgPath:"/static/img/default.jpg",
	}
	AddBook(book)
}

func testDeleteBook(t *testing.T) {
	DeleteBook("31")
}

func testGetBook(t *testing.T) {
	book, err := GetBookByID("10")
	fmt.Println(book, err)
}
