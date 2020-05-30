/*
 * @Author: gongluck
 * @Date: 2020-05-30 21:29:16
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-05-30 21:31:11
 */

package model

type Book struct {
	ID      int
	Title   string
	Author  string
	Price   float64
	Sales   int
	Stock   int
	ImgPath string
}
