/*
 * @Author: gongluck
 * @Date: 2020-05-30 13:44:52
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-01 13:52:52
 */

package controller

import (
	"GoWeb/dao"
	"GoWeb/model"
	"GoWeb/utils"
	"html/template"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	user, _ := dao.CheckUserNameAndPassword(username, password)
	if user.ID != 0 {
		sess := &model.Session{
			SessionID : utils.CreateUUID(),
			UserName : user.Username,
			UserID : user.ID,
		}
		dao.AddSession(sess)
		cookie := http.Cookie{
			Name:"user",
			Value:sess.SessionID,
			HttpOnly:true,
		}
		http.SetCookie(w, &cookie)
		t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
		t.Execute(w, user)
	} else {
		t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		t.Execute(w, "用户名或密码不正确！")
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("user")
	if cookie != nil{
		cookieValue := cookie.Value
		dao.DeleteSession(cookieValue)
		cookie.MaxAge = -1
	}
	GetPageBooksByPrice(w, r)
}

func Regist(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")

	user, _ := dao.CheckUserName(username)
	if user.ID != 0 {
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w, "用户名已存在！")
	} else {
		dao.SaveUser(username, password, email)
		t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
		t.Execute(w, "")
	}
}

func CheckUserName(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	user, _ := dao.CheckUserName(username)
	if user.ID != 0 {
		w.Write([]byte("用户名已存在!"))
	} else {
		w.Write([]byte("用户名可用!"))
	}
}
