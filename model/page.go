/*
 * @Author: gongluck
 * @Date: 2020-05-31 16:07:21
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-05-31 16:24:41
 */

package model

type Page struct {
	Books       []*Book
	PageNo      int64
	PageSize    int64
	TotalPageNo int64
	TotalRecord int64
}
