/*
 * @Author: gongluck
 * @Date: 2020-05-30 21:44:15
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-05-30 21:53:39
 */

package dao

import (
	"GoWeb/model"
	"GoWeb/utils"
)

func GetBooks() ([]*model.Book, error) {
	sqlStr := "select id, title, author, price, sales, stock, img_path from books"
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}

	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}

	return books, nil
}
