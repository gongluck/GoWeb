/*
 * @Author: gongluck
 * @Date: 2020-05-28 16:55:28
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-05-28 17:01:11
 */

package utils

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	Db, err = sql.Open("mysql", "root:123456@localhost:3306/test")
	if err != nil {
		panic(err.Error())
	}
}
