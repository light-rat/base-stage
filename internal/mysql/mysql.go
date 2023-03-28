package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"stage/http_server/dao"
)
func Init() {
	dsn := "root:kaixin2022@tcp(43.143.184.15:3306)/d_self?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
	db.NamingStrategy = schema.NamingStrategy{
		TablePrefix:   "t_",
		SingularTable: true,
	}
	dao.Init(db)
}