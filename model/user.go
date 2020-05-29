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

func (user *User) AddUser() error {
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
	if err != nil {
		fmt.Println("插入异常：", err)
		return err
	}

	return nil
}

func (user *User) GetUserById() (*User, error) {
	sqlStr := "select id, username, password, email from users where id=?"
	row := utils.Db.QueryRow(sqlStr, user.ID)

	var id int
	var username string
	var password string
	var email string
	err := row.Scan(&id, &username, &password, &email)
	if err != nil {
		fmt.Println("查询异常：", err)
		return nil, err
	}
	u := &User{
		ID:       id,
		Username: username,
		Password: password,
		Email:    email,
	}
	return u, err
}

func (user *User) GetUsers() ([]*User, error) {
	sqlStr := "select id, username, password, email from users"
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		fmt.Println("查询异常：", err)
		return nil, err
	}

	var id int
	var username string
	var password string
	var email string
	var users []*User
	for rows.Next() {
		err := rows.Scan(&id, &username, &password, &email)
		if err != nil {
			fmt.Println("查询-异常：", err)
			return nil, err
		}
		u := &User{
			ID:       id,
			Username: username,
			Password: password,
			Email:    email,
		}
		users = append(users, u)
	}

	return users, nil
}
