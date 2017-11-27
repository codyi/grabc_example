package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	. "goCronJob/libs"
)

type User struct {
	Id         int    `json:"id" label:"id"`
	Phone      string `json:"phone" label:"用户手机号"`
	RealName   string `jsob:"real_name" label:"真实姓名"`
	Password   string `json:"password" label:"登录密码"`
	CreateTime int    `json:"create_time" label:"创建时间"`
	UpdateTime int    `json:"update_time" label:"更新时间"`
	IsLock     int    `json:"is_lock" label:"用户状态 0正常  1锁定"`
}

//return current model's table name
func (this *User) TableName() string {
	return GetTableName("users")
}

//Find one user by phone from database
func (this *User) FindByPhone(phone string) error {
	if phone == "" {
		return errors.New("用户手机号不能为空")
	}

	o := orm.NewOrm()
	return o.QueryTable(this.TableName()).Filter("phone", phone).One(this)
}

//validate user password is correct
func (this *User) ValidatePassword(password string) bool {
	fmt.Println(this.Password)
	fmt.Println(Md5(password))
	return this.Password == Md5(password)
}

//Modify user passowrd
func (this *User) ModifyPassword(password string) bool {
	if this == nil || this.Id <= 0 {
		return false
	}

	this.Password = Md5(password)

	o := orm.NewOrm()
	if _, err := o.Update(this); err == nil {
		return true
	}

	return false
}
