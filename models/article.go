package models

import "github.com/astaxie/beego/orm"

type Article struct {
	Id         int
	Title      string `orm:"unique"`
	Subtitle   string `orm:"null"`
	Author     string
	Tags       []string
	Text       string `orm:"type(text)"`
	Createtime int64
	Status     bool
}

func (u *Article) Insert() (Id int64, err error) {
	return orm.NewOrm().Insert(u)
}
func (u *Article) GetByTitle() (*Article, error) {
	a := new(Article)

	err := orm.NewOrm().QueryTable("article").Filter("Title", u.Title).One(a)
	if err != nil {
		return nil, err
	}
	return a, nil
}
func (u *Article) Exist() bool {
	return orm.NewOrm().QueryTable("article").Filter("Title", u.Title).Exist()
}
