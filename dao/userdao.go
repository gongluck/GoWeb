/*
 * @Author: gongluck
 * @Date: 2020-05-29 22:28:16
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-05-29 23:02:35
 */

package dao

import (
	"GoWeb/model"
	"GoWeb/utils"
)

func CheckUserName(username string) (*model.User, error) {
	sqlStr := "select id, username, password, email from users where username=?"
	row := utils.Db.QueryRow(sqlStr, username)
	user := &model.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	return user, err
}

func SaveUser(username string, password string, email string) error {
	sqlStr := "insert into users(username, password, email) values(?, ?, ?)"
	_, err := utils.Db.Exec(sqlStr, username, password, email)
	return err
}
