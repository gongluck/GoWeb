/*
 * @Author: gongluck
 * @Date: 2020-06-01 09:23:06
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-01 09:27:29
 */

package model

type Session struct {
	SessionID string
	UserName  string
	UserID    int
	Cart      *Cart
}
