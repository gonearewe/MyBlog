package routers

import (
	"MyBlog/controllers"
	"MyBlog/controllers/article"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/exit", &controllers.ExitController{})

	beego.Router("/article/add", &article.ArticleAddController{})
	beego.Router("/article/view", &article.ArticleViewController{})
	beego.Router("/article/read", &article.ArticleReadController{})

}
