package session

import (
	"github.com/gorilla/sessions"
	"goblog/pkg/logger"
	"net/http"
)

var (
	Store = sessions.NewCookieStore([]byte("33446a9dcf9ea060a0a6532b166da32f304af0de"))
	Session *sessions.Session
	Request *http.Request
	Response http.ResponseWriter
)

func StartSession(w http.ResponseWriter, r *http.Request)  {
	var err error
	Session, err = Store.Get(r, "goblog-session")
	logger.LogError(err)
	Request = r
	Response = w
}

// Put 写入键值对应的会话数据
func Put(key string, value interface{})  {
	Session.Values[key] = value
	Save()
}

// Get 获取会话数据，获取数据时请做类型检测
func Get(key string) interface{} {
	return Session.Values[key]
}

// Forget 删除某个会话项
func Forget(key string)  {
	delete(Session.Values, key)
	Save()
}

func Flush()  {
	Session.Options.MaxAge = -1
	Save()
}

// Flush 删除当前会话
func Save()  {
	err := Session.Save(Request, Response)
	logger.LogError(err)
}