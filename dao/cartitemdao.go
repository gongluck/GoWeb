/*
 * @Author: gongluck
 * @Date: 2020-06-01 16:13:20
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-01 16:56:47
 */

package dao

import (
	"GoWeb/model"
	"GoWeb/utils"
)

func AddCartItem(cartItem *model.CartItem) error {
	sqlStr := "insert into cart_items(count, amount, book_id, cart_id) values(?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, cartItem.Count, cartItem.GetAmount(), cartItem.Book.ID, cartItem.CartID)
	return err
}

func GetCartItemByBookID(bookID string) (*model.CartItem, error) {
	sqlStr := "select id, count, amount, cart_id from cart_items where book_id=?"
	row := utils.Db.QueryRow(sqlStr, bookID)
	cartItem := &model.CartItem{}
	err := row.Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount, &cartItem.CartID)
	if err != nil {
		return nil, err
	}
	return cartItem, nil
}

func GetCartItemsByCartID(CartID string) ([]*model.CartItem, error) {
	sqlStr := "select id, count, amount, book_id, cart_id from cart_items where cart_id=?"
	rows, err := utils.Db.Query(sqlStr, CartID)
	if err != nil {
		return nil, err
	}
	var cartItems []*model.CartItem
	for rows.Next() {
		var bookID string
		cartItem := &model.CartItem{}
		err := rows.Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount, &bookID, &cartItem.CartID)
		if err != nil {
			return nil, err
		}
		book, _ := GetBookByID(bookID)
		cartItem.Book = book
		cartItems = append(cartItems, cartItem)
	}
	return cartItems, nil
}

func GetCartItemByBookIDAndCartID(bookID, cardID string) (*model.CartItem, error) {
	sqlStr := "select id, count, amount from cart_items where book_id=? and cart_id=?"
	row := utils.Db.QueryRow(sqlStr, bookID, cardID)
	cartItem := &model.CartItem{}
	err := row.Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount)
	if err != nil {
		return nil, err
	}
	book, _ := GetBookByID(bookID)
	cartItem.Book = book
	return cartItem, nil
}

func UpdateBookCount(bookCount int64, bookID int, cartID string) error {
	sqlStr := "update cart_items set count=? where book_id=? and cart_id=?"
	_, err := utils.Db.Exec(sqlStr, bookCount, bookID, cartID)
	return err
}
