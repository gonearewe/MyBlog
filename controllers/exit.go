package controllers

package controllers

import (
	"MyBlog/models"
	"MyBlog/utils"
	"fmt"
	"github.com/astaxie/beego"
)

type ExitController struct {
	BaseController
}

func (c *ExitController) Get() {
	c.Data["title"] = "Log Out"
	c.DelSession("LoginUser")
	c.Redirect("/", 302)
}