package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	db_host := beego.AppConfig.String("db.host")
	db_user := beego.AppConfig.String("db.user")
	db_password := beego.AppConfig.String("db.password")
	db_port := beego.AppConfig.String("db.port")
	db_name := beego.AppConfig.String("db.database_name")

	if db_password == "" {
		orm.RegisterDataBase("default", "mysql", db_user+"@tcp("+db_host+":"+db_port+")/"+db_name)
	} else {
		orm.RegisterDataBase("default", "mysql", db_user+":"+db_password+"@tcp("+db_host+":"+db_port+")/"+db_name)
	}

	orm.RegisterModel(new(User))
}

func GetTableName(tableName string) string {
	return tableName
}
