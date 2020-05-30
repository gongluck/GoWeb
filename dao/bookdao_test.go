/*
 * @Author: gongluck
 * @Date: 2020-05-30 21:49:00
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-05-30 21:54:33
 */

package dao

import (
	"fmt"
	"testing"
)

func TestBook(t *testing.T) {
	t.Run("测试获取图书", testGetBooks)
}

func testGetBooks(t *testing.T) {
	books, _ := GetBooks()
	for k, v := range books {
		fmt.Printf("第%v本书信息：%v\n", k+1, v)
	}
}
