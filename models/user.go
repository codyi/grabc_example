package models

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	. "grabc_example/libs"
)

func init() {
	db_host := beego.AppConfig.String("db.host")
	db_user := beego.AppConfig.String("db.user")
	db_password := beego.AppConfig.String("db.password")
	db_port := beego.AppConfig.String("db.port")
	db_name := beego.AppConfig.String("db.database_name")

	orm.RegisterDataBase("default", "mysql", db_user+":"+db_password+"@tcp("+db_host+":"+db_port+")/"+db_name)
	orm.RegisterModel(new(User))
}

type User struct {
	Id       int    `json:"id" label:"id"`
	Phone    string `json:"phone" label:"用户手机号"`
	RealName string `jsob:"real_name" label:"真实姓名"`
	Password string `json:"password" label:"登录密码"`
}

//return current model's table name
func (this *User) TableName() string {
	return "users"
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
	return this.Password == Md5(password)
}

//retrieve all user name and ids
func (this User) UserList(pageIndex, pageCount int) (userList map[int]string, totalNum int, err error) {
	userList = make(map[int]string, 0)
	var users []*User
	var total int64
	o := orm.NewOrm()
	_, err = o.QueryTable(this.TableName()).Limit(pageCount).Offset(pageCount * (pageIndex - 1)).All(&users)

	if err != nil {
		return userList, int(total), err
	} else {
		for _, u := range users {
			userList[u.Id] = u.RealName
		}
	}

	total, err = o.QueryTable(this.TableName()).Count()
	return userList, int(total), err
}

//Find user by id
func (this *User) FindById(id int) error {
	o := orm.NewOrm()
	return o.QueryTable(this.TableName()).Filter("id", id).One(this)
}

func (this User) GetId() int {
	return this.Id
}

//Find user by id
func (this User) FindNameById(id int) string {
	o := orm.NewOrm()
	err := o.QueryTable(this.TableName()).Filter("id", id).One(&this)

	if err != nil {
		return ""
	}

	return this.RealName
}
