/*
 * @Author: gongluck
 * @Date: 2020-05-28 19:47:13
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-05-28 20:34:01
 */

package model

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("开始测试用户模块：")
	m.Run()
}

func TestAddUser(t *testing.T) {
	t.Run("测试添加用户", testAddUser)
	t.Run("测试查询用户", testGetUserById)
	t.Run("测试查询所有用户", testGetUsers)
}

func testAddUser(t *testing.T) {
	user := &User{
		Username: "gongluck",
		Password: "hello",
		Email:    "1039994845@qq.com",
	}
	user.AddUser()
}

func testGetUserById(t *testing.T) {
	user := &User{
		ID: 1,
	}
	u, _ := user.GetUserById()
	fmt.Println("查询结果：", u)
}

func testGetUsers(t *testing.T) {
	var user User
	users, _ := user.GetUsers()
	for k, v := range users {
		fmt.Printf("第%v个用户是%v\n", k+1, v)
	}
}
