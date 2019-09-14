package article

import (
	"MyBlog/models"
	."MyBlog/controllers"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
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
	c.TplName = "article/add.html"
}
func (c *ArticleAddController) Post() {
	if c.IsLogin==false{
		c.Redirect("/login", 302)
		return
	}
	if c.GetString("title") == "" {
		c.Data["json"] = map[string]interface{}{"code": 7, "message": "文章标题为空"}
		c.ServeJSON()
		return
	}

	newArticle := new(models.Article)
	// tags:=strings.Split(c.GetString("tags"), "&")
	// for _,tag:=range tags{
	// 	if s:=strings.TrimSpace(tag);s!=""{
	// 		var i int
	// 		newArticle.Tags[i]=s
	// 		i++
	// 	}
	// }
	// if newArticle.Tags[0]==""{
	// 	c.Data["json"] = map[string]interface{}{"code": 7, "message": "文章标签为空"}
	// 	c.ServeJSON()
	// 	return
	// }

	newArticle.Title = c.GetString("title")
	newArticle.Subtitle = c.GetString("subtitle")
	newArticle.Createtime = time.Now().Unix()
	newArticle.Status = true
	if user, ok := c.LoginUser.(string); ok {
		newArticle.Author = user
	}

	if id, err := newArticle.Insert(); err != nil {
		fmt.Println(err)
		c.Data["json"] = map[string]interface{}{"code": 8, "message": "数据库插入失败"}
		c.ServeJSON()
		return
	} else {
		_, _, err := c.GetFile("article")
		if err != nil {
			if newArticle.Delete()!=nil{
				fmt.Println(err)
			}
			c.Data["json"] = map[string]interface{}{"code": 9, "message": "文章上传失败"}
			c.ServeJSON()
			return
		}
		path := beego.AppConfig.String("articlepath") + strconv.FormatInt(id,10) + ".md"
		fmt.Println(beego.AppConfig.String("articlepath"))
		fmt.Println(path)
		c.SaveToFile("article", path)
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "文章上传成功", "article_id": id}
		c.ServeJSON()
		return
	}

}
