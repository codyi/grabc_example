package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root@tcp(127.0.0.1)/goCronJob")
	orm.RegisterModel(new(User))
}

func GetTableName(tableName string) string {
	return tableName
}
