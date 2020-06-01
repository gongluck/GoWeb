/*
 * @Author: gongluck
 * @Date: 2020-06-01 09:38:17
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-01 09:44:00
 */

package dao

import (
	"GoWeb/model"
	"testing"
)

func TestSession(t *testing.T) {
	t.Run("测试添加会话", testAddSession)
	t.Run("测试删除会话", testDeleteSession)
}

func testAddSession(t *testing.T) {
	session := model.Session{
		SessionID: "testsession",
		UserName:  "gongluck",
		UserID:    5,
	}
	AddSession(&session)
}

func testDeleteSession(t *testing.T) {
	DeleteSession("testsession")
}
