package model

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var mysqlInstance *gorm.DB

func InitMysql(ctx context.Context) {
	username := "root"
	password := "376227"
	host := "82.156.205.9"
	port := 3306
	Dbname := "test"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, Dbname)
	var err error
	mysqlInstance, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("mysql connected")
	fmt.Println(mysqlInstance.Raw("show tables"))
}

func getMysqlClient(ctx context.Context) *gorm.DB {
	return mysqlInstance.WithContext(ctx)
}
