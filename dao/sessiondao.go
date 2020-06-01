/*
 * @Author: gongluck 
 * @Date: 2020-06-01 09:30:24 
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-01 14:57:24
 */

package dao

import (
	"GoWeb/model"
	"GoWeb/utils"
	"net/http"
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

func IsLogin(r *http.Request) (bool, string) {
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		cookieValue := cookie.Value
		session, _ := GetSession(cookieValue)
		if session.UserID > 0 {
			return true, session.UserName
		}
	}
	return false, ""
}
