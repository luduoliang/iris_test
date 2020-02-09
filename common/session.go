package common

import (
	"iris_test/config"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris"
)

var (
	Sess = NewSession()
)

//返回session，单例
func NewSession() *sessions.Sessions{
	session_cookie_name := config.GetString("session_cookie_name", "")
	return sessions.New(sessions.Config{Cookie: session_cookie_name})
}


//获取
func GetSession(c iris.Context, key string) interface{} {
	session := Sess.Start(c)
	return session.Get(key)
}

//设置
func SetSession(c iris.Context, key string, value interface{}) {
	session := Sess.Start(c)
	session.Set(key, value)
}

//删除
func DelSession(c iris.Context, key string) {
	session := Sess.Start(c)
	session.Delete(key)
}

//清空
func ClearSession(c iris.Context) {
	session := Sess.Start(c)
	session.Clear()
}
