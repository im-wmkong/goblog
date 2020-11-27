package auth

import (
	"errors"
	"goblog/app/models/user"
	"goblog/pkg/session"
	"gorm.io/gorm"
)

// getUID 获取字符串uid
func getUID() string {
	uid := session.Get("uid")
	if uidstr, ok := uid.(string); ok && len(uidstr) > 0 {
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

// Attempt 尝试登录
func Attempt(email, password string) error {
	_user, err := user.GetByEmail(email)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("账号不存在或密码错误")
		} else {
			return errors.New("内部错误，请稍后尝试")
		}
	}

	if !_user.ComparePassword(password) {
		return errors.New("账号不存在或密码错误")
	}
	session.Put("uid", _user.GetStringID())

	return nil
}

func Login(user user.User) {
	session.Put("uid", user.GetStringID())
}

func Logout() {
	session.Forget("uid")
}

func Check() bool {
	return len(getUID()) > 0
}
