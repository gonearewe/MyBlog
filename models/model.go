package models
import (	
	"github.com/astaxie/beego/orm"
)

type User struct{
	Id int
	Username string
	Password string
	Createtime int64
	Status bool
}

func (u *User) Insert() (Id int64,err error) {
	return orm.NewOrm().Insert(u)
}
func (u *User) GetByName() (*User, error) {
	a := new(User)

	err := orm.NewOrm().QueryTable("user").Filter("Username", u.Username).One(a)
	if err != nil {
		return nil, err
	}
	return a, nil
}
func (u *User) Exist() bool {
	return orm.NewOrm().QueryTable("user").Filter("Username",u.Username).Exist()
}