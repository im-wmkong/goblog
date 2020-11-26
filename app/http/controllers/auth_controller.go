package controllers

import (
	"fmt"
	"goblog/app/models/user"
	"goblog/app/requests"
	"goblog/pkg/view"
	"net/http"
)

type AuthController struct {

}

func (*AuthController) Register(w http.ResponseWriter, r *http.Request)  {
	view.RenderSimple(w, view.D{}, "auth.register")
}

func (*AuthController) DoRegister(w http.ResponseWriter, r *http.Request)  {
	// 1. 初始化数据
	_user := user.User{
		Name:            r.PostFormValue("name"),
		Email:           r.PostFormValue("email"),
		Password:        r.PostFormValue("password"),
		PasswordComfirm: r.PostFormValue("password_comfirm"),
	}

	// 2. 表单规则
	errs := requests.ValidateRegistrationForm(_user)

	if len(errs) > 0 {
		// 3. 表单不通过 —— 重新显示表单
		view.RenderSimple(w, view.D{
			"Errors": errs,
			"User":   _user,
		}, "auth.register")
	} else {
		// 4. 验证成功，创建数据
		_user.Create()

		if _user.ID > 0 {
			// 登录用户并跳转到首页
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "创建用户失败，请联系管理员")
		}
	}
}

func (*AuthController) Login(w http.ResponseWriter, r *http.Request)  {
	view.RenderSimple(w, view.D{}, "auth.login")
}

func (*AuthController) DoLogin(w http.ResponseWriter, r *http.Request)  {

}