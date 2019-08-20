package article

import (
	"github.com/astaxie/beego"
	"MyBlog/models"
	. "MyBlog/controllers"
	"fmt"
)

type ArticleViewController struct {
	BaseController
}

func (c *ArticleViewController) Get() {
	if pageIndex, err := c.GetInt64("page_index",1); err != nil {
		fmt.Println(err)
		c.Abort("404")
	}else{
		if pagenum,err:=models.ArticleNum();err!=nil{
			fmt.Println(err)
			c.Abort("404")
		}else{
			numPerPage,_:=beego.AppConfig.Int64("articlePerPage")
			if (pageIndex-1)*numPerPage>=pagenum{ //check if index of page is out of range
				fmt.Printf("Fail to Load %d : No Enough Article For That Index !!!\n",pageIndex)
				c.Abort("404")
			}
			if articles,err:=models.ListArticleByPage(numPerPage, pageIndex);err!=nil{
				fmt.Println(err)
				c.Abort("404")
			}else{
				c.Data["articles"]=articles //pass articles to template
				
				for _,i:=range articles{
					fmt.Println(*i)
				}
				
				pagination:=new(models.Pagination)
				pagination.Set(pageIndex,numPerPage,pagenum)

				fmt.Println(pagination)  //debug

				c.Data["pagination"]=pagination
			}
		}

	}

	c.Data["activeArticle"] = true
	c.Data["title"] = "View Article"
	c.TplName = "article/view.html"
}
