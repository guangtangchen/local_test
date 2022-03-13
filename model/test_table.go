package model

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"
)

type TestTable struct {
	ID          int64     `gorm:"column:id"`
	OrderNo     string    `gorm:"column:order_no"`
	Name        string    `gorm:"column:name"`
	Content     string    `gorm:"column:content"`
	Num         int64     `gorm:"column:num"`
	GmtCreated  time.Time `gorm:"column:gmt_created"`
	GmtModified time.Time `gorm:"column:gmt_modified"`
}

// TableName 重写表名方法
func (*TestTable) TableName() string {
	return "test_table"
}

// TestTableDao 数据访问对象
type TestTableDao struct {
}

var once sync.Once
var testTableDao *TestTableDao

func TestTableDaoInstance() *TestTableDao {
	once.Do(func() {
		testTableDao = &TestTableDao{}
	})
	return testTableDao
}

func (dao *TestTableDao) Create(ctx context.Context, order *TestTable) (bool, error) {
	db := getMysqlClient(ctx)
	err := db.Create(order).Error
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate") {
			return true, nil
		}
		return false, err
	}
	fmt.Println("Create TestTable success = ", order)
	return false, nil
}
