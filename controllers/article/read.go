package article

import (
	. "MyBlog/controllers"
	."MyBlog/utils"
	"MyBlog/models"
	"fmt"
	"io/ioutil"
	"strconv"
	"github.com/astaxie/beego"
)

type ArticleReadController struct {
	BaseController
}

func (c *ArticleReadController) Get() {
	id, err := c.GetInt("id")
	if err != nil {
		fmt.Println(err)
		c.Abort("404")
	}
	article := new(models.Article)
	article.Id = id
	article, err = article.GetById()
	if err != nil {
		fmt.Println(err)
		c.Abort("404")
	}
	filename := beego.AppConfig.String("articlepath") + strconv.Itoa(article.Id) + ".md"
	if content,err:=ioutil.ReadFile(filename);err!=nil{
		fmt.Printf("Can't Read Article File %s :",filename)
		fmt.Println(err)
	}else{
		article.Text=string(content)
	}
	c.Data["title"] = "Read Article"
	c.Data["article"] = article
	c.Data["text"]= MarkdownToHtml(article.Text)
	c.TplName = "article/read.html"

}
