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

func (p *Page)IsHasPrev() bool {
	return p.PageNo > 1
}

func (p *Page)IsHasNext() bool {
	return p.PageNo < p.TotalPageNo
}

func (p *Page)GetPrevPageNo() int64 {
	if p.IsHasPrev() {
		return p.PageNo -1
	} else {
		return 1
	}
}

func (p *Page)GetNextPageNo() int64 {
	if p.IsHasNext() {
		return p.PageNo +1
	} else {
		return p.TotalPageNo
	}
}
