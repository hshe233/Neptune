package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "login.html"
}

func (c *LoginController) Login() {

	username := c.GetString("username")
	password := c.GetString("password")

	//校验门户的登录用户名和密码
	if checkID(username, password) {
		c.Redirect("/home", 302)

		//在Session中给其他系统的用户名密码赋值
		c.initID(username, password)
	} else {
		c.Redirect("/", 302)
	}
	return
}

func checkID(username, password string) bool {
	if username == "admin" && password == "admin" {
		return true
	} else {
		return false
	}
}

func (c *LoginController)initID(username, password string) {

	c.SetSession("username", username)
	c.SetSession("password", password)
	c.SetSession("username_edap","123")
	c.SetSession("password_edap","123")

}
