/*
 * @Author: gongluck 
 * @Date: 2020-06-01 15:22:20 
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-01 15:26:22
 */

package model

type CartItem struct {
	CartItemID int64
	Book *Book
	Count int64
	Amount float64
	CartID string
}

func (cartItem *CartItem)GetAmount() float64 {
	price := cartItem.Book.Price
	return float64(cartItem.Count) * price
}
