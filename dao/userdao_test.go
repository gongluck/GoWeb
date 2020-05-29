/*
 * @Author: gongluck
 * @Date: 2020-05-29 22:41:12
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-05-29 22:58:18
 */

package dao

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("开始测试用户登录模块：")
	m.Run()
}

func TestLogin(t *testing.T) {
	t.Run("测试登录用户", testLogin)
}

func testLogin(t *testing.T) {
	user, err := CheckUserName("gongluck")
	fmt.Println(err, user)
}
