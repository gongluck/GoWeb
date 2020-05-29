/*
 * @Author: gongluck 
 * @Date: 2020-05-28 19:47:13 
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-05-28 20:34:01
 */

package model

import(
	"testing"
	"fmt"
)

func TestMain(m *testing.M){
	fmt.Println("开始测试用户模块：")
	m.Run()
}

func TestAddUser(t *testing.T){
	testAddUser(t)
}

func testAddUser(t *testing.T){
	user := &User{
		Username: "gongluck",
		Password: "hello",
		Email:"1039994845@qq.com",
	}
	user.AddUser()
}
