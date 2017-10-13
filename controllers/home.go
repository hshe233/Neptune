package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	username := c.GetSession("username")
	fmt.Println(username)
	c.TplName = "home.html"
}
