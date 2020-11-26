package auth

import (
	"goblog/app/models/user"
	"goblog/pkg/session"
)

// getUID 获取字符串uid
func getUID() string {
	uid := session.Get("uid")
	if uidstr, ok := uid.(string); ok && len(uidstr) > 0{
		return uidstr
	}
	return ""
}

// User 获取登录用户信息
func User() user.User {
	uid := getUID()
	if len(uid) > 0 {
		_user, err := user.Get(uid)
		if err == nil {
			return _user
		}
	}
	return user.User{}
}
//
//func ()  {
//
//}