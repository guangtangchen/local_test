package model

import (
	"context"
	"fmt"
	"github.com/guangtangchen/common/util"
	"github.com/hedzr/assert"
	"testing"
	"time"
)

func TestTestTableDao_Create(t *testing.T) {
	InitMysql(context.Background())
	order := &TestTable{
		OrderNo:     util.GenId(),
		Name:        "chenguangtang" + util.GenId(),
		Content:     "今天在罗庄西里",
		Num:         1,
		GmtCreated:  time.Now(),
		GmtModified: time.Now(),
	}
	dup, err := TestTableDaoInstance().Create(context.Background(), order)
	assert.Equal(t, false, dup)
	assert.NoError(t, err)
	fmt.Println(dup, err)
}
