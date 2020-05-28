/*
 * @Author: gongluck
 * @Date: 2020-05-28 19:26:02
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-05-28 20:32:11
 */

package model

import (
	"GoWeb/utils"
	"fmt"
)

type User struct {
	ID       int
	Username string
	Password string
	Email    string
}

func (user *User) AddUser() error{
	sqlStr := "insert into users(username, password, email) values(?, ?, ?)"
	
	// inStmt, err := utils.Db.Prepare(sqlStr)
	// if err != nil{
	// 	fmt.Println("预编译异常：", err)
	// 	return err
	// }
	// _, err = inStmt.Exec("admin", "123456", "1039994845@qq.com")
	// if err != nil{
	// 	fmt.Println("插入异常：", err)
	// 	return err
	// }

	_, err := utils.Db.Exec(sqlStr, user.Username, user.Password, user.Email)
	if err != nil{
		fmt.Println("插入异常：", err)
		return err
	}
	
	return nil
}
