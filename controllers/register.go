/*
注册页面的控制层

@author liuyufeng
@version 1.0
@package controllers
*/
package controllers

import (
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Get() {
	c.TplName = "register.html"
	c.Layout = "common/basic-layout.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = "common/header.html"
	c.LayoutSections["HtmlFooter"] = "common/footer.html"

	c.Data["Isregister"] = true

}

func (c *RegisterController) Post() {

}
