/*
 * @Author: gongluck 
 * @Date: 2020-06-01 16:13:20 
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-01 16:20:36
 */

package dao

import (
	"GoWeb/model"
	"GoWeb/utils"
)

func AddCartItem(cartItem *model.CartItem) error {
	sqlStr := "inset into cart_items(count, amount, book_id, cart_id) values(?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, cartItem.Count, cartItem.GetAmount(), cartItem.Book.ID, cartItem.CartID)
	return err
}
