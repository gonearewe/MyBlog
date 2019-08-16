package controllers

import (
	"github.com/astaxie/beego"
	"MyBlog/models"
	"fmt"
	"strings"
	"time"
)

type ArticleAddController struct {
	BaseController
}
type ArticleUpdateController struct {
	BaseController
}

func (c *ArticleAddController) Get() {
	c.Data["activeArticle"] = true
	c.Data["title"] = "Add Article"
	c.TplName = "article_add.html"
}
func (c *ArticleAddController) Post() {
	newArticle := new(models.Article)
	newArticle.Title = c.GetString("title")
	newArticle.Subtitle = c.GetString("subtitle")
	newArticle.Author = c.GetString("author")
	newArticle.Tags = strings.Split(c.GetString("tags"), "&")

	for {
		switch {
		case newArticle.Title == "":
			errorReport(c, "文章标题为空")
			return
		case newArticle.Author == "":
			errorReport(c, "文章作者为空")
			return
		case newArticle.Tags == nil || newArticle.Tags[0] == "":
			errorReport(c, "标签为空")
			return
		default:
			break
		}
	}
	newArticle.Createtime = time.Now().Unix()
	newArticle.Status = true
	if id, err := newArticle.Insert(); err != nil {
		fmt.Println(err)
		c.Data["json"] = map[string]interface{}{"code": 8, "message": "数据库插入失败"}
		c.ServeJSON()
		return
	}else{
		c.SaveToFile("article", beego.AppConfig.String("articlepath")+string(id)+".md")
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "文章添加成功"}
		c.ServeJSON()
		return
	}

}
func errorReport(c *ArticleAddController, s string) {
	c.Data["json"] = map[string]interface{}{"code": 7, "message": s}
	c.ServeJSON()
	return
}
