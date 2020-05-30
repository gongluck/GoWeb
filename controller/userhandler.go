/*
 * @Author: gongluck 
 * @Date: 2020-05-30 13:44:52 
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-05-30 14:12:12
 */

package controller

import (
	"net/http"
	"GoWeb/dao"
	"html/template"
)

func Login(w http.ResponseWriter, r *http.Request){
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	user, _ := dao.CheckUserNameAndPassword(username, password)
	if user.ID != 0{
		t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
		t.Execute(w, "")
	}else{
		t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		t.Execute(w, "用户名或密码不正确！")
	}
}

func Regist(w http.ResponseWriter, r *http.Request){
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")

	user, _ := dao.CheckUserName(username)
	if user.ID != 0{
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w, "用户名已存在！")
	}else{
		dao.SaveUser(username, password, email)
		t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
		t.Execute(w, "")
	}
}
