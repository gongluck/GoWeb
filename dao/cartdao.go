/*
 * @Author: gongluck 
 * @Date: 2020-06-01 16:19:53 
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-01 16:36:31
 */

package dao

import (
	"GoWeb/model"
	"GoWeb/utils"
)

func AddCart(cart *model.Cart) error {
	sqlStr:="insert into carts(id, total_count, total_amount, user_id) values(?, ?, ?, ?)"
	_, err := utils.Db.Exec(sqlStr, cart.CartID, cart.GetTotalCount(), cart.GetTotalAmount(), cart.UserID)
	for _, cartItem := range cart.CartItems {
		AddCartItem(cartItem)
	}
	return err
}