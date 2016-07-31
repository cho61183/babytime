/*
	登录页面的控制

	@author liuyufeng
	@version 1.0
	@package controllers
*/
package controllers

import (
	"github.com/astaxie/beego"
	"github.com/cho61183/babytime/models"
)

type LoginController struct {
	beego.Controller
}

type LoginInfo struct {
	Name interface{} `form:"username"`
	Pwd  string      `form:"password"`
}

func (c *LoginController) Get() {
	device_id := models.Get_user_info_name(16)
	c.TplName = "login.html"
	c.Layout = "common/basic-layout.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = "common/header.html"
	c.LayoutSections["HtmlFooter"] = "common/footer.html"
	c.Data["Islogin"] = true
	c.Data["Device_id"] = device_id
}

func (c *LoginController) Post() {
	login := LoginInfo{}
	if err := c.ParseForm(&login); err != nil {
		c.Ctx.WriteString("Para_error")
	} else {
		c.Ctx.WriteString("Name:=" + login.Name.(string) + "\nPwd:=" + login.Pwd)
	}
	//c.Redirect("/", 301)
}
