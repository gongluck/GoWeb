/*
 * @Author: gongluck 
 * @Date: 2020-06-01 09:30:24 
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-01 10:23:00
 */

package dao

import (
	"GoWeb/model"
	"GoWeb/utils"
)

func AddSession(sess *model.Session) error {
	sqlStr:= "insert into sessions values(?, ?, ?)"
	_, err := utils.Db.Exec(sqlStr, sess.SessionID, sess.UserName, sess.UserID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteSession(sessID string) error {
	sqlStr:= "delete from sessions where session_id=?"
	_, err := utils.Db.Exec(sqlStr, sessID)
	if err != nil {
		return err
	}
	return nil
}

func GetSession(sessId string)(*model.Session, error){
	sqlStr:= "select session_id, username, user_id from sessions where session_id=?"
	row := utils.Db.QueryRow(sqlStr, sessId)
	sess := &model.Session{}
	err := row.Scan(&sess.SessionID, &sess.UserName, &sess.UserID)
	return sess, err
}